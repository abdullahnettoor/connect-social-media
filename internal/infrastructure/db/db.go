package db

import (
	"context"
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectDb(cfg *config.Config) (neo4j.DriverWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(
		cfg.DbUri,
		neo4j.BasicAuth(cfg.DbUsername, cfg.DbPassword, ""),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	// defer driver.Close(context.Background())

	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Neo4j Connection Established")
	return driver, nil
}
