package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// REST API endpoint
	r.GET("/consume/rest", consumeREST)
	// SOAP API endpoint
	r.GET("/consume/soap", consumeSOAP)
	r.GET("/consume/event", ConsumeEvent)
	r.GET("/consume/calculator/add", consumeSoapCalculator)
	r.GET("/consume/soap/cal/add", consumeSoapAdd)
	r.POST("/api/getUserInfo", getUserInfo)

	r.Run(":8081") // Jalankan server pada port 8080
}
