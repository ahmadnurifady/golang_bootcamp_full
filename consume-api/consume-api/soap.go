package main

import (
	"bytes"
	"consume-api/domain"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func consumeSOAP(c *gin.Context) {
	soapRequest := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">
      <ubiNum>500</ubiNum>
    </NumberToWords>
  </soap:Body>
</soap:Envelope>`

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.dataaccess.com/webservicesserver/NumberConversion.wso", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim permintaan"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
		return
	}

	fmt.Println(string(body))

	var envelope domain.Envelope
	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println("HASIL CONVERT KE STRUCT ==", envelope)
	fmt.Println("body ==", envelope.Body.XMLName)

	c.Data(http.StatusOK, "text/xml", body)
}

func consumeSoapCalculator(c *gin.Context) {
	resp, err := http.Get("http://www.dneonline.com/calculator.asmx?WSDL")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	var definitions domain.Definitions
	body, err := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(body, &definitions)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// if err := json.NewDecoder(resp.Body).Decode(&add); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendekode data"})
	// 	return
	// }

	c.JSON(http.StatusOK, definitions)
}

func consumeSoapAdd(c *gin.Context) {
	soapRequest := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
	    <Add xmlns="http://tempuri.org/">
	      <intA>10</intA>
	      <intB>2</intB>
	    </Add>
	  </soap:Body>
	</soap:Envelope>`

	// jsonData := []byte(`{
	// 	"Body": {
	// 		"Add": {
	// 			"IntA": 10,
	// 			"IntB": 2
	// 		}
	// 	}
	// }`)

	// // Unmarshal JSON
	// var envelope Envelope
	// if err := json.Unmarshal(jsonData, &envelope); err != nil {
	// 	log.Fatalf("Error unmarshalling JSON: %v", err)
	// }

	// // Convert to XML
	// xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Error marshalling XML: %v", err)
	// }

	// // Add XML header
	// xmlData = append([]byte(xml.Header), xmlData...)

	// request, err := io.ReadAll(c.Request.Body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim permintaan"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
		return
	}

	// fmt.Println(string(body))

	// var envelope domain.Envelope
	// err = xml.Unmarshal(body, &envelope)
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	return
	// }

	// fmt.Println("HASIL CONVERT KE STRUCT ==", envelope)
	// fmt.Println("body ==", envelope.Body.XMLName)

	c.Data(http.StatusOK, "text/xml", body)
}
