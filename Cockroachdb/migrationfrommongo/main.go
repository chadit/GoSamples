package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/lib/pq"
)

var (
	// ErrNilDB indicates that DB is nil.
	ErrNilDB = errors.New("mc: DB cannot be nil")
	// ErrEmptyURL indicates that URL is empty.
	ErrEmptyURL = errors.New("mc: URL is empty")
	// ErrEmptySchema indicates that schema is empty.
	ErrEmptySchema = errors.New("mc: schema is empty")
	// ErrNilDBScript indicates that DBScript is nil.
	ErrNilDBScript = errors.New("mc: DBScript cannot be nil")
	// ErrSchemaExists indicates that the database already exists.
	ErrSchemaExists = errors.New("pq: schema %q already exists")
)

func main() {
	var err error
	d := getDbSettings()
	if err = d.initDatabases(); err != nil {
		fmt.Println("initDatabases : ", err)
	}

	// db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/bank?sslmode=disable")
	// if err != nil {
	// 	log.Fatalf("error connection to the database: %s", err)
	// }

}

func (db *DB) initDatabases() error {
	parsedURL, err := url.Parse(db.URL)
	if err != nil {
		return err
	}

	db.DB, err = sql.Open("postgres", parsedURL.String())
	if err != nil {
		return err
	}
	db.name = strings.TrimPrefix(parsedURL.Path, "/")
	db.user = parsedURL.User.Username()
	fmt.Println("db.name : ", db.name)
	fmt.Println("db.user : ", db.user)

	//db.databaseSetup()

	if _, err = db.Exec(fmt.Sprintf("CREATE SCHEMA %s AUTHORIZATION %s", db.Schema, db.user)); err != nil {
		if err.Error() == fmt.Sprintf(ErrSchemaExists.Error(), db.Schema) {
			fmt.Println("ErrSchemaExists : ", err)
			//log.Println(err)
		} else {
			return err
		}
	}

	for _, script := range db.Script {
		if _, err = db.Exec(strings.Replace(script, "{{DBNAME}}.{{DBSCHEMA}}", fmt.Sprint(db.name, ".", db.Schema), -1)); err != nil {
			//if _, err = db.Exec(strings.Replace(script, "{{DBNAME}}", db.name, -1)); err != nil {
			return err
		}
	}

	return nil
}

// func (db *DB) databaseSetup() error {
// 	rootConn := strings.Replace(db.URL, db.user, "root", 1)
// 	parsedURL, err := url.Parse(rootConn)
// 	if err != nil {
// 		fmt.Println("parse error")
// 		return err
// 	}

// 	dbroot, err := sql.Open("postgres", parsedURL.String())
// 	if err != nil {
// 		fmt.Println("parse open db err")
// 		return err
// 	}

// 	if _, err := dbroot.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db.name)); err != nil {
// 		fmt.Println("exec create db error")
// 		return err
// 	}
// 	return nil
// }

// DB holds all methods necessary for SQL database logic.
type DB struct {
	URL    string
	Script DBScript
	Schema string
	name   string
	user   string
	*sql.DB
}

func getDbSettings() DB {
	return DB{
		Script: DBTables,
		URL:    "postgres://postgres:postgres@127.0.0.1:5432/qiktracker?sslmode=disable",
		//URL:    "postgresql://maxroach@localhost:26257/QikTracker?sslmode=disable",
		//	Name:   "QikTracker",
		Schema: "qik",
	}
}

// DBScript holds table collection information
type DBScript map[string]string

// DBTables holds scripts for the table database.
var DBTables = DBScript{
	"ProductTrackings": `CREATE TABLE IF NOT EXISTS {{DBNAME}}."ProductTracking"
(
  ID serial primary key NOT NULL,
  DateCreated timestamp,
  MarketplaceID varchar(15) NOT NULL,
  Asin char(10),
  Upc varchar(15) NOT NULL,
  Domain varchar(100) NOT NULL,
  Title varchar(1000) NOT NULL,
  Category varchar(50) NOT NULL,
  Author varchar(100),
  Condition varchar(25) NOT NULL,
  SubCondition varchar(25),
  PathName varchar NOT NULL,
  ImageURL varchar,
  CurrencyCode varchar(10),
  RegularAmount float NOT NULL,
  SaleAmount float NOT NULL,
  ShippingAmount float NOT NULL,
  TotalAmount float NOT NULL,
  SalesRank int,
  SellerFeedbackCount int,
  SellerPositiveFeedbackRating varchar(25),
  Count int,
  Channel varchar(10) NOT NULL,
  IsSoldByAmazon bool NOT NULL,
  IsBuyBoxEligible bool NOT NULL,
  PackageLength float NOT NULL,
  PackageWidth float NOT NULL,
  PackageHeight float NOT NULL,
  PackageWeight float NOT NULL,
  AmazonFees float NOT NULL,
  RetryCount int NOT NULL,
  Message varchar(1000)
);`,
}

// var DBTables = DBScript{
// 	"org": `CREATE TABLE IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."org"
// (
//   id serial primary key NOT NULL,
//   "name" character varying(255) CHECK (name !=''),
//   "createdat" timestamp,
//   "createdby" integer,
//   "updatedat" timestamp,
//   "updatedby" integer
// );`,
// 	"user": `CREATE TABLE IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."user"
// (
//   id serial primary key NOT NULL,
//   "orgid" integer NOT NULL,
//   "firstname" character varying(255) CHECK ("firstname" !=''),
//   "lastname" character varying(255) CHECK ("lastname" !=''),
//   email character varying(255) CHECK ("email" !=''),
//   "jobtitle" character varying(255),
//   "isactive" boolean,
//   "createdat" timestamp,
//   "createdby" integer,
//   "updatedat" timestamp,
//   "updatedby" integer
// );`,
// 	"serviceuser": `CREATE TABLE IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."serviceuser"
// (
//   id serial primary key NOT NULL,
//   "userid" integer NOT NULL,
//   "key" character varying(255) CHECK ("key" !=''),
//   "secret" character varying(255) CHECK ("secret" !=''),
//   "isactive" boolean,
//   "createdat" timestamp,
//   "createdby" integer,
//   "updatedat" timestamp,
//   "updatedby" integer
// );`,
// 	"identity": `CREATE TABLE IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."identity"
// (
//   id serial primary key NOT NULL,
//   "userid" integer NOT NULL,
//   "uid" character varying(255) CHECK (uid !=''),
//   "provider" character varying(255) CHECK (provider !=''),
//   "token" text CHECK (token !=''),
//   "secret" character varying(2048),
//   "expiry" timestamp NOT NULL
// );`,
// 	"service": `CREATE TABLE IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."service"
// (
//   uuid character varying(36) primary key NOT NULL,
//   "name" character varying(255) CHECK (name !=''),
//   "createdat" timestamp,
//   "createdby" integer,
//   "updatedat" timestamp,
//   "updatedby" integer
// );
// `,
// 	"servicepage": `CREATE TABLE IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."servicepage"
// (
//   uuid character varying(36) primary key NOT NULL,
//   serviceuuid character varying(36) CHECK (serviceuuid !=''),
//   parentuuid character varying(36) CHECK (parentuuid !=''),
//   "title" character varying(255) CHECK (title !=''),
//   "uri" character varying(2048) CHECK (uri !=''),
//   "createdat" timestamp,
//   "createdby" integer,
//   "updatedat" timestamp,
//   "updatedby" integer);
// `,
// 	"userservicepageperm": `CREATE TABLE  IF NOT EXISTS {{DBNAME}}.{{DBSCHEMA}}."userservicepageperm"
// (
//   id serial primary key NOT NULL,
//   userid integer NOT NULL,
//   servicepageid character varying(36) CHECK (servicepageid !=''),
//   canread boolean,
//   canupdate boolean,
//   cancreate boolean,
//   candelete boolean,
//   "createdat" timestamp,
//   "createdby" integer,
//   "updatedat" timestamp,
//   "updatedby" integer
// );`,
// }
