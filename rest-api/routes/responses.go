package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func errorResponse(message string, status int, ctx *gin.Context, err error) {
	ctx.JSON(
		status,
		gin.H{
			"message": message,
			"error":   fmt.Sprintf("%v", err),
		},
	)
}
