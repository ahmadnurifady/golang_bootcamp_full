package main

import (
	"consume-api/domain"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func consumeREST(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendekode data"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func ConsumeEvent(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/api/v1/event/find-by-id?id=EVENT-004")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	var response domain.BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendekode data"})
		return
	}

	c.JSON(http.StatusOK, response)
}
