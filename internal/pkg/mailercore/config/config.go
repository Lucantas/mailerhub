package config

import (
	"encoding/json"
	"os"
)

// Config contains
type Config struct {
	DBConnections []DBConfig `json:"dbConnections"`
	WebAPIs       []WebAPI   `json:"webApis"`
}

// DBConfig is responsible to create a connection
// with the database
type DBConfig struct {
	Type          string `json:"type"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Username      string `json:"username"`
	DBName        string `json:"dbName"`
	Password      string `json:"password"`
	PlainPassword bool   `json:"plainPassword"`
	SSL           bool   `json:"ssl"`
}

// WebAPI is the struct representing a type of API
// the most important is the internal one, that will be
// used for the system communicate with its frontend service
// where the user will actually use the service
type WebAPI struct {
	Name   string `json:"name"`
	Secret `json:"secret"`
}

// Secret is a base string for building api keys,
// the period of update is not decided yet
type Secret struct {
	Current    string `json:"current"`
	Expiration string `json:"expiration"`
}

// New reads the main config file to setup the application
func New() Config {
	bytes := readJSON(os.Getenv("MAILER_SETUP"))

	var c Config

	json.Unmarshal(bytes, &c)

	return c
}
