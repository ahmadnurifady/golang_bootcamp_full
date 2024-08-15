package domain

import "encoding/xml"

// Root struct for SOAP Envelope
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

// Body struct for SOAP Body
type Body struct {
	XMLName               xml.Name              `xml:"Body"`
	NumberToWordsResponse NumberToWordsResponse `xml:"NumberToWordsResponse"`
}

// NumberToWordsResponse struct for the response
type NumberToWordsResponse struct {
	XMLName             xml.Name `xml:"NumberToWordsResponse"`
	xmlns               string   `xml:"xmlns,attr"`
	NumberToWordsResult string   `xml:"NumberToWordsResult"`
}
