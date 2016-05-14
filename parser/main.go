package main

import (
	"GoRepositories/Mongo"
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// 	"errors"
	// 	"fmt"
	// 	"log"
	// 	"strconv"
	// 	"strings"
	// 	"time"
	// "GoRepositories/Mongo"
	// 	"gopkg.in/mgo.v2"
	// 	"gopkg.in/mgo.v2/bson"
)

func main() {
	// connectionString := "mongodb://10.0.12.90:27117,10.0.11.124:27117/DigitalFleet/?replicaSet=replDigital;maxPoolSize=10000;w=1;readPreference=primaryPreferred;journal=true"
	connectionString := "mongodb://127.0.0.1:27017/DigitalFleet/?maxPoolSize=10000;w=1;readPreference=primaryPreferred;journal=true"

	fmt.Printf("started print\n")
	collection := Mongo.InitCollectionAndDatabaseFromConnectionString(connectionString, "Users")
	findQuery, _ := findUserQuery("Demo", "User", "demo123@digitalfleet.com", "5589532fe645110ffc4d305f")
	item := new(User)
	//	personQuery := Mongo.FindOne(collection, bson.M{})
	errQuery := Mongo.FindOne(collection, findQuery).One(&item)
	if errQuery != nil {
		fmt.Println("error FindOne", errQuery)
	}
	fmt.Println(item)
	//	db, _ := Mongo.InitDatabaseFromConnection(connectionString, "GoLangTest")
	//	collection := Mongo.InitCollectionFromConnectionString(connectionString, "GoLangTest", "people")

	// lifeCycleTest(collection)
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("")
	// get(collection, bson.M{"name": "Ale"})

}

func lifeCycleTest(collection *mgo.Collection) {
	ID1 := bson.NewObjectId().Hex()
	//	newPerson, _ := bson.Marshal(&Person{ID1, "TestUser" + ID1, "+55 53 8116 9639", ID1, time.Now().UTC()})
	newPerson := &Person{ID1, "TestUser" + ID1, "+55 53 8116 9639", "test", time.Now().UTC()}

	// create
	insertErr := Mongo.Insert(collection, newPerson)
	if insertErr != nil {
		fmt.Println(insertErr)
	}

	findByID := bson.M{"_id": ID1}

	// update
	newPerson.Phone = "1234"
	updateErr := Mongo.Update(collection, findByID, newPerson)
	if updateErr != nil {
		fmt.Println(updateErr)
	}

	// update by query
	updateQuery := bson.M{"$set": bson.M{"phone": "12345"}}
	updateinfo, updateQueryError := Mongo.UpdateByQuery(collection, findByID, updateQuery)
	if updateQueryError != nil {
		fmt.Println(updateQueryError)
	}
	fmt.Println("updateinfo : ", updateinfo)

	// update and return
	fmt.Println("UpdateOneAndReturn")
	updateReturnQuery := bson.M{"$set": bson.M{"phone": "123456"}}
	findUpdateAndReturnQuery := bson.M{"testid": "test"}

	returnItems, updatereturninfo, updateReturnQueryError := Mongo.UpdateOneAndReturn(collection, findUpdateAndReturnQuery, updateReturnQuery)
	if updateReturnQueryError != nil {
		fmt.Println(updateReturnQueryError)
	}
	fmt.Println("updatereturninfo : ", updatereturninfo)
	fmt.Println("updatereturnitem : ", returnItems, returnItems["name"])

	// deleteQuery := bson.M{"testid": "test"}
	// deleteInfo, deleteErr := Mongo.DeleteByQuery(collection, deleteQuery)
	// if deleteErr != nil {
	// 	fmt.Println(deleteErr)
	// }
	// fmt.Println("deleteInfo : ", deleteInfo)

}

func get(collection *mgo.Collection, query bson.M) {
	//	result1 := Person{}
	queryOptions := Mongo.QueryOptions{}
	queryOptions.SetQueryOptionDefaults()
	queryOptions.Skip = 0
	queryOptions.Projection = bson.M{"testid": 0}
	//queryOptions.Sort = "-datecreated"
	//	queryOptions.Sort = ""

	fmt.Println("Get all")
	// itemss, err2 := Mongo.Find(collection, bson.M{}, queryOptions)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	person := new(Person)
	//	personQuery := Mongo.FindOne(collection, bson.M{})
	errQuery := Mongo.FindOne(collection, bson.M{}).One(&person)
	if errQuery != nil {
		fmt.Println("error FindOne", errQuery)
	}

	fmt.Println("**********************************************************************")
	fmt.Println(person.ID, person.Name, person.Phone, person.TestID, person.DateCreated)
	fmt.Println("**********************************************************************")

	//for i := range itemss {
	//	name := itemss[i]["name"].(string)
	//	person := itemss[i].(Person)
	//name := asString(itemss[i], "name1")
	//		fmt.Println("**********************************************************************")
	//		fmt.Println(person.Name, itemss[i]["phone"], itemss[i]["testid"], itemss[i]["datecreated"])
	//	}
	// fmt.Println("Get first")
	// findFirst, errFirst := Mongo.FindOne(collection, bson.M{"_id": "5705b6bd1d505c08ec68d2c9j"})
	// if errFirst != nil {
	// 	//log.Fatal(errFirst)
	// 	fmt.Println(errFirst)
	// 	fmt.Println(findFirst)
	// 	if findFirst == nil {
	// 		fmt.Println("yes")
	// 	}
	// } else {
	// 	fmt.Println(findFirst["name"], findFirst["phone"], findFirst["testid"], findFirst["datecreated"])
	// }

	fmt.Println("Get byId")
	person1 := new(Person)
	errByID := Mongo.FindByID(collection, "570ea1ef1d505c12a8c380e8").One(&person1)
	if errByID != nil {
		log.Println(errByID)
	} else {
		fmt.Println("**********************************************************************")
		fmt.Println(person1.ID, person1.Name, person1.Phone, person1.TestID, person1.DateCreated)
		fmt.Println("**********************************************************************")
	}

	fmt.Println("Get Count")
	findCount, errCount := Mongo.Count(collection, bson.M{"_id": "570ea1ef1d505c12a8c380e8"})
	if errCount != nil {
		fmt.Println(errCount)
	}
	fmt.Println(findCount)

	fmt.Println("Get Distinct")
	var result []string
	findDistinct, errDistinct := Mongo.Distinct(collection, bson.M{}, "name", result)
	if errDistinct != nil {
		fmt.Println(errDistinct)
	}
	fmt.Println(findDistinct)

}

// data bson.M,
// func asString1(data bson.M, propertyNames []string) string {
// 	//var itemValue interface{}
//
// 	itemValue := data
// 	for i := range propertyNames {
// 		propertyName := itemValue[propertyName]["test"]
// 		itemValue, ok := itemValue[propertyName]["test"]
// 		if ok == false {
// 			return ""
// 		}
// 	}
// 	var test string
// 	test = itemValue.(string)
// 	value, ok1 := itemValue.(string)
//
// 	//value, ok := data[propertyName].(string)
// 	if ok1 == false {
// 		return ""
// 	}
// 	return value
// }

// func asString(data bson.M, propertyName string) string {
// 	value, ok := data[propertyName].(string)
// 	if ok == false {
// 		return ""
// 	}
// 	return value
// }

// Person test
type Person struct {
	ID          string `bson:"_id"`
	Name        string
	Phone       string
	TestID      string
	DateCreated time.Time
}

func findUserQuery(firstName string, lastName string, email string, tenantID string) (bson.M, error) {
	findQuery := bson.M{}
	if tenantID == "" {
		return findQuery, errors.New("tenantid cannot be empty")
	}
	findQuery["TenantId"] = tenantID
	if firstName != "" {
		findQuery["FirstName"] = getRegexQuery("FirstName", firstName)
	}

	if lastName != "" {
		findQuery["LastName"] = getRegexQuery("LastName", lastName)
	}
	if email != "" {
		findQuery["Email"] = getRegexQuery("Email", email)
	}
	return findQuery, nil
}

func getRegexQuery(propertyName string, propertyValue string) bson.M {
	regexQuery := bson.RegEx{}
	regexQuery.Pattern = propertyValue + "$"
	regexQuery.Options = "i"
	//return bson.M{propertyName: bson.M{"$regex": regexQuery}}
	return bson.M{"$regex": regexQuery}
}

// User model
type User struct {
	// (ReadOnly) Id of the document, created by system utilizing Mongo Bson Id
	ID string `json:"Id"  bson:"_id" binding:"required"`
	// (ReadOnly) Date the document was created (UTC)
	DateCreated time.Time `json:"DateCreated" bson:"DateCreated" binding:"required"`
	// (ReadOnly) Date the document was modified (UTC)
	DateModified time.Time `json:"DateModified" bson:"DateModified"`
	// (ReadOnly) What user modified the document
	UserModified string `json:"UserModified" bson:"UserModified,omitempty"`
	// (ReadOnly) Has the record been deleted
	IsDeleted bool `json:"IsDeleted" bson:"IsDeleted,omitempty"`
	// (Migration/Sync only)ExternalId is used to sync with external systems (Not used for anything internally)
	ExternalID string `json:"ExternalId" bson:"ExternalId,omitempty"`
	// (Migration/Sync only)ExternalNumber is used to sync with external systems (Not used for anything internally)
	ExternalNumber string `json:"ExternalNumber" bson:"ExternalNumber,omitempty"`
	// (Migration/Sync only)used to validate systems and keep them in sync (Not used for anything internally)
	HashCode string `json:"HashCode"  bson:"HashCode,omitempty"`
	// TenantID pertains to the tenant the document belongs to.  this is not serialized out
	TenantID string `json:"-" bson:"TenantId"`

	Title                   string     `json:"Title" bson:"Title"`
	FirstName               string     `json:"FirstName"  bson:"FirstName"`
	LastName                string     `json:"LastName"  bson:"LastName"`
	UserName                string     `json:"UserName"  bson:"UserName"`
	Email                   string     `json:"Email"  bson:"Email"`
	Type                    string     `json:"Type"  bson:"Type"`
	TimeZone                string     `json:"TimeZone"  bson:"TimeZone"`
	CommunicationPreference string     `json:"CommunicationPreference"  bson:"CommunicationPreference"`
	IsDriver                bool       `json:"IsDriver" bson:"IsDriver"`
	Pin                     string     `json:"Pin" bson:"Pin"`
	Department              string     `json:"Department" bson:"Department"`
	HireDate                *time.Time `json:"HireDate" bson:"HireDate"`
	CustomerID              string     `json:"Customer_Id" bson:"Customer_Id"`
	HomePlantID             string     `json:"Home_Plant_Id" bson:"Home_Plant_Id"`
	VehicleID               string     `json:"Vehicle_Id" bson:"Vehicle_Id"`
	PrimaryCustomerIds      []string   `json:"Primary_Customer_Ids" bson:"-"`
	PrimaryJobIds           []string   `json:"Primary_Job_Ids" bson:"-"`
	PrimaryOrderIds         []string   `json:"Primary_Order_Ids" bson:"-"`
	PrimaryTicketIds        []string   `json:"Primary_Ticket_Ids" bson:"-"`
}
