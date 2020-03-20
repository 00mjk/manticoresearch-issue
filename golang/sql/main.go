package main

import (
	"log"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cl, err := sql.Open("mysql", "@tcp(127.0.0.1:9306)/")
	if err != nil {
		log.Fatal(err)
	}



	res, err := cl.Exec(query)
	fmt.Println(res, err)
}

