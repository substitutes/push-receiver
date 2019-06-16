package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/substitutes/push-receiver/model"
	"github.com/substitutes/substitutes/models"
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
