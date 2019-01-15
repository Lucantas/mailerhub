package main

import (
	"fmt"
	"mailer-service/internal/pkg/mailercore/config"
)

type screen struct {
	Config config.Config
}

func (s screen) currentConfig() {
	s.Config = config.New()

	fmt.Println("Current mailer configuration:")
	fmt.Println("DBConnections : ", len(s.Config.DBConnections))
	fmt.Println("APIConfigs : ", len(s.Config.WebAPIs))

	if len(s.Config.DBConnections) > 0 {
		for i, dbConn := range s.Config.DBConnections {
			fmt.Println("DB Connection ", i)
			fmt.Println("Type: ", dbConn.Type)
			fmt.Println("Host: ", dbConn.Host)
			fmt.Println("Port: ", dbConn.Port)
			fmt.Println("Username: ", dbConn.Username)
			fmt.Println("Database name: ", dbConn.DBName)
			fmt.Println("Password: ", dbConn.Password)
			fmt.Println("PlainPassword: ", dbConn.PlainPassword)
			fmt.Println("SSL: ", dbConn.SSL)
		}
	}

	if len(s.Config.WebAPIs) > 0 {
		for i, api := range s.Config.WebAPIs {
			fmt.Println("Web Api ", i)
			fmt.Println("Name: ", api.Name)
			fmt.Println("Secret: ", api.Secret.Current)
			fmt.Println("Expiration: ", api.Secret.Expiration)
		}
	}
}

func (s screen) showOptions() {
	fmt.Println("OPTIONS: ")
	fmt.Println("1 - Create Configuration")
	fmt.Println("2 - Edit existing configuration")
	fmt.Println("3 - Exit")
}

func (s screen) handleOptions(option int) {
	switch option {
	case 1:
		s.createConfiguration()
	case 2:
		s.editConfiguration()
	case 3:
		s.handleExit()
	}
}

func (s screen) createConfiguration() {

}

func (s screen) editConfiguration() {

}

func (s screen) handleExit() {}
