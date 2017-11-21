package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Person test
type Person struct {
	Name        string
	Phone       string
	TestID      string
	DateCreated time.Time
}

// Order test
type Order struct {
	Address struct {
		City        string `json:"City"`
		Country     string `json:"Country"`
		CountryCode string `json:"CountryCode"`
		County      string `json:"County"`
		Gps         struct {
			Radius   float64   `json:"Radius"`
			Bbox     []float64 `json:"bbox"`
			Geometry struct {
				Coordinates []float64 `json:"coordinates"`
				Type        string    `json:"type"`
			} `json:"geometry"`
			Type string `json:"type"`
		} `json:"Gps"`
		PostCode      string `json:"PostCode"`
		StateProvince string `json:"StateProvince"`
		Street        string `json:"Street"`
	} `json:"Address"`
	CancelReason    string    `json:"CancelReason"`
	Category        string    `json:"Category"`
	CustomerID      string    `json:"Customer_Id"`
	DateClosed      time.Time `json:"DateClosed"`
	DateCompleted   time.Time `json:"DateCompleted"`
	DateCreated     time.Time `json:"DateCreated"`
	DateModified    time.Time `json:"DateModified"`
	DateToShip      time.Time `json:"DateToShip"`
	ExternalID      string    `json:"ExternalId"`
	ExternalNumber  string    `json:"ExternalNumber"`
	HashCode        string    `json:"HashCode"`
	IsCancelled     bool      `json:"IsCancelled"`
	IsDeleted       bool      `json:"IsDeleted"`
	JobSiteLocation struct {
		City        string `json:"City"`
		Country     string `json:"Country"`
		CountryCode string `json:"CountryCode"`
		County      string `json:"County"`
		Gps         struct {
			Radius   float64   `json:"Radius"`
			Bbox     []float64 `json:"bbox"`
			Geometry struct {
				Coordinates []float64 `json:"coordinates"`
				Type        string    `json:"type"`
			} `json:"geometry"`
			Type string `json:"type"`
		} `json:"Gps"`
		PostCode      string `json:"PostCode"`
		StateProvince string `json:"StateProvince"`
		Street        string `json:"Street"`
	} `json:"JobSiteLocation"`
	JobID     string `json:"Job_Id"`
	LineItems []struct {
		Category          string  `json:"Category"`
		Code              string  `json:"Code"`
		DeliveredQuantity float64 `json:"DeliveredQuantity"`
		Description       string  `json:"Description"`
		LoadSize          string  `json:"LoadSize"`
		Name              string  `json:"Name"`
		Price             float64 `json:"Price"`
		Quantity          float64 `json:"Quantity"`
		Slump             float64 `json:"Slump"`
		//	SubLineItems      interface{} `json:"SubLineItems"`
		TaxCode       string `json:"TaxCode"`
		UnitOfMeasure string `json:"UnitOfMeasure"`
		ID            string `json:"_id"`
	} `json:"LineItems"`
	Loads               float64  `json:"Loads"`
	MiscNotes           string   `json:"MiscNotes"`
	Note                string   `json:"Note"`
	OrderID             string   `json:"OrderId"`
	OrderNumber         string   `json:"OrderNumber"`
	OrderTicketIds      []string `json:"Order_Ticket_Ids"`
	PrimaryUserID       string   `json:"Primary_User_Id"`
	PurchaseOrderNumber string   `json:"PurchaseOrderNumber"`
	TenantID            string   `json:"TenantId"`
	Type                string   `json:"Type"`
	UserModified        string   `json:"UserModified"`
	VehicleSpacing      float64  `json:"VehicleSpacing"`
	ID                  string   `json:"_id"`
}

const databaseName = "GoLangTest"

func main() {
	//complexTest()
	basicTest()
}

func complexTest() {
	//session, err := mgo.Dial("mongodb://127.0.0.1/DigitalFleet/?w=majority;journal=true;maxPoolSize=1000")

	session, err := mgo.Dial("mongodb://127.0.0.1/")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// http://stackoverflow.com/questions/28775929/using-mgo-with-nested-documents-that-convert-to-mapstringinterface-how-to-i

	c := session.DB("DigitalFleet").C("Orders")

	var result bson.M
	//result := Order{}
	err = c.Find(bson.M{"_id": "5663a6a757ea3d28e4fdee04"}).One(&result)
	//fmt.Println("Id", result.ID)
	fmt.Println(result)
	result = addProperty(result)
	fmt.Println(result)

	newOrder := new(Order)
	newOrder.ID = bson.NewObjectId().Hex()
	testOrder, _ := json.Marshal(newOrder)
	fmt.Println(testOrder)
	//testOrder["test2"] = "testtest"
	//result["test"]
}

func addProperty(item bson.M) bson.M {
	item["test"] = "testdata"
	return item
}

func basicTest() {
	session, err := mgo.Dial("mongodb://127.0.0.1/")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(databaseName).C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639", bson.NewObjectId().Hex(), time.Now().UTC()},
		&Person{"Cla", "+55 53 8402 8510", bson.NewObjectId().Hex(), time.Now().UTC()})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
	session.Close()
	fmt.Println("2")
	session.Close()
}
