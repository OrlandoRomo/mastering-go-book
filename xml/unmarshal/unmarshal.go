package main

import (
	"encoding/xml"
	"fmt"
)

var (
	GetProductSellable         string = "getProductSellable"
	GetProduct                 string = "getProduct"
	GetMediaContent            string = "getMediaContent"
	GetInventoryLevels         string = "getInventoryLevels"
	GetFobPoints               string = "getFobPoints"
	GetConfigurationAndPricing string = "getConfigurationAndPricing"
)

const (
	// versions for the xmlns attribute
	V1_0_0 = "1.0.0"
	V1_1_0 = "1.1.0"
	V1_2_0 = "1.2.1"
	V2_0_0 = "2.0.0"

	// xmlns attribute url
	SchemaInstanceXMLNs          = "http://www.w3.org/2001/XMLSchema-instance"
	ProductDataServiceXMLNs      = "http://www.promostandards.org/WSDL/ProductDataService"
	MediaDataServiceXMLNs        = "http://www.promostandards.org/WSDL/MediaService"
	InventoryServiceXMLNs        = "http://www.promostandards.org/WSDL/InventoryService"
	InventoryXMLNs               = "http://www.promostandards.org/WSDL/Inventory"
	PricingAndConfigurationXMLNs = "http://www.promostandards.org/WSDL/PricingAndConfiguration"
	SharedObjectsXMLNs           = "SharedObjects"
)

type PreRequest struct {
	URL  string
	Body interface{}
}

var StaticRequests map[string]*PreRequest

type WsVersion struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type Customer struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type Password struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type Sellable struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type LocalizationCountry struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type LocalizationLanguage struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type ProductID struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type ProductType struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type MediaType struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type Currency struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type Fob struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type PriceType struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

type ConfigurationType struct {
	XMLName xml.Name
	Value   string `xml:",chardata" validate:"min=1"`
}

// getProductSellable operation for v1.0.0, v.2.0.0
type ProductSellableRequest struct {
	XMLName    xml.Name
	XMLNsXsi   string `xml:"xmlns:xsi,attr"`
	WsVersion  *WsVersion
	CustomerID *Customer
	Password   *Password
	IsSellable *Sellable
}

// getProduct operation for v1.0.0, v.2.0.0
type ProductRequest struct {
	XMLName              xml.Name
	XMLNsXsi             string                `xml:"xmlns:xsi,attr"`
	WsVersion            *WsVersion            `xml:",omitempty"`
	CustomerID           *Customer             `xml:",omitempty"`
	Password             *Password             `xml:",omitempty"`
	LocalizationCountry  *LocalizationCountry  `xml:",omitempty"`
	LocalizationLanguage *LocalizationLanguage `xml:",omitempty"`
	ProductID            *ProductID            `xml:",omitempty"`
}

// getMediaContent operation for v.1.1.0
type MediaContentRequest struct {
	XMLName    xml.Name
	XMLNsXsi   string     `xml:"xmlns:xsi,attr"`
	WsVersion  *WsVersion `xml:",omitempty"`
	CustomerID *Customer  `xml:",omitempty"`
	Password   *Password  `xml:",omitempty"`
	MediaType  *MediaType `xml:",omitempty"`
	ProductID  *ProductID `xml:",omitempty"`
}

// getInventoryLevels operation for v1.0.0, v.1.2.1 and v2.0.0
type InventoryLevelsRequest struct {
	XMLName       xml.Name
	XMLNsXsi      string       `xml:"xmlns:xsi,attr"`
	WsVersion     *WsVersion   `xml:",omitempty"`
	CustomerID    *Customer    `xml:",omitempty"`
	Password      *Password    `xml:",omitempty"`
	ProductID     *ProductID   `xml:",omitempty"`
	ProductIDtype *ProductType `xml:",omitempty"`
}

// getFobPoints operation for v1.0.0
type FobPointsRequest struct {
	XMLName              xml.Name
	XMLNsXsi             string                `xml:"xmlns:xsi,attr"`
	WsVersion            *WsVersion            `xml:",omitempty"`
	CustomerID           *Customer             `xml:",omitempty"`
	Password             *Password             `xml:",omitempty"`
	ProductID            *ProductID            `xml:",omitempty"`
	LocalizationCountry  *LocalizationCountry  `xml:",omitempty"`
	LocalizationLanguage *LocalizationLanguage `xml:",omitempty"`
}

// getConfigurationAndPricing operation for v1.0.0
type ConfigurationAndPricingRequest struct {
	XMLName              xml.Name
	XMLNsXsi             string                `xml:"xmlns:xsi,attr"`
	WsVersion            *WsVersion            `xml:",omitempty"`
	CustomerID           *Customer             `xml:",omitempty"`
	Password             *Password             `xml:",omitempty"`
	ProductID            *ProductID            `xml:",omitempty"`
	Currency             *Currency             `xml:",omitempty"`
	FobID                *Fob                  `xml:",omitempty"`
	PriceType            *PriceType            `xml:",omitempty"`
	LocalizationCountry  *LocalizationCountry  `xml:",omitempty"`
	LocalizationLanguage *LocalizationLanguage `xml:",omitempty"`
	ConfigurationType    *ConfigurationType    `xml:",omitempty"`
}

func main() {
	productSellable := &ProductSellableRequest{
		XMLName: xml.Name{
			Local: "GetProductSellable",
			Space: fmt.Sprintf("%s/%s/", InventoryServiceXMLNs, V2_0_0),
		},
		XMLNsXsi: SchemaInstanceXMLNs,

		WsVersion: &WsVersion{
			XMLName: xml.Name{
				Local: "wsVersion",
				Space: fmt.Sprintf("%s/%s/%s/", InventoryServiceXMLNs, V2_0_0, SharedObjectsXMLNs),
			},
			Value: "Token1",
		},
		CustomerID: &Customer{
			XMLName: xml.Name{
				Local: "id",
				Space: fmt.Sprintf("%s/%s/%s/", InventoryServiceXMLNs, V2_0_0, SharedObjectsXMLNs),
			},
			Value: "Token1",
		},
		Password: &Password{
			XMLName: xml.Name{
				Local: "password",
				Space: fmt.Sprintf("%s/%s/%s/", InventoryServiceXMLNs, V2_0_0, SharedObjectsXMLNs),
			},
			Value: "$%&$%&",
		},
		IsSellable: &Sellable{
			XMLName: xml.Name{
				Local: "isSellable",
				Space: fmt.Sprintf("%s/%s/%s/", InventoryServiceXMLNs, V2_0_0, SharedObjectsXMLNs),
			},
		},
	}
	raw, err := xml.MarshalIndent(productSellable, "", "\t")
	fmt.Println(string(raw), err)
}
