package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type GetProductSellableRequest struct {
	XMLName    xml.Name `xml:"GetProductSellableRequest"`
	XMLNS      string   `xml:"xmlns,attr"`
	WSVersion  WSVersion
	CustomerID CustomerID
	Password   Password
	IsSellable IsSellable
}

type WSVersion struct {
	WSVersion string `xml:"wsVersion" validate:"min=1"`
	XMLNS     string `xml:"xmlns,attr"`
}
type CustomerID struct {
	CustomerID string `xml:"id" validate:"min=1"`
	XMLNS      string `xml:"xmlns,attr"`
}
type Password struct {
	Password string `xml:"password" validate:"min=1"`
	XMLNS    string `xml:"xmlns,attr"`
}
type IsSellable struct {
	IsSellable bool   `xml:"isSellable" validate:"min=1"`
	XMLNS      string `xml:"xmlns,attr"`
}

func main() {
	request := &GetProductSellableRequest{}
	output, err := xml.MarshalIndent(request, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))

}
