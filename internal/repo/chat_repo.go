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

type ChatRepository struct {
	db neo4j.DriverWithContext
}

func NewChatRepository(db neo4j.DriverWithContext) ChatRepositoryInterface {
	return &ChatRepository{db}
}

func (r *ChatRepository) CreateMessage(ctx context.Context, message *entity.Message) (*entity.Message, error) {

	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	params, err := conv.StructToMap(message)
	if err != nil {
		log.Println("Error occurred while parsing struct to map:", err)
		return nil, err
	}

	fmt.Println("Params is:", params)

	cypher := `
	MATCH (sender:User {userId: $SenderID})
	MATCH (recipient:User {userId: $RecipientID})
	CREATE (msg:Message {
		messageId: $ID,
		message: $Message,
		createdAt:$CreatedAt,
		receivedAt: $ReceivedAt,
		readAt: $ReadAt
	})
	CREATE (sender)-[:SEND_MSG]->(msg)<-[:RECEIVED_MSG]-(recipient)
	RETURN msg;
	`

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		log.Println("Error occurred while Executing cypher:", err)
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		log.Println("Error occurred while retrieving new msg:", err)
		return nil, err
	}

	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, message); err != nil {
		log.Println("Error occurred while converting commentMap to msg:", err)
		return nil, err
	}

	return message, nil
}
