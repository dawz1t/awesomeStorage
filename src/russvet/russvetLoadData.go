package russvet

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"os"
)

type RussvetDocument struct {
	XMLName        xml.Name           `xml:"Document"`
	DocType        string             `xml:"DocType"`
	SenderGln      uint64             `xml:"SenderGln"`
	ReceiverGln    uint64             `xml:"ReceiverGln"`
	Currency       string             `xml:"Currency"`
	DocumentNumber uint64             `xml:"DocumentNumber"`
	DocumentDate   uint32             `xml:"DocumentDate"`
	DocDetail      []RussvetDocDetail `xml:"DocDetail"`
}

type RussvetDocDetail struct {
	XMLName         xml.Name               `xml:"DocDetail"`
	EAN             RussvetEAN             `xml:"EAN"`
	SenderPrdCode   uint64                 `xml:"SenderPrdCode"`
	ReceiverPrdCode uint64                 `xml:"ReceiverPrdCode"`
	ProductName     string                 `xml:"ProductName"`
	UOM             string                 `xml:"UOM"`
	AnalitCat       string                 `xml:"AnalitCat"`
	ParentProdCode  string                 `xml:"ParentProdCode"`
	ParentProdGroup string                 `xml:"ParentProdGroup"`
	ProductCode     string                 `xml:"ProductCode"`
	ProductGroup    string                 `xml:"ProductGroup"`
	ItemsPerUnit    float32                `xml:"ItemsPerUnit"`
	QTY             float32                `xml:"QTY"`
	SumQTY          float32                `xml:"SumQTY"`
	Price2          float32                `xml:"Price2"`
	Brand           string                 `xml:"Brand"`
	RetailPrice     float32                `xml:"RetailPrice"`
	RetailCurrency  string                 `xml:"RetailCurrency"`
	CustPrice       float32                `xml:"CustPrice"`
	VendorProdNum   string                 `xml:"VendorProdNum"`
	SupOnhandDetail RussvetSupOnhandDetail `xml:"SupOnhandDetail"`
	Multiplicity    int                    `xml:"Multiplicity"`
	QtyLots         string                 `xml:"QtyLots"`
	ItemId          uint64                 `xml:"ItemId"`
	BlockExpAll     string                 `xml:"BlockExpAll"`
	BlockExpBy      string                 `xml:"BlockExpBy"`
	BlockExpKz      string                 `xml:"BlockExpKz"`
}

type RussvetEAN struct {
	XMLName     xml.Name `xml:"EAN"`
	Value       uint64   `xml:"Value"`
	Description string   `xml:"Description"`
}

type RussvetSupOnhandDetail struct {
	XMLName     xml.Name `xml:"SupOnhandDetail"`
	PartnerQTY  float32  `xml:"PartnerQTY"`
	PartnerUOM  string   `xml:"PartnerUOM"`
	LastUpdDate uint32   `xml:"LastUpdDate"`
}

func ParseRussvet() {

	var err error
	xmlFile, err := os.Open("/russvet/siberia/PRICAT_261342_892627694.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var russvet RussvetDocument

	err = Decode(&russvet, byteValue)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(russvet.DocDetail); i++ {
		fmt.Println("Product Name: " + russvet.DocDetail[i].ProductName)
	}
}

func Decode(document *RussvetDocument, byteValue []byte) error {

	var err error

	r := bytes.NewReader([]byte(byteValue))
	d := xml.NewDecoder(r)

	d.CharsetReader = charset.NewReaderLabel
	err = d.Decode(&document)

	return err
}
