package main

// @title Push Receiver API
// @version 1.0
// @description Substitutes project push receiver for push-cli
// @license.name AGPLv3
// @contact.email support@steinbart.xyz
// @schemes http https
// @securityDefinitions.basic BasicAuth
// @BasePath /api/v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginlogrus "github.com/toorop/gin-logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	_ "github.com/substitutes/push-receiver/docs"
)

// Global context
var Context context.Context
var Database *mongo.Database

func main() {

	r := gin.New()

	log := logrus.New()

	r.Use(ginlogrus.Logger(log), gin.Recovery())

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to create MongoDB driver: ", err)
	}

	Context, _ = context.WithTimeout(context.Background(), 10*time.Second)

	if client.Connect(Context) != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}
	if client.Ping(Context, readpref.Primary()) != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	Database = client.Database("push")

	c := NewController()

	log.Info("Connected to database: ", Database.Name())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		substitute := v1.Group("/substitute")
		{
			substitute.PUT("class", c.AddClass)
		}
	}

	if r.Run() != nil {
		logrus.Info("Terminated program")
	}

}
