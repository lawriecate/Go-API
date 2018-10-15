package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conv err")
	}
	result := lookupItem(idInt)

	b, err := json.Marshal(result)
	fmt.Fprintf(w, "%s", b)

}

func getItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	result := lookupItems()

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Err getting results")
	}

	fmt.Fprintf(w, "%s", b)

}

func postItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")

	insertItem(name, description, price)

}

func main() {
	router := httprouter.New()
	router.GET("/items/:id", getItem)
	router.GET("/items", getItems)
	router.POST("/items", postItem)
	http.ListenAndServe(":8080", router)
}
