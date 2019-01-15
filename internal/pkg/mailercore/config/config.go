package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// TODO: parse the config as xml

// Config contains
type Config struct {
	DBConnections []DBConfig `xml:"config>db-connections>db-connection"`
}

// DBConfig is responsible to create a connection
// with the database
type DBConfig struct {
	XMLName       xml.Name `xml:"db-connection"`
	Type          string   `xml:"type,attr"`
	Host          string   `xml:"host"`
	Port          string   `xml:"port"`
	Username      string   `xml:"username"`
	DBName        string   `xml:"db-name"`
	Password      string   `xml:"password"`
	PlainPassword bool     `xml:"plain-password"`
	SSL           bool     `xml:"use-ssl"`
}

// New reads the main config file to setup the application
func New() Config {
	file, err := os.Open(os.Getenv("MAILER_SETUP"))

	if err != nil {
		// TODO: properly handle if error on this part
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var c Config

	xml.Unmarshal(byteValue, &c)

	return c
}
