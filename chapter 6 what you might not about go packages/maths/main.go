package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var Data []Entry
var TemplatesDir = "templates/*"
var DataTable = "data"

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	table := template.Must(template.ParseGlob(TemplatesDir))
	table.ExecuteTemplate(w, "index.html", Data)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	database := arguments[1]
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Emptying database table")
	deleteQuery, _, _ := sq.Delete(DataTable).ToSql()
	_, err = db.Exec(deleteQuery)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Populating", database)
	for i := 20; i < 50; i++ {
		random := rand.Intn(100-1) + 100
		insertData, args, _ := sq.Insert(DataTable).
			Columns("number", "double", "square").
			Values(random, 2*random, random*random).
			ToSql()
		_, _ = db.Exec(insertData, args...)
	}
	selectData, _, err := sq.Select("*").From(DataTable).ToSql()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := db.Query(selectData)
	if err != nil {
		fmt.Println(err)
	}
	number, double, square := 0, 0, 0
	for rows.Next() {
		err = rows.Scan(&number, &double, &square)
		entry := Entry{
			Number: number,
			Double: double,
			Square: square,
		}
		Data = append(Data, entry)
	}

	http.HandleFunc("/", Handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
