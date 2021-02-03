package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/iot/libs"
	"github.com/iot/stru"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func allUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	libs.SetupResponse(&w)
	r.ParseForm()

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
	var find bson.M
	if value := r.FormValue("account"); value != "" {
		find = bson.M{"account": value}
	} else {
		find = bson.M{}
	}
	cur, err := collection.Find(context.Background(), find)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	// var test []interface{}
	// err = cur.All(context.Background(), &test)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Fprintln(w, test)

	results := []stru.UserInfo{}
	for cur.Next(context.Background()) {
		result := stru.UserInfo{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	j, _ := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("./find.html")
	t.Execute(w, struct{ Message string }{Message: string(j)})

}

func createUser(w http.ResponseWriter, r *http.Request) {
	libs.SetupResponse(&w)
	libs.CreateUser(w, r)
}

func main() {
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/", allUser)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
