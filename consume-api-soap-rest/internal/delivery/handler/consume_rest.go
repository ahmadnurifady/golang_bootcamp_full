package handler

import (
	"consume-api-soap-rest/internal/domain/response"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsumeRestEvent interface {
	ConsumeGetEvent
	RouteConsumeRestEvent
}

type ConsumeGetEvent interface {
	ConsumeGetEvent(ctx *gin.Context)
}

type RouteConsumeRestEvent interface {
	Route()
}

type consumeRestEvent struct {
	rg *gin.RouterGroup
}

// ConsumeGetEvent implements ConsumeRestEvent.
func (c *consumeRestEvent) ConsumeGetEvent(ctx *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/v1/event/find-by-id?id=EVENT-004")
	if err != nil {
		// fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	var response response.BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendekode data"})
		// fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Route implements ConsumeRestEvent.
func (c *consumeRestEvent) Route() {
	cg := c.rg.Group("/consume-rest")
	cg.GET("/get-add", c.ConsumeGetEvent)
}

func NewConsumeRestEvent(rg *gin.RouterGroup) ConsumeRestEvent {
	return &consumeRestEvent{rg: rg}
}
