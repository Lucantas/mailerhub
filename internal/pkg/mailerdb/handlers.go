package mailerdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// dialect to use postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConn returns a db with connection to gorm database
func DBConn(
	host string,
	port string,
	username string,
	DBName string,
	password string,
	PlainPassword bool,
	SSL bool,
) *gorm.DB {
	c := conn{
		Host:     host,
		Port:     port,
		User:     username,
		DBName:   DBName,
		Password: password,
		SSL:      SSL,
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password))
	if err != nil {
		panic(err)
	}
	return db
}

// SaveClient uses a connection with db to create a client
func SaveClient(dbConn *gorm.DB, mail string, password string) {
	defer dbConn.Close()
	dbConn.AutoMigrate(&client{})
	dbConn.Create(client{SenderEmail: "admin@mailercloud.com", Password: "123456"})
}
