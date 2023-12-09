package db

import (
	"context"

	"github.com/bhoopendrau/tailscale-ui-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func Init() {
	c := config.GetConfig()
	uri := c.GetString("db.uri")
	var err error
	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

func GetDB() *mongo.Client {
	return db
}
