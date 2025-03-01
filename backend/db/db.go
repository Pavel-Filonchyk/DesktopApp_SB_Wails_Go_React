package db

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx = context.TODO()

func ConnectDB() {
    var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://pfilonchyk:Luky2024PF@cluster0.cdb89.mongodb.net/Flats"))

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("MongoDB успешно подключен")
}

// Функция для получения клиента MongoDB
func GetClient() *mongo.Client {
    return client
}

// Получение коллекции из базы данных
func GetCollection(collectionName string) *mongo.Collection {
    return client.Database("Flats").Collection(collectionName)
}
