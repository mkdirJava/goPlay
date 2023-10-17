package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database_name = "work_flow"

func Save(collectionName string, documents []interface{}) []interface{} {
	return doMongoAction(collectionName, func(col *mongo.Collection) []interface{} {
		if saveResult, err := col.InsertMany(context.Background(), documents); err != nil {
			panic(err)
		} else {
			return saveResult.InsertedIDs
		}
	})
}

func Get[T any](collectionName string, ids []uint) []T {
	return doMongoAction(collectionName, func(c *mongo.Collection) []T {
		var returnItems []T
		if cursor, err := c.Find(context.Background(), bson.D{}); err != nil {
			cursor.All(context.Background(), &returnItems)
		} else {
			panic(err)
		}
		return returnItems
	})
}

func Delete[T any](collectionName string, ids []uint) {
	doMongoAction(collectionName, func(c *mongo.Collection) interface{} {
		c.UpdateMany(
			context.Background(),
			bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}},
			bson.D{
				{Key: "$set", Value: bson.D{{Key: "deleted", Value: "true"}}},
			})
		return nil
	})
}

func doMongoAction[T any](collection_name string, action func(*mongo.Collection) T) T {
	connectionCtx := getConnectonContext()
	if con, err := mongo.Connect(connectionCtx, options.Client().ApplyURI(getConnectionURI())); err != nil {
		panic("Cannot get mongo for connection")
	} else {
		defer con.Disconnect(connectionCtx)
		return action(con.Database(database_name).Collection(collection_name))
	}
}

func getConnectionURI() string {
	mongo_uri := "MONGO_URI"
	uri := os.Getenv(mongo_uri)
	if uri == "" {
		panic(fmt.Sprintf("Env var must be set %v", mongo_uri))
	}
	return mongo_uri
}

func getConnectonContext() context.Context {
	if ctx, err := context.WithTimeout(context.Background(), 10*time.Second); err != nil {
		panic("Cannot get background for connection")
	} else {
		return ctx
	}
}
