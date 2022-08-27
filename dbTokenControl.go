package dbTokenControl

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

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

	stmt, err := db.Prepare("INSERT INTO tokens(access_token, token_type, expires_in, scope, user_id, refresh_token) values(?,?,?,?,?,?)")
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
	sql := "SELECT access_token, token_type, expires_in, scope, user_id, refresh_token FROM tokens ORDER BY id DESC LIMIT 1"
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
		fmt.Println(strConn)
	}

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("database connection failed, trying new connection.")
		fmt.Println(strConn)

		var i int = 5		
				
		for {

			log:=fmt.Sprintf("connection %d of %d",i,5)
			fmt.Println(log)

			time.Sleep(10 * time.Second)
			i--;
			errPing2 := db.Ping()
	
			if errPing2 == nil {
				return db, nil 
			}

			if i==0 {				
				db.Close()				
				defer os.Exit(3)
				return nil, err				
			}

		}
		
		

		
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

	stmt, err := db.Prepare("TRUNCATE TABLE products")
	if err != nil {
		fmt.Println(err.Error())
	}

	stmt.Exec()

	defer db.Close()

	return nil

}

func MakeSelectStatement(table string) string {

	item := reflect.TypeOf(Items{})

	names := make([]string, item.NumField())
	ids := make([]string, item.NumField())

	for i := range names {
		name := strings.ToLower(item.Field(i).Name)

		if strings.Contains("status", name) || strings.Contains("condition", name) {
			names[i] = "`" + name + "`"
		} else {
			names[i] = name
		}
		ids[i] = "?"

	}

	return "insert into " + table + " (" + strings.Join(names, ", ") + ") values (" + strings.Join(ids, ",") + ")"

}
func Save(item Items) error {

	db, errConn := openConnection()
	if errConn != nil {
		return errConn
	}

	stmt, err := db.Prepare(MakeSelectStatement("products"))
	if err != nil {
		return err
	}

	stmt.Exec(
		item.Id,
		item.Site_id,
		item.Title,
		item.Subtitle,
		item.Seller_id,
		item.Category_id,
		item.Official_store_id,
		item.Price,
		item.Base_price,
		item.Original_price,
		item.Currency_id,
		item.Initial_quantity,
		item.Available_quantity,
		item.Sold_quantity,
		item.Sale_terms,
		item.Buying_mode,
		item.Listing_type_id,
		item.Start_time,
		item.Stop_time,
		item.Condition,
		item.Permalink,
		item.Thumbnail_id,
		item.Thumbnail,
		item.Secure_thumbnail,
		item.Status,
		item.Warranty,
		item.Catalog_product_id,
		item.Domain_id,
		item.Health,
		item.Pictures,
		item.Description)

	defer db.Close()

	return nil
}
