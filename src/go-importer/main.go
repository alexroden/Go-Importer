package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	
	_ "github.com/go-sql-driver/mysql"
)

var rowStruct = [5]string{
		"sku",
		"plu",
		"name",
		"size",
		"size_sort",
	}

func clean(s string) string {
	if len(s) > 0 {
		if s[0] == '"' {
			s = s[1:]
		}
	
		if s[len(s) - 1] == '"' {
			s = s[:len(s) - 1]
		}
	}

	return s
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "go"
	dbPass := "password"
	dbName := "go"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:4306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func insert(data map[string]string) {
	db := dbConn()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO `products`(`sku`, `plu`, `name`, `size`, `size_sort`) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(data["sku"], data["plu"], data["name"], data["size"], data["size_sort"])
} 

func truncate() {
	db := dbConn()
	defer db.Close()

	_, err := db.Query("TRUNCATE TABLE `products`")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	_, err :=  ioutil.ReadFile("products.csv")
	if (err != nil) {
		panic(err.Error())
	}

	f, err := os.Open("products.csv")
	if (err != nil) {
		panic(err.Error())
	}
	defer f.Close()

	truncate()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var mapper = make(map[string]string)
		row := strings.Split(scanner.Text(), ", ")
		for i, data := range row {
			mapper[rowStruct[i]] = clean(data)
		}

		insert(mapper)
	} 

	fmt.Printf("Done! 🎉\n")
	return
}