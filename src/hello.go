package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Define metadata struct
type ToDo struct {
	ID          int
	DueDate     time.Time
	Description string
	Priority    int
}

// Define a to do list slice
var toDoList []ToDo

func pop(index int) {
	// If the element to be removed is at the head
	if (index - 1) == 0 {
		toDoList = toDoList[index:]
	} else if (index + 1) > len(toDoList) {
		toDoList = toDoList[:index]
	} else {
		toDoList = append(toDoList[:(index-1)], toDoList[(index+1):]...)
	}
}

func createToDoItem(id int, dueDate time.Time, description string, priority int) ToDo {
	newTask := ToDo{
		ID:          id,
		DueDate:     dueDate,
		Description: description,
		Priority:    priority,
	}

	toDoList = append(toDoList, newTask)
	return newTask
}

func readToDoList() int {
	item, err := fmt.Println(toDoList)
	if err != nil {
		fmt.Println(err)
	}

	return item
}

func updateToDoListDescription(id int, description string) {
	if id < 0 || id >= len(toDoList) {
		fmt.Println("Invalid index provided")
		return
	}
	toDoList[id].Description = description
}

func updateToDoListDueDate(id int, dueDate time.Time) {
	if id < 0 || id >= len(toDoList) {
		fmt.Println("Invalid index provided")
		return
	}
	toDoList[id].DueDate = dueDate
}

func updateToDoListPriority(id int, priority int) {
	if id < 0 || id >= len(toDoList) {
		fmt.Println("Invalid index provided")
	}
	toDoList[id].Priority = priority
}

func deleteToDoList(id int) {
	pop(1)
}

func main() {
	// Create a new ToDo item
	// createToDoItem(1, time.Date(2023, 03, 18, 13, 30, 30, 30, time.UTC), "Finish laying out metadata struct for my TODO list project!", 1)
	// createToDoItem(2, time.Date(2020, 03, 18, 13, 30, 30, 30, time.UTC), "Finish laying out metadata struct for my TODO list project!", 1)
	// readToDoList()
	// updateToDoListDescription(1, "Shower")
	// readToDoList()
	// Create some example ToDo items
	toDoList = []ToDo{
		{ID: 1, Description: "Finish writing report", DueDate: time.Date(2023, time.March, 31, 23, 59, 59, 0, time.UTC), Priority: 2},
		{ID: 2, Description: "Submit report to boss", DueDate: time.Date(2023, time.April, 7, 23, 59, 59, 0, time.UTC), Priority: 1},
	}

	// Define a function that will be used to render our to-do list in HTML
	renderToDoList := func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templ.Execute(w, toDoList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/delete/"))
		if err != nil {
			http.Error(w, "Invalid item ID", http.StatusBadRequest)
			return
		}

		deleteToDoList(id)

		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()
		description := r.Form.Get("description")
		dueDate := r.Form.Get("dueDate")
		priority := r.Form.Get("priority")

		// type conversions
		pr, err := strconv.Atoi(priority)
		if err != nil {
			panic(err)
		}

		date, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			panic(err)
		}

		newItem := createToDoItem(len(toDoList)+1, date, description, pr)
		toDoList = append(toDoList, newItem)

		http.Redirect(w, r, "/", http.StatusFound)
	})
	// Define a route that will render our to-do list when requested
	http.HandleFunc("/", renderToDoList)

	// Start the server and listen for incoming requests
	http.ListenAndServe(":8080", nil)
}
