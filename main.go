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
			// GET Function
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(todos)
		case "POST":
			// POST Function
			var todo Todo
			json.NewDecoder(r.Body).Decode(&todo)
			todos = append(todos, todo)
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(todo)
		case "DELETE":
			// Create a DELETE request 
			for index, todo := range todos {
				if todo.ID == id {
					todos = append(todos[:index], todos[index+1:]...)
					rw.WriteHeader(http.StatusOK)
					rw.Write([]byte(`{"message": "Success to delete todo"}`))
				}
			}
		case "PUT":
			// Create a PUT request
			for index, todo := range todos {
				if todo.ID == id {
					json.NewDecoder(r.Body).Decode(&todo)
					todos[index].ID = todo.ID
					todos[index].Task = todo.Task
					rw.WriteHeader(http.StatusOK)
					rw.Write([]byte(`{"message": "Success to update todo"}`))
				}
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}