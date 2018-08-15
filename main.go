package main

import (
  "context"
  // "testing"
  "fmt"
  "github.com/mongodb/mongo-go-driver/bson"
  "github.com/mongodb/mongo-go-driver/mongo"
  // "github.com/mongodb/mongo-go-driver/mongo/findopt"
  // "github.com/stretchr/testify/require"
)

func main () {
  client, _ := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
  db := client.Database("gotest")
  // _, _ = db.RunCommand(context.Background(), bson.NewDocument(bson.EC.Int32("dropDatabase", 1)))

  coll := db.Collection("inventory")

  {
    // Start Example 1
    result, err := coll.InsertOne(
      context.Background(),
      bson.NewDocument(
        bson.EC.String("item", "canvas"),
        bson.EC.Int32("qty", 100),
        bson.EC.ArrayFromElements("tags",
        bson.VC.String("cotton"),
      ),
      bson.EC.SubDocumentFromElements("size",
        bson.EC.Int32("h", 28),
        bson.EC.Double("w", 35.5),
        bson.EC.String("uom", "cm"),
      ),
    ))

    if err != nil { fmt.Println("Shit, this is bad.")}
    fmt.Printf("%+v\n", result)
    fmt.Println(result.InsertedID)
  }
}
