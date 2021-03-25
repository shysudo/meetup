package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func InitDb() {
	service := readFromfile()
	driverName := "logs"
	var err error
	sql.Register(driverName, &mysql.MySQLDriver{})
	// Create the connection string
	connStr := service.Id + ":" + service.Password + "@tcp(" + service.Server + ":" + service.Port + ")/" + "meetup"
	// Connect to the db
	DB, err = sql.Open(driverName, connStr)
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Println("Database Connection Established")
	}
	// Set the connection parameters
	DB.SetMaxOpenConns(60)
	DB.SetMaxIdleConns(0)
	DB.SetConnMaxLifetime(time.Duration(time.Second * 20))
}

type Service struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Id       string `json:"id"`
	Password string `json:"password"`
}

func readFromfile() Service {
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var service Service

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &service)
	return service
}
