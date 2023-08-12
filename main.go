package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database")
	fmt.Println("Listening to port 4000")
	http.HandleFunc("/", helloHandler)

	coll := client.Database("test").Collection("products")
	name := "Industrial Bookshelf"
	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	error := http.ListenAndServe(":4000", nil)
	if error != nil {
		panic(error)
	}

}
func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World!")
}
