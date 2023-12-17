package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorHandlers struct {
}

func (u *ErrorHandlers) HandleInternalError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (u *ErrorHandlers) HandleNotFoundError(c *gin.Context, status bool) {
	if !status {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
}
