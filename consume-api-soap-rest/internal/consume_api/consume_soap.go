package consumeapi

import (
	"consume-api-soap-rest/internal/domain/response"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumeApiCalculatorSoapAdd(c *gin.Context) {
	resp, err := http.Get("http://www.dneonline.com/calculator.asmx?WSDL")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var definitions response.Definitions
	body, err := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(body, &definitions)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(definitions)

}
