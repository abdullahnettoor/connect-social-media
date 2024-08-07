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

func NewUserRepository(driver neo4j.DriverWithContext) UserRepositoryInterface {
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
		userId: $ID,
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

	return newUser, nil
}

func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (u:User {username: $username}) RETURN u`

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

	return user, nil
}

func (r *UserRepository) FindUserByUserId(ctx context.Context, id string) (*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (n:User {userId: $id}) RETURN n`

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

	return user, nil
}

func (r *UserRepository) UpdateUserStatus(ctx context.Context, id string, status string, updatedAt string) (*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (n:User {userId: $id}) SET n.status = $status, n.updatedAt = $updatedAt RETURN n`

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

func (r *UserRepository) FollowUser(ctx context.Context, userId, followedId string) error {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `
	MATCH (n:User{userId:$userId})
	MATCH (f:User{userId:$followedId})
	MERGE (n)-[:FOLLOWS]->(f)
	`

	params := map[string]any{
		"userId": userId, "followedId": followedId,
	}

	fmt.Println("Params is", params)

	_, err := session.Run(ctx, cypher, params)

	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return err
	}
	return nil
}

func (r *UserRepository) UnfollowUser(ctx context.Context, userId, followedId string) error {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `
	MATCH (n:User {userId: $userId})
	MATCH (f:User {userId: $followedId})
	MATCH (n)-[r:FOLLOWS]->(f)
	DELETE r
	`

	params := map[string]any{
		"userId": userId, "followedId": followedId,
	}

	fmt.Println("Params is", params)

	_, err := session.Run(ctx, cypher, params)

	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetFollowers(ctx context.Context, userId string) ([]*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `
	MATCH (n:User {userId: $userId})
	MATCH (f:User)
	MATCH (f)-[r:FOLLOWS]->(n)
	RETURN f.username as username, 
			f.avatar as avatar, 
			f.userId as userId
	`

	params := map[string]any{
		"userId": userId,
	}

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}

	var followers = make([]*entity.User, 0)
	for result.Next(ctx) {
		f := &entity.User{}
		fmt.Println(result.Record().AsMap())
		err := conv.MapToStruct(result.Record().AsMap(), &f)
		if err != nil {
			log.Println("Error occurred while converting followerMap to follower:", err)
			return nil, err
		}
		followers = append(followers, f)
	}

	return followers, nil
}

func (r *UserRepository) GetFollowing(ctx context.Context, userId string) ([]*entity.User, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `
	MATCH (n:User {userId: $userId})
	MATCH (f:User)
	MATCH (n)-[r:FOLLOWS]->(f)
	RETURN f.username as username, 
			f.avatar as avatar, 
			f.userId as userId
	`

	params := map[string]any{
		"userId": userId,
	}

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}

	var following = make([]*entity.User, 0)
	for result.Next(ctx) {
		f := &entity.User{}
		fmt.Println(result.Record().AsMap())
		err := conv.MapToStruct(result.Record().AsMap(), &f)
		if err != nil {
			log.Println("Error occurred while converting followingMap to following:", err)
			return nil, err
		}
		following = append(following, f)
	}

	return following, nil
}
