package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	"github.com/abdullahnettoor/connect-social-media/pkg/conv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type PostRepository struct {
	db neo4j.DriverWithContext
}

func NewPostRepository(db neo4j.DriverWithContext) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) Create(ctx context.Context, userId string, post *entity.Post) (*entity.Post, error) {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params, err := conv.StructToMap(post)
	if err != nil {
		log.Println("Error occurred while parsing struct to map:", err)
		return nil, err
	}
	params["UserId"] = userId
	fmt.Println("Params is", params)

	cypher := `CREATE (p:Post {
		id :$ID,
		description: $Description, 
		location: $Location, 
		mediaUrls: $MediaUrls, 
		isBlocked: $IsBlocked,
		createdAt: $CreatedAt,
		updatedAt: $UpdatedAt
	}) 
	WITH p
	MATCH (u:User { id: $UserId})
	CREATE (u)-[:POSTED]->(p)
	RETURN p	
	`
	// RETURN p, u.id AS userId, u.username AS username, u.profileUrl AS profileUrl
	// cypher := `CREATE (p:Post {
	// 	description: $Description,
	// 	location: $Location,
	// 	mediaUrls: $MediaUrls,
	// 	isBlocked: $IsBlocked,
	// 	createdAt: $CreatedAt,
	// 	updatedAt: $UpdatedAt})
	// WITH p
	// MATCH (u:User) WHERE id(u) = $UserId
	// CREATE (u)-[:POSTED]->(p)
	// RETURN p`

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		log.Println("Error occurred while retrieving new post:", err)
		return nil, err
	}

	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, post); err != nil {
		log.Println("Error occurred while converting userMap to user:", err)
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) LikePost(ctx context.Context, userId, postId string) error {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (p:Post {
		id :$postId}) WITH p
	MATCH (u:User { id: $userId})
	MERGE (u)-[:LIKED]->(p)
	`

	_, err := session.Run(ctx, cypher, map[string]any{"userId": userId, "postId": postId})
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return err
	}

	return nil
}

func (r *PostRepository) UnlikePost(ctx context.Context, userId, postId string) error {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH 
	(:Post {id :$postId})
	-[r:LIKED]->
	(:User { id: $userId})
	DELETE r
	`

	_, err := session.Run(ctx, cypher, map[string]any{"userId": userId, "postId": postId})
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return err
	}

	return nil
}
