package main

import (
	"mailer-service/internal/pkg/mailerdb"
)

func main() {
	config := newConfig()

	dbConfig := config.DBConnections[0]

	gormConn := mailerdb.DBConn(
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.DBName,
		dbConfig.Password,
		dbConfig.PlainPassword,
		dbConfig.SSL,
	)

	mailerdb.SaveClient(
		gormConn,
		"admin@mailercloud.com",
		"123456",
	)
}
