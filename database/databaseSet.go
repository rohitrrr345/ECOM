package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongo.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func DBSet() *mongo.Client {
	client,err:=mongo.NewClient(options.Client().ApplyURI("mongodb://2707"))
	if err!=nil {
		log.Fatal(err)
	}
	ctx,cancel:=context.WithTImeout(context.Background(),10*time.Second)
	defer.cancel()
	err=client.connect(ctx)
	if err!=nil {
		log.Fatal(err)
	}
}
