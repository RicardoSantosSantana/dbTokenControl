package dbTokenControl

import "fmt"

func checkErr(err error) {

	if err != nil {
		fmt.Println(err.Error())
		//apiMeli.Logger.Log_message(err.Error())
		panic(err)
	}
}
