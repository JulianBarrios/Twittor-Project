package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = DBconnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jubarrios:<password>@cluster0.da5zz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func DBconnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	log.Println("Conexion exitosa con la DB")
	return client
}

func CheckConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
