package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var ctx = context.Background()

func getDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:pass@localhost:27017/admin")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database("animal_db")
	return db, nil
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the world of animals.")
	})

	r.GET("/animals", func(c *gin.Context) {
		db, err := getDB()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Fatal(err)
		}

		cursor, err := db.Collection("animal_tb").Find(ctx, bson.M{})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Fatal(err)
		}

		var animals []bson.M
		if err = cursor.All(ctx, &animals); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"animals": animals})
	})

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}
