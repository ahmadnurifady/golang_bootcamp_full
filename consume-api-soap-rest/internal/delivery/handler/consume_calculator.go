package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsumeCalculator interface {
	Add
	Route
	AddRequestJson
}

type AddRequestJson interface {
	AddRequestJson(ctx *gin.Context)
}

type Add interface {
	CalculatorAdd(ctx *gin.Context)
}

type Route interface {
	Route()
}

type consumeCalculator struct {
	rg *gin.RouterGroup
}

// AddRequestJson implements ConsumeCalculator.
func (c *consumeCalculator) AddRequestJson(ctx *gin.Context) {
	panic("unimplemented")
}

// CalculatorAdd implements ConsumeCalculator.
func (c *consumeCalculator) CalculatorAdd(ctx *gin.Context) {
	soapRequest := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
	    <Add xmlns="http://tempuri.org/">
	      <intA>10</intA>
	      <intB>2</intB>
	    </Add>
	  </soap:Body>
	</soap:Envelope>`

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengirim permintaan"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
		return
	}

	ctx.Data(http.StatusOK, "text/xml", body)

}

// Route implements ConsumeCalculator.
func (c *consumeCalculator) Route() {
	cg := c.rg.Group("/consume-calculator")
	cg.GET("/get-add", c.CalculatorAdd)
}

func NewConsumeCalculator(rg *gin.RouterGroup) ConsumeCalculator {
	return &consumeCalculator{rg: rg}
}
