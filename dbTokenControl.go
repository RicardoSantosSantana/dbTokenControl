package dbTokenControl

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/*
	  type STRConn struct {
		DbName     string
		DbHost     string
		DbUser     string
		DbPassword string
		DbPort     string
	}
*/
var StringConnection STRConn

// will save message on the database
func LogMessage(message string) {

	if message == "" {
		return
	}

	db, err := openConnection()
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO log(message) values(?)")
	if err != nil {
		panic(err)
	}

	stmt.Exec(message)
	defer db.Close()

}

func (token Token) Add() error {

	db, err := openConnection()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO config_token(access_token, token_type, expires_in, scope, user_id, refresh_token) values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	stmt.Exec(token.Access_token, token.Token_type, token.Expires_in, token.Scope, token.User_id, token.Refresh_token)
	defer db.Close()

	return nil
}

func Active() (Token, error) {

	db, err := openConnection()
	if err != nil {
		return Token{}, err
	}

	token := Token{}
	sql := "SELECT access_token, token_type, expires_in, scope, user_id, refresh_token FROM config_token ORDER BY id DESC LIMIT 1"
	err = db.QueryRow(sql).Scan(&token.Access_token, &token.Token_type, &token.Expires_in, &token.Scope, &token.User_id, &token.Refresh_token)

	if err != nil {
		return Token{}, err
	}

	defer db.Close()

	return token, nil

}

func openConnection() (*sql.DB, error) {

	strConn := stringConnection()

	db, err := sql.Open("mysql", strConn)
	if err != nil {
		return nil, err
	}

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("database connection failed > " + strConn)
		db.Close()
		os.Exit(3)
	}

	return db, err
}

func stringConnection() string {

	sConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		StringConnection.DbUser,
		StringConnection.DbPassword,
		StringConnection.DbHost,
		StringConnection.DbPort,
		StringConnection.DbName)

	return sConnection
}

func TruncateAllItems() error {

	db, errConn := openConnection()
	if errConn != nil {
		return errConn
	}

	stmt, err := db.Prepare("TRUNCATE TABLE product")
	if err != nil {
		fmt.Println(err.Error())
	}

	stmt.Exec()

	defer db.Close()

	return nil

}

func Save(item Items) error {

	db, errConn := openConnection()
	if errConn != nil {
		return errConn
	}

	stmt, err := db.Prepare("INSERT INTO product(id, title, subtitle, price, base_price, original_price, permalink, thumbnail, pictures,description) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	// sqlv := fmt.Sprintf("INSERT INTO product(id, title, subtitle, price, base_price, original_price, permalink, thumbnail, pictures,description) values('%s','%s','%s',%f,%f,%f,'%s','%s','%s','%s')", item.Id,
	// 	item.Title,
	// 	item.Subtitle,
	// 	item.Price,
	// 	item.Base_price,
	// 	item.Original_price,
	// 	item.Permalink,
	// 	item.Thumbnail,
	// 	item.Pictures,
	// 	item.Description)

	// fmt.Println(sqlv)

	stmt.Exec(
		item.Id,
		item.Title,
		item.Subtitle,
		item.Price,
		item.Base_price,
		item.Original_price,
		item.Permalink,
		item.Thumbnail,
		item.Pictures,
		item.Description)

	defer db.Close()

	return nil
}
