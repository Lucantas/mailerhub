package mailerdb

import (
	"fmt"
	"log"
	"mailer-service/internal/pkg/mailercore/components"
	"mailer-service/internal/pkg/mailercore/config"

	"github.com/jinzhu/gorm"
	// dialect to use postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConn returns a db with connection to gorm database
func DBConn() *gorm.DB {

	dbConfig := config.New().DBConnections[0]

	c := conn{
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		User:     dbConfig.Username,
		DBName:   dbConfig.DBName,
		Password: dbConfig.Password,
		SSL:      dbConfig.SSL,
	}
	db, err := gorm.Open("postgres", c.toString())
	if err != nil {
		panic(err)
	}
	db.Set("gorm:table_options", "charset=utf8")
	return db
}

func SaveUser(u components.User) error {
	dbConn := DBConn()
	defer dbConn.Close()
	dbConn.LogMode(true)

	dbConn.AutoMigrate(&components.User{})
	err := dbConn.Save(&u).GetErrors()

	if len(err) > 0 {
		return err[0]
	}
	return nil
}

// SaveClient uses a connection with db to create a client
func SaveClient(c components.Client) {
	dbConn := DBConn()
	defer dbConn.Close()

	dbConn.AutoMigrate(&components.Client{})
	dbConn.Create(&c)
}

func saveCampaign(c components.Campaign) {
	dbConn := DBConn()
	defer dbConn.Close()

	dbConn.AutoMigrate(&components.Campaign{})
	dbConn.Save(c)
}

func saveContactForm(c components.Contact) {
	dbConn := DBConn()
	defer dbConn.Close()

	dbConn.AutoMigrate(&components.Contact{})
	dbConn.Save(c)
}

// GetClient function internally run a getClient function
// that retrieve a client based on an id
func GetClient(id int) components.Client {
	return getClient(id)
}

func getClient(id int) components.Client {
	dbConn := DBConn()
	defer dbConn.Close()

	var c components.Client

	dbConn.First(&components.Client{}, id)

	return c
}

// GetUser function internally run a getUser function
// that retrieve a user based on an id
func GetUser(id int) (components.User, error) {
	return getUser(id)
}

func GetUserByLogin(login string) (components.User, error) {
	return getUserByLogin(login)
}

func getUser(id int) (components.User, error) {
	dbConn := DBConn()
	dbConn.LogMode(true)
	defer dbConn.Close()

	var u components.User
	log.Println(id)
	dbConn.First(&u, id)

	if u.ID == 0 {
		return u, newErrUserNotFound(
			fmt.Sprintf("Could not find a user by id %d", id),
		)
	}

	return u, nil
}

func getUserByLogin(login string) (components.User, error) {
	db := DBConn()
	defer db.Close()

	var u components.User
	db.Where("email = ?", login).First(&u)

	if u.ID == 0 {
		return u, newErrUserNotFound(
			fmt.Sprintf("Could not find a user by login %s", login),
		)
	}

	return u, nil
}

/* func getContactService(id int) mailer {
	dbConn := DBConn()
	defer dbConn.Close()

	/* switch v := dbContact.(type) {
	case Campaign:
		contact = v
	}
} */

func (c conn) toString() string {
	var sslEnabled string

	if c.SSL {
		sslEnabled = "enable"
	} else {
		sslEnabled = "disable"
	}

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password, sslEnabled)
}
