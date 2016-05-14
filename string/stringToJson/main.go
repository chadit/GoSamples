package main

import (
	"encoding/json"
	"fmt"
)

// QueueData model
type QueueData struct {
	Payload      string `json:"Payload" bson:"Payload"`
	PayloadType  string `json:"PayloadType"  bson:"PayloadType"`
	Tag          string `json:"Tag"  bson:"Tag"`
	HardwareType string `json:"HardwareType"  bson:"HardwareType"`
}

func main() {
	convertStringToStruct()
	fmt.Println("")
	fmt.Println("")
	convertStructToString()
}

func convertStringToStruct() {
	payload := "{\"Payload\":\"IVBOICAgMmM2MDBjZmFiOGU3MzY2MTMwNDE2MjA1OTIy\",\"PayloadType\":\"DigiTrack\",\"Tag\":\"!PN   2c600cfab8e7366130416205922\",\"HardwareType\":\"tablet\"}"
	fmt.Println(payload)
	var e QueueData
	err := json.Unmarshal([]byte(payload), &e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(e.Payload)
}

func convertStructToString() {
	items := QueueData{}
	items.Payload = "IVBOICAgMmM2MDBjZmFiOGU3MzY2MTMwNDE2MjA1OTIy"
	items.PayloadType = "DigiTrack"
	items.Tag = "!PN   2c600cfab8e7366130416205922"
	items.HardwareType = "tablet"
	byte2, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byte2))
}
