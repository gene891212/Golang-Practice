package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Account string `json:"account"`
	Password string `json:"password"`
}

func insertData(account, password string)  {
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
	podcastResult, err := collection.InsertOne(ctx, bson.D{
		{"account", account},
		{"password", password},
	})
	fmt.Println(podcastResult)
}

func allUser(w http.ResponseWriter, r *http.Request)  {
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
	results := []User{}
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := User{}
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result...
		results = append(results, result)
		fmt.Println(result)
	}
	j, _ :=json.Marshal(results)
	fmt.Fprintf(w, string(j))
}

func addAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println(k, v)
	}
	infomation := User{
		r.FormValue("account"),
		r.FormValue("password"),
	}
	insertData(r.FormValue("account"), r.FormValue("password"))
	j, err := json.Marshal(infomation)
	if err != nil {
		log.Fatal("json format error", err)
	}
	fmt.Println(j)
	//fmt.Fprintf(w, string(j))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", addAccount)
	http.HandleFunc("/all_user", allUser)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}