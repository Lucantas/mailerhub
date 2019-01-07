package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Config contains
type Config struct {
	XMLName       xml.Name   `xml:"db-connections"`
	DBConnections []DBConfig `xml:"db-connection"`
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

func newConfig() Config {
	fmt.Println(os.Getenv("MAILER_SETUP"))
	file, err := os.Open(os.Getenv("MAILER_SETUP"))

	if err != nil {
		// TODO: properly handle if error on this part
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	fmt.Println(byteValue)

	var config Config

	xml.Unmarshal(byteValue, &config)

	return config
}
