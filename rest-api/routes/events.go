package routes

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		errorResponse("Could not fetch events", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorResponse("Could not parse event id", http.StatusBadRequest, ctx, err)
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		errorResponse("Could not fetch event", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		errorResponse("Could not parse event", http.StatusBadRequest, ctx, err)
		return
	}

	event.UserID = ctx.GetInt64("userId")

	err = event.Save()

	if err != nil {
		errorResponse("Could not create event", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorResponse("Could not parse event id", http.StatusBadRequest, ctx, err)
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		errorResponse("Could not find event", http.StatusInternalServerError, ctx, err)
		return
	}

	if event.UserID != userId {
		errorResponse("Unauthorized!", http.StatusUnauthorized, ctx, errors.New("Not owner"))
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		errorResponse("Could not parse event", http.StatusBadRequest, ctx, err)

		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		errorResponse("Could not update event", http.StatusBadRequest, ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorResponse("Could not parse event id", http.StatusBadRequest, ctx, err)
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		errorResponse("Could not find event", http.StatusInternalServerError, ctx, err)
		return
	}

	if event.UserID != userId {
		errorResponse("Unauthorized!", http.StatusUnauthorized, ctx, errors.New("Not owner"))
		return
	}

	err = event.Delete()

	if err != nil {
		errorResponse("Could not delete event", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
