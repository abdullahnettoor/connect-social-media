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

type CommentRepository struct {
	db neo4j.DriverWithContext
}

func NewCommentRepository(db neo4j.DriverWithContext) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params, err := conv.StructToMap(comment)
	if err != nil {
		log.Println("Error occurred while parsing struct to map:", err)
		return nil, err
	}

	fmt.Println("Params is:", params)

	cypher := `
	MATCH (user:User {userId: $UserID})
	MATCH (post:Post {postId: $PostID})
	CREATE (comment:Comment {
		commentId: $ID,
		comment: $Comment,
		createdAt:$CreatedAt
	})
	CREATE (user)-[:WROTE_COMMENT]->(comment)<-[:HAS_COMMENT]-(post)
	RETURN comment;
	`

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		log.Println("Error occurred while retrieving new comment:", err)
		return nil, err
	}

	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, comment); err != nil {
		log.Println("Error occurred while converting commentMap to comment:", err)
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) Delete(ctx context.Context, userId, commentId string) error {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params := map[string]any{
		"userId":    userId,
		"commentId": commentId,
	}
	fmt.Println("Params", params)

	cypher := `MATCH 
	(user:User {userId: $userId})
	-[:WROTE_COMMENT]->
	(comment:Comment {commentId: $commentId})
	DETACH DELETE comment
	`

	_, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return err
	}
	fmt.Println("Hooray..........")

	return nil
}
