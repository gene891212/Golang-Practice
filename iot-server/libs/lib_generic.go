package libs

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/iot/stru"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertData ...
func InsertData(user stru.UserInfo) {
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

// SetupResponse ...
func SetupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// RaiseError ...
func RaiseError(w http.ResponseWriter, status int, message string, path string) {
	now, _ := time.Now().MarshalText()

	respError := stru.ErrorResp{
		Error: stru.ErrorMessage{
			Timestamp: string(now),
			Status:    status,
			Message:   message,
			Path:      path,
		},
	}
	json.NewEncoder(w).Encode(respError)
}
