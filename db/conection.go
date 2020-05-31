package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ConnectionURI = "mongodb+srv://amendez:Judoneyba*1234@chat-mega-cluster-fdu9q.mongodb.net/test?retryWrites=true&w=majority";

/* Object connection to db */
var MongoCN = ConnectDB()

var clientOptions =  options.Client().ApplyURI(ConnectionURI)

/* ConnectDB return db connection */
func ConnectDB() *mongo.Client {
 client, err := mongo.Connect(context.TODO(), clientOptions)

 if err != nil {
	log.Fatal(err.Error())
 }

 log.Println("Conexión Exitosa")
 return client
}