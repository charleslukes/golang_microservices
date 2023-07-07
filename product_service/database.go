package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "users"

type Product struct {
	Name      string `json:"name" bson:"name"`
	Desc      string `json:"desc" bson:"desc"`
	Type      string `json:"type" bson:"type"`
	Unit      string `json:"unit" bson:"unit"`
	Price     string `json:"price" bson:"price"`
	Available bool   `json:"available" bson:"available"`
	Supplier  string `json:"supplier" bson:"supplier"`
}

type MongoHandler struct {
	client   *mongo.Client
	database string
}

// MongoHandler Constructor
func NewHandler(address string) *MongoHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))

	if err != nil {
		fmt.Println("Panic: ", err)
		panic(err)
	}

	mh := &MongoHandler{
		client:   client,
		database: DefaultDatabase,
	}

	fmt.Println("connected to db successfully")
	return mh
}

func (mh *MongoHandler) GetOne(u *Product, filter interface{}) error {
	//Will automatically create a collection if not available
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(u)
	return err
}

func (mh *MongoHandler) Get(filter interface{}) []*Product {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result []*Product
	for cur.Next(ctx) {
		contact := &Product{}
		er := cur.Decode(contact)
		if er != nil {
			log.Fatal(er)
		}
		result = append(result, contact)
	}
	return result
}

func (mh *MongoHandler) AddOne(u *Product) (*mongo.InsertOneResult, error) {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, u)
	return result, err
}

func (mh *MongoHandler) Update(u *Product, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateMany(ctx, filter, update)
	return result, err
}

func (mh *MongoHandler) RemoveOne(filter interface{}) (*mongo.DeleteResult, error) {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, filter)
	return result, err
}
