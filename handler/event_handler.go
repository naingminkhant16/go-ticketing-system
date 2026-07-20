package handler

import (
	"net/http"
	"ticketing-system/common/response"
	"ticketing-system/entity/dto"
	"ticketing-system/service"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService *service.EventService
}

func NewEventHandler(eventService *service.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

//TODO : add pagination

func (handler *EventHandler) GetAllEvents(ctx *gin.Context) {
	events, err := handler.eventService.GetAllEvents()
	if err != nil {
		response.NotFound(ctx)
		return
	}

	response.Success(ctx, "Event List", events)
}

func (handler *EventHandler) CreateEvent(ctx *gin.Context) {
	var input dto.EventCreateDto

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	event, err := handler.eventService.Create(&input)
	if err != nil {
		ctx.Error(err)
		return
	}
	
	response.Created(ctx, event)
}
