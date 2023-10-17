package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database_name = "work_flow"

type Updates struct {
	Deleted_at time.Time `bson:"deleted_at"`
}

func Save(collectionName string, documents []interface{}) []primitive.ObjectID {
	return doMongoAction(collectionName, func(col *mongo.Collection) []primitive.ObjectID {
		if saveResult, err := col.InsertMany(context.Background(), documents); err != nil {
			panic(err)
		} else {
			savedObjectIds := make([]primitive.ObjectID, 0)
			for _, saved := range saveResult.InsertedIDs {
				if objectId, ok := saved.(primitive.ObjectID); ok {
					savedObjectIds = append(savedObjectIds, objectId)
				} else {
					panic(fmt.Errorf("saved id is not an objectId"))
				}
			}
			return savedObjectIds
		}
	})
}

func Get[T any](collectionName string, ids []primitive.ObjectID) []T {
	return doMongoAction(collectionName, func(c *mongo.Collection) []T {
		var returnItems []T
		if cursor, err := c.Find(context.Background(), bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}); err == nil {
			cursor.All(context.Background(), &returnItems)
		} else {
			panic(err)
		}
		return returnItems
	})
}

func Delete[T any](collectionName string, ids []primitive.ObjectID) {
	doMongoAction(collectionName, func(c *mongo.Collection) interface{} {
		c.UpdateMany(
			context.Background(),
			bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}},
			bson.D{
				{Key: "$set", Value: bson.D{{Key: "updates.deleted_at", Value: primitive.NewDateTimeFromTime(time.Now())}}},
			})
		return nil
	})
}

func doMongoAction[T any](collection_name string, action func(*mongo.Collection) T) T {
	connectionCtx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()
	if con, err := mongo.Connect(connectionCtx, options.Client().ApplyURI(getConnectionURI())); err != nil {
		panic(fmt.Errorf("cannot get mongo for connection %v", err))
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
	return uri
}
