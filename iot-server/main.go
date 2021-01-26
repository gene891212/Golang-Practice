package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func insertData(user User) {
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
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
}

func allUser(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(body)
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

	//var test []interface{}
	//err = cur.All(context.Background(), &test)
	//if err != nil {
	//	log.Fatal(err)
	//}

	results := []User{}
	for cur.Next(context.Background()) {
		result := User{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	j, _ := json.Marshal(results)
	_, err = fmt.Fprintf(w, string(j))
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	infomation := User{
		r.FormValue("account"),
		r.FormValue("password"),
	}
	insertData(infomation)

	j, err := json.Marshal(infomation)
	if err != nil {
		log.Fatal("json format error", err)
	}
	//w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, string(j))
}

func main() {
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/", allUser)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
