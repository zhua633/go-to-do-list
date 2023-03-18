package main

import (
	"fmt"
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

func createToDoItem(id int, dueDate time.Time, description string, priority int) {
	newTask := ToDo{
		ID:          id,
		DueDate:     dueDate,
		Description: description,
		Priority:    priority,
	}

	toDoList = append(toDoList, newTask)
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
	createToDoItem(1, time.Date(2023, 03, 18, 13, 30, 30, 30, time.UTC), "Finish laying out metadata struct for my TODO list project!", 1)
	createToDoItem(2, time.Date(2020, 03, 18, 13, 30, 30, 30, time.UTC), "Finish laying out metadata struct for my TODO list project!", 1)
	readToDoList()
	updateToDoListDescription(1, "Shower")
	readToDoList()
}
