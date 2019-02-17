package mailerdb

import (
	"log"
	"mailer-service/internal/pkg/mailercore/components"
)

func GetUsersAddressLists(id int) (components.AddressList, error) {
	dbConn := DBConn()
	dbConn.LogMode(true)
	defer dbConn.Close()

	var aL components.AddressList
	log.Println(id)
	dbConn.First(&aL, id)

	return aL, nil
}

func saveAddressList(list components.AddressList, name string) error {
	dbConn := DBConn()
	dbConn.LogMode(true)
	defer dbConn.Close()

	dbConn.AutoMigrate(&components.AddressList{})
	err := dbConn.Save(&list).GetErrors()

	if len(err) > 0 {
		return err[0]
	}
	return nil
}
