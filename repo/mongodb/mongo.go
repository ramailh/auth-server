package mongodb

import (
	"context"
	"fmt"
	"github.com/ramailh/auth-server/props"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserCollection = "user"
)

type mongoClient struct {
	client *mongo.Client
}

func NewMongoClient() mongoClient {
	hostUrl := fmt.Sprintf("mongodb://%s/?connect=direct", props.MongoHost)
	opt := options.Client().ApplyURI(hostUrl)
	if props.MongoUsername != "" {
		opt.SetAuth(options.Credential{Username: props.MongoUsername, Password: props.MongoPassword})
	}

	client, err := mongo.NewClient(opt)
	if err != nil {
		log.Println(err)
	}

	client.Connect(context.Background())

	// client, err := mongo.Connect(context.Background())
	// if err != nil {
	// 	log.Println(err)
	// }

	return mongoClient{client: client}
}

func (cl mongoClient) Insert(doc interface{}) (interface{}, error) {
	defer cl.client.Disconnect(context.Background())

	return cl.client.Database(props.DBName).Collection(UserCollection).InsertOne(context.Background(), doc)
}

func (cl mongoClient) Find(filter bson.M, opt ...*options.FindOptions) ([]map[string]interface{}, error) {
	defer cl.client.Disconnect(context.Background())

	res, err := cl.client.Database(props.DBName).Collection(UserCollection).Find(context.Background(), filter, opt...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data []map[string]interface{}
	if err = res.All(context.Background(), &data); err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

func (cl mongoClient) FindOne(filter bson.M) (map[string]interface{}, error) {
	defer cl.client.Disconnect(context.Background())

	data := make(map[string]interface{})
	err := cl.client.Database(props.DBName).Collection(UserCollection).FindOne(context.Background(), filter).Decode(&data)

	return data, err
}

func (cl mongoClient) Update(filter, updateDoc bson.M) (interface{}, error) {
	defer cl.client.Disconnect(context.Background())

	return cl.client.Database(props.DBName).Collection(UserCollection).UpdateOne(context.Background(), filter, updateDoc)
}

func (cl mongoClient) Delete(filter bson.M) (interface{}, error) {
	defer cl.client.Disconnect(context.Background())

	return cl.client.Database(props.DBName).Collection(UserCollection).DeleteOne(context.Background(), filter)
}
