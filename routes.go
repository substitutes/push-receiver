package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/substitutes/push-receiver/model"
	"github.com/substitutes/substitutes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Controller is a routing manager
type Controller struct{}

// NewController creates a new routing manager
func NewController() *Controller { return &Controller{} }

// AddClass godoc
// @Summary Add a class
// @Description Add a class that has substitutes
// @Accept json
// @Produce json
// @Success 201 {array} model.SuccessResponse
// @Failure 500 {object} model.APIResponse
// @Param substitute body models.SubstituteResponse true "Substitute data"
// @Router /substitute/class [put]
func (ctl *Controller) AddClass(c *gin.Context) {
	var substitutes models.SubstituteResponse
	if c.BindJSON(&substitutes) != nil {
		model.NewMessage("Failed to bind JSON", nil).Throw(500, c)
		return
	}

	collection := Database.Collection("substitutes")
	_, err := collection.InsertOne(context.Background(), substitutes)
	if err != nil {
		model.NewMessage("Failed to save data to database!", err).Throw(500, c)
		return
	}

	c.JSON(201, substitutes.Meta.Class)
}

// DeleteRecord godoc
// @Summary Delete a class
// @Description Delete a class
// @Accept json
// @Produce json
// @Success 201 {array} model.SuccessResponse
// @Failure 500 {object} model.APIResponse
// @Failure 400 {object} model.APIResponse
// @Param id path string true "Object ID"
// @Router /substitute/class [put]
func (ctl *Controller) DeleteRecord(c *gin.Context) {
	// Validate ID
	collection := Database.Collection("substitutes")
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		model.NewMessage("Invalid ID", err).Throw(400, c)
		return
	}
	d, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		model.NewMessage("Failed to delete record", err).Throw(500, c)
		return
	}
	if d.DeletedCount > 0 {
		model.NewMessage("Deleted record", nil).Throw(200, c)
		return
	}
	model.NewMessage("No records found", errors.New("failed to find record")).Throw(400, c)
}

// ListClasses godoc
// @Summary list available classes
// @Description List all available classes and their last update
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Failure 500 {object} model.APIResponse
// Router /substitute/classes [get]
func (ctl *Controller) ListClasses(c *gin.Context) {
	collection := Database.Collection("substitutes")
	// Query
	classes, err := collection.Distinct(context.Background(), "meta.class", bson.D{{
		"meta.date", bson.D{{"$gt", time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)}},
	}})

	if err != nil {
		model.NewMessage("Failed to search for classes", err).Throw(500, c)
		return
	}

	c.JSON(200, classes)
}
