package dbTokenControl

import (
	"database/sql"
	"fmt"
	"os"
)

var sconn StringConnection

func stringConnection() string {
	sConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", sconn.DbUser, sconn.DbPassword, sconn.DbHost, sconn.DbPort, sconn.DbName)
	return sConnection
}

func AddMessage(message string) {

	if message == "" {
		return
	}

	db, err := OpenConnection()
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO log(message) values(?)")
	if err != nil {
		panic(err)
	}

	stmt.Exec(message)
	defer db.Close()

}

func OpenConnection() (*sql.DB, error) {

	strConn := stringConnection()
	db, err := sql.Open("mysql", strConn)
	checkErr(err)

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("database connection failed | " + strConn)
		db.Close()
		os.Exit(3)
	}

	return db, err
}
