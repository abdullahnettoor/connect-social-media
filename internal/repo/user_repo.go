package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/pkg/conv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type UserRepository struct {
	db neo4j.DriverWithContext
}

func NewUserRepository(driver neo4j.DriverWithContext) *UserRepository {
	return &UserRepository{db: driver}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params, err := conv.StructToMap(user)
	if err != nil {
		log.Println("Error occurred while parsing struct to map:", err)
		return nil, err
	}
	fmt.Println("Params is", params)

	cypher := `MATCH (u:User)
	WHERE u.username = $Username OR u.email = $Email
	RETURN u`

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}
	if result.Peek(ctx) {
		rec, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}

		email := rec[0].Values[0].(dbtype.Node).Props["email"].(string)
		username := rec[0].Values[0].(dbtype.Node).Props["username"].(string)

		if user.Email == email && user.Username == username {
			err = e.ErrEmailAndUsernameConflict
		} else if user.Email == email {
			err = e.ErrEmailConflict
		} else {
			err = e.ErrUsernameConflict
		}
		return nil, err
	}

	cypher = `CREATE (u:User {
		username: $Username, 
		email: $Email, 
		password: $Password, 
		status: $Status, 
		fullName: $FullName, 
		accountType: $AccountType,
		createdAt: $CreatedAt,
		updatedAt: $UpdatedAt
	})
    RETURN u`

	result, err = session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		log.Println("Error occurred while retrieving new user:", err)
		return nil, err
	}

	var newUser = &entity.User{}
	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, newUser); err != nil {
		log.Println("Error occurred while converting userMap to user:", err)
		return nil, err
	}
	newUser.ID = record.Values[0].(dbtype.Node).GetId()

	return newUser, nil
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (u:User) WHERE u.username = $username RETURN u`

	result, err := session.Run(ctx, cypher, map[string]any{"username": username})
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}
	if !result.Peek(ctx) {
		return nil, e.ErrUserNotFound
	}

	record := result.Record()
	var user = &entity.User{}
	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, user); err != nil {
		log.Println("Error occurred while converting userMap to user:", err)
		return nil, err
	}
	user.Password = record.Values[0].(dbtype.Node).Props["password"].(string)
	user.ID = record.Values[0].(dbtype.Node).GetId()

	return user, nil
}

func (r *UserRepository) FindUserByUserId(ctx context.Context, id int64) (*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (n:User) WHERE id(n) = $id RETURN n`

	result, err := session.Run(ctx, cypher, map[string]any{"id": id})
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}
	if !result.Peek(ctx) {
		return nil, e.ErrUserNotFound
	}

	record := result.Record()
	var user = &entity.User{}
	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, user); err != nil {
		log.Println("Error occurred while converting userMap to user:", err)
		return nil, err
	}
	user.ID = record.Values[0].(dbtype.Node).GetId()

	return user, nil
}

func (r *UserRepository) UpdateUserStatus(ctx context.Context, id int64, status string, updatedAt string) (*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (n:User) WHERE id(n) = $id SET n.status = $status, n.updatedAt = $updatedAt RETURN n`

	result, err := session.Run(ctx, cypher, map[string]any{"id": id, "status": status, "updatedAt": updatedAt})
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}
	if !result.Peek(ctx) {
		return nil, e.ErrUserNotFound
	}

	record := result.Record()
	var user = &entity.User{}
	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, user); err != nil {
		log.Println("Error occurred while converting userMap to user:", err)
		return nil, err
	}
	user.ID = record.Values[0].(dbtype.Node).GetId()

	return user, nil
}

func (r *UserRepository) RemoveUserByEmail(ctx context.Context, email string) error {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (n:User{email:$email}) DELETE n`

	_, err := session.Run(ctx, cypher, map[string]any{"email": email})
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return err
	}

	return nil
}
