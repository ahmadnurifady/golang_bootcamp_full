package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	router := gin.Default()

// 	router.Run(":8080")
// }

func getUserInfo(c *gin.Context) {
	// SOAP request payload
	requestPayload := `
  <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://example.com/soap/service">
   <soapenv:Header/>
   <soapenv:Body>
    <ser:getUserDetails>
     <ser:userId>%s</ser:userId>
    </ser:getUserDetails>
   </soapenv:Body>
  </soapenv:Envelope>
 `

	// Replace %s with the user ID received in the request
	userID := c.PostForm("userID")
	payload := fmt.Sprintf(requestPayload, userID)

	// SOAP endpoint URL
	endpointURL := "http://example.com/soap/service"

	// Create a new HTTP request
	req, err := http.NewRequest("POST", endpointURL, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	// Set the SOAP headers and content type
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Set("SOAPAction", "getUserDetails")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the SOAP response into a struct
	type Response struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			GetUserDetailsResponse struct {
				XMLName   xml.Name `xml:"getUserDetailsResponse"`
				UserInfo  string   `xml:"userInfo"`
				ErrorCode string   `xml:"errorCode"`
			}
		}
	}
	var soapResponse Response
	err = xml.Unmarshal(body, &soapResponse)
	if err != nil {
		log.Fatal(err)
	}

	// Check for errors in the SOAP response
	if soapResponse.Body.GetUserDetailsResponse.ErrorCode != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user details"})
		return
	}

	// Return the user info from the SOAP response
	c.JSON(http.StatusOK, gin.H{"userInfo": soapResponse.Body.GetUserDetailsResponse.UserInfo})
}
