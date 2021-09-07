package bd

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var user = ""
var password = ""

//connection var as object
var connection = Connection()

var uri = options.Client().ApplyURI("mongodb+srv://"+ user +":"+ password +"@cluster0.eguxz.mongodb.net/test?authSource=admin&replicaSet=atlas-89eeig-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

// Connection method
func Connection() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), uri)

    if err!= nil {
        log.Fatal(err)
        return client
    }

    err = client.Ping(context.TODO(), nil)

    if err!= nil {
        log.Fatal(err.Error())
        return client
    }

    log.Printf("Success connection.")
    return client
}

// connectionIsSuccess method
func connectionIsSuccess() int {
    err := connection.Ping(context.TODO(), nil)
    if err!= nil {
        log.Fatal(err)
        return 0
    }
    return 1
}