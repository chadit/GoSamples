package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

// Employees is a collection of employee
type Employees struct {
	employee []Employee
}

type Employees1 []Employee

// Employee model for DigitalFleet database
type Employee struct {
	EmployeeID int
	ClientID   int
	LastName   string `sql:"Last_Name"`
	FirstName  string `sql:"First_Name"`
	Login      string
	PIN        string
}

var debug = flag.Bool("debug", true, "enable debugging")
var sqlPassword = flag.String("password", "ky1ttk#1", "the database password")
var port = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "192.168.0.52", "the database server")
var sqlUser = flag.String("user", "chad", "the database user")
var database = flag.String("dbname", "DigitalFleet", "Database name")

func main() {

	//connString := "server=192.168.0.52;user id=chad;password=ky1ttk#1;database=DigitalFleet;port=1433"
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d", *server, *sqlUser, *sqlPassword, *database, *port)

	fmt.Printf(" connString:%s\n", connString)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	// Is the database running?
	strResult := PingServer(db)
	fmt.Println("(sqltest) Ping of Server Result Was: " + strResult)
	fmt.Println("********************************")

	// Does the database exist?
	boolDBExist, err := CheckDB(db, *database)
	if err != nil {
		fmt.Println("(sqltest) Error running CheckDB: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("(sqltest) Database Existence Check: " + strconv.FormatBool(boolDBExist))
	fmt.Println("********************************")

	// select EmployeeID,ClientID,Last_Name,First_Name, LOGIN, PIN from employees where ClientID = '18'
	//rows, err := db.Query("select EmployeeID,ClientID,Last_Name,First_Name, LOGIN, PIN from employees where PIN = '440440'")
	rows, err := db.Query("select EmployeeID,ClientID,Last_Name,First_Name, LOGIN, PIN from employees where ClientID = '18'")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	c := new(Employees)
	d := Employees1{}

	for rows.Next() {
		m := Employee{}
		err := rows.Scan(&m.EmployeeID, &m.ClientID, &m.LastName, &m.FirstName, &m.Login, &m.PIN)
		if err != nil {
			panic("Error reading row: " + err.Error())
		}
		fmt.Println(&m)
		d = append(d, m)
		c.employee = append(c.employee, m)
	}
	fmt.Println("********************************")
	count := len(d)
	fmt.Printf("items found %d\n ", count)
	for i := 0; i < count; i++ {
		fmt.Println(d[i])
	}

}

// PingServer uses a passed database handle to check if the database server works
func PingServer(db *sql.DB) string {

	err := db.Ping()
	if err != nil {
		return ("From Ping() Attempt: " + err.Error())
	}

	return ("Database Ping Worked...")

}

// CheckDB checks if the database "strDBName" exists on the MSSQL database engine.
func CheckDB(db *sql.DB, strDBName string) (bool, error) {

	// Does the database exist?
	result, err := db.Query("SELECT db_id('" + strDBName + "')")
	defer result.Close()
	if err != nil {
		return false, err
	}

	for result.Next() {
		var s sql.NullString
		err := result.Scan(&s)
		if err != nil {
			return false, err
		}

		// Check result
		if s.Valid {
			return true, nil
		}
		return false, nil
	}

	// This return() should never be hit...
	return false, err
}
