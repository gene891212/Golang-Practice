package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("iot").Collection("account")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	fmt.Println(cur)
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := struct {
			Account  string
			Password string
		}{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...
		fmt.Println(result)
		fmt.Printf("%+v\n", result)

	}
}
