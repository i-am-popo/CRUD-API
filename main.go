package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")

		query := r.URL.Query()
		id, _ := strconv.Atoi(query.Get("id"))
		fmt.Println(id)

		switch r.Method {

			
		case "GET":
						//ISSUE: Create a GET request to get all todos
		case "POST":

			//ISSUE: Create a POST request to add a new todo
			
		case "DELETE":
			  //ISSUE: Create a DELETE request to delete a todo
		case "PUT":
			//ISSUE: Create a PUT request to update a todo
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}