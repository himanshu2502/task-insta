package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	FieldStr  string `json: "Field Str"`
	FieldInt  int    `json: "Field Int"`
	FieldBool bool   `json: "Field Bool"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("Database").Collection(" COllection")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := MongoField{
		Id        string `json:"id"`
	    Caption   string `json:"caption"`
	    Image_url string `json:"Image_url"`
	    time      string `json:"time"`
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}
}
