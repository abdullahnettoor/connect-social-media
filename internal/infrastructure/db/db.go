package db

import (
	"context"
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectDb(cfg *config.Config) (neo4j.DriverWithContext, error) {
	log.Println("NEO4J: Trying to connect with LOCALHOST")
	driver, err := neo4j.NewDriverWithContext(
		// Try connecting with localhost
		cfg.DbUri,
		neo4j.BasicAuth(cfg.DbUsername, cfg.DbPassword, ""),
	)
	if err != nil {
		log.Println("NEO4J:", err.Error())
		return nil, err
	}

	log.Println("NEO4J: Verifying Connectivity")
	err = driver.VerifyConnectivity(context.Background())
	if !neo4j.IsConnectivityError(err) && err != nil {
		log.Println("NEO4J:", err.Error())
		return nil, err
	}

	// Try connecting to containerized db
	if neo4j.IsConnectivityError(err) {
		log.Println("NEO4J: Trying to connect with CONTAINER")
		driver, err = neo4j.NewDriverWithContext(
			cfg.DbContainerUri,
			neo4j.BasicAuth(cfg.DbUsername, cfg.DbPassword, ""),
		)
		if err != nil {
			log.Println("NEO4J:", err.Error())
			return nil, err
		}

		log.Println("NEO4J: Verifying Connectivity")
		err = driver.VerifyConnectivity(context.Background())
		if err != nil {
			log.Println("NEO4J:", err.Error())
			return nil, err
		}
	}

	log.Println("NEO4J: Connection Established")
	return driver, nil
}
