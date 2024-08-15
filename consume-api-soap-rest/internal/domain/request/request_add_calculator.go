package request

import "encoding/xml"

// Define the structures based on the XML schema
type Envelope struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
	XmlnsSoap string   `xml:"xmlns:soap,attr"`
	Body      Body     `xml:"soap:Body"`
}

type Body struct {
	XMLName xml.Name `xml:"soap:Body"`
	Add     Add      `xml:"Add"`
}

type Add struct {
	XMLName xml.Name `xml:"Add"`
	Xmlns   string   `xml:"xmlns,attr"`
	IntA    int      `xml:"intA"`
	IntB    int      `xml:"intB"`
}

type RequestTwoVarAdd struct {
	IntA int `json:"intA"`
	IntB int `json:"intB"`
}
