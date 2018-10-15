package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*fun insertItem(name string, description string, price float32) {
	db, err := sql.Open("mysql", "root@tcp(localhost)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// TODO run an insert query
}*/

type Item struct {
	ID          int32
	Name        string
	Description string
	Price       float32
	next        *Item
}

type Items []Item

func lookupItem(queryId int) Item {
	db, err := sql.Open("mysql", "root@tcp(localhost)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/*var (
		id   int
		name string
		description string
		price float32
	)*/
	item := Item{}
	rows, err := db.Query("select id, name, description, price from items where id = ?", queryId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(&item.ID, &item.Name)
		//rts := fmt.Sprintf("%d %s", item.ID, item.Name)
		//return string(rts)
		return item
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return item
}

func lookupItems() Items {
	db, err := sql.Open("mysql", "root@tcp(localhost)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/*var (
		id   int
		name string
		description string
		price float32
	)*/
	items := Items{}
	rows, err := db.Query("select id, name, description, price from items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price)
		if err != nil {
			fmt.Println("Err reading row")
		}
		items = append(items, item)
		log.Println(item.ID, item.Name)
		//rts := fmt.Sprintf("%d %s", item.ID, item.Name)
		//return string(rts)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return items
}

func insertItem(name string, description string, price string) bool {
	db, err := sql.Open("mysql", "root@tcp(localhost)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insertStatement := fmt.Sprintf("INSERT INTO items (name,description,price) VALUES('%s','%s','%s') ", name, description, price)
	result, err := db.Exec(insertStatement)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted row %s", result)
	return true

}
