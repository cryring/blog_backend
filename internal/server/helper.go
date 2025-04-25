package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}
