package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorResponse("Could not parse event id", http.StatusBadRequest, ctx, err)
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		errorResponse("Could not find event", http.StatusInternalServerError, ctx, err)
		return
	}

	err = event.Register(userId)

	if err != nil {
		errorResponse("Could not register", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.Cancel(userId)

	if err != nil {
		errorResponse("Could not cancel", http.StatusInternalServerError, ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Canceled!"})
}
