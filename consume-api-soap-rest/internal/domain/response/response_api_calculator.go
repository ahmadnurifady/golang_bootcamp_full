package response

import "encoding/xml"

type Definitions struct {
	XMLName   xml.Name   `xml:"definitions"`
	TargetNS  string     `xml:"targetNamespace,attr"`
	Types     Types      `xml:"types"`
	Messages  []Message  `xml:"message"`
	PortTypes []PortType `xml:"portType"`
	Bindings  []Binding  `xml:"binding"`
	Services  []Service  `xml:"service"`
}

type Types struct {
	Schema Schema `xml:"schema"`
}

type Schema struct {
	ElementFormDefault string    `xml:"elementFormDefault,attr"`
	TargetNamespace    string    `xml:"targetNamespace,attr"`
	Elements           []Element `xml:"element"`
}

type Element struct {
	Name        string      `xml:"name,attr"`
	ComplexType ComplexType `xml:"complexType"`
}

type ComplexType struct {
	Sequence Sequence `xml:"sequence"`
}

type Sequence struct {
	Elements []SubElement `xml:"element"`
}

type SubElement struct {
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	MinOccurs int    `xml:"minOccurs,attr"`
	MaxOccurs int    `xml:"maxOccurs,attr"`
}

type Message struct {
	Name  string `xml:"name,attr"`
	Parts []Part `xml:"part"`
}

type Part struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type PortType struct {
	Name       string      `xml:"name,attr"`
	Operations []Operation `xml:"operation"`
}

type Operation struct {
	Name          string      `xml:"name,attr"`
	Documentation string      `xml:"documentation"`
	Input         InputOutput `xml:"input"`
	Output        InputOutput `xml:"output"`
}

type InputOutput struct {
	Message string `xml:"message,attr"`
}

type Binding struct {
	Name       string             `xml:"name,attr"`
	Type       string             `xml:"type,attr"`
	Operations []OperationBinding `xml:"operation"`
}

type OperationBinding struct {
	Name       string      `xml:"name,attr"`
	SoapAction string      `xml:"soapAction,attr"`
	Style      string      `xml:"style,attr"`
	Input      BodyBinding `xml:"input"`
	Output     BodyBinding `xml:"output"`
}

type BodyBinding struct {
	Use string `xml:"use,attr"`
}

type Service struct {
	Name  string `xml:"name,attr"`
	Ports []Port `xml:"port"`
}

type Port struct {
	Name    string  `xml:"name,attr"`
	Binding string  `xml:"binding,attr"`
	Address Address `xml:"address"`
}

type Address struct {
	Location string `xml:"location,attr"`
}
