package consumeapi

import (
	"consume-api-soap-rest/internal/domain/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func ConsumeEvent() {
	resp, err := http.Get("http://localhost:8080/api/v1/event/find-by-id?id=EVENT-004")
	if err != nil {
		fmt.Println(err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	var response response.BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendekode data"})
		fmt.Println(err)
		return
	}

	// c.JSON(http.StatusOK, response)
	fmt.Println(response)

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Message : ", response.Message)
	fmt.Println("Data : ", response.Data)
}
