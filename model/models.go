package model

import (
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/structs"
)

type APIResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func NewMessage(message string, error error) *APIResponse {
	return &APIResponse{Message: message, Error: error}
}

func (a *APIResponse) Throw(code int, c *gin.Context) {
	c.JSON(code, a)
}

// Substitutes is a meta struct for holding substitute data
type Substitutes struct {
	Substitutes []structs.Substitute   `json:"substitutes"`
	Meta        structs.SubstituteMeta `json:"meta"`
}

type SuccessResponse struct {
	Class string `json:"class"`
}
