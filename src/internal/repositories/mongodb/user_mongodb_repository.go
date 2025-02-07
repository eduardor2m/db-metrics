package mongodb

import (
	"context"
	"fmt"
	"os"

	"github.com/eduardor2m/db-metrics/src/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// var _ repository. = &UserMongodbRepository{}

type UserMongodbRepository struct {
	connectorManager
}

func (instance *UserMongodbRepository) Create(userReceived models.User) error {
	fmt.Println("usuario", userReceived)
	conn, err := instance.connectorManager.getConnection()
	if err != nil {
		return err
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	document := bson.M{
		"id":          userReceived.ID(),
		"name":        userReceived.Name(),
		"email":       userReceived.Email(),
		"preferences": userReceived.Preferences(),
	}

	collectionName := os.Getenv("MONGODB_COLLECTION_USER")

	_, err = conn.Collection(collectionName).InsertOne(ctx, document)
	if err != nil {
		return err
	}

	return nil
}
