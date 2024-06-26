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

func (r *CommentRepository) GetCommentsOfPost(ctx context.Context, postId string) ([]*entity.Comment, error) {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params := map[string]any{
		"postId": postId,
	}
	cypher := `MATCH 
	(post:Post {postId: $postId})
	-[:HAS_COMMENT]-
	(comment:Comment)
	-[:WROTE_COMMENT]-
	(u:User)
	RETURN comment, u.username as username, u.avatar as avatar
	`

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}
	records, err := result.Collect(ctx)
	if err != nil {
		return nil, err
	}

	var comments []*entity.Comment
	for _, r := range records {
		fmt.Println("Record is", r.AsMap())

		var c = &entity.Comment{}
		err := conv.MapToStruct(r.AsMap(), c)
		if err != nil {
			log.Print("Error is: ", err)
			// return nil, err
		}
		err = conv.MapToStruct(r.AsMap()["comment"].(dbtype.Node).Props, c)
		if err != nil {
			log.Print("Error is: ", err)
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}