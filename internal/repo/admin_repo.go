package repo

import (
	"context"
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/pkg/conv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type AdminRepository struct {
	db neo4j.DriverWithContext
}

func NewAdminRepository(db neo4j.DriverWithContext) *AdminRepository {
	return &AdminRepository{db}
}

func (r *AdminRepository) FindAdminByEmail(ctx context.Context, email string) (*entity.Admin, error) {
	session := r.db.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	cypher := `MATCH (a:Admin) WHERE a.email = $email RETURN a`

	result, err := session.Run(ctx, cypher, map[string]any{"email": email})
	if err != nil {
		log.Println("Admin Repo error:", err)
		return nil, err
	}
	if !result.Peek(ctx) {
		return nil, e.ErrAdminNotFound
	}

	record := result.Record()
	var admin = &entity.Admin{}
	if err := conv.MapToStruct(record.Values[0].(dbtype.Node).Props, admin); err != nil {
		log.Println("Error occurred while converting userMap to admin:", err)
		return nil, err
	}

	return admin, nil
}
