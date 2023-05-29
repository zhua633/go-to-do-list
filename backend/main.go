package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Define metadata struct
type ToDo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	DueDate     time.Time `json:"dueDate"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Done        bool      `json:"done"`
}

// // Define a to do list slice
// var toDoList []ToDo

// func pop(index int) {
// 	// If the element to be removed is at the head
// 	if (index - 1) == 0 {
// 		toDoList = toDoList[index:]
// 	} else if (index + 1) > len(toDoList) {
// 		toDoList = toDoList[:index]
// 	} else {
// 		toDoList = append(toDoList[:(index-1)], toDoList[(index+1):]...)
// 	}
// }

// func createToDoItem(id int, dueDate time.Time, description string, priority int) {
// 	newTask := ToDo{
// 		ID:          id,
// 		DueDate:     dueDate,
// 		Description: description,
// 		Priority:    priority,
// 	}

// 	toDoList = append(toDoList, newTask)
// }

// func readToDoList() int {
// 	item, err := fmt.Println(toDoList)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return item
// }

// func updateToDoListDescription(id int, description string) {
// 	if id < 0 || id >= len(toDoList) {
// 		fmt.Println("Invalid index provided")
// 		return
// 	}
// 	toDoList[id].Description = description
// }

// func updateToDoListDueDate(id int, dueDate time.Time) {
// 	if id < 0 || id >= len(toDoList) {
// 		fmt.Println("Invalid index provided")
// 		return
// 	}
// 	toDoList[id].DueDate = dueDate
// }

// func updateToDoListPriority(id int, priority int) {
// 	if id < 0 || id >= len(toDoList) {
// 		fmt.Println("Invalid index provided")
// 	}
// 	toDoList[id].Priority = priority
// }

// func deleteToDoList(id int) {
// 	pop(1)
// }

func main() {
	app := fiber.New()

	// CORS middleware configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
		AllowHeaders: "Content-Type",
	}))

	app.Options("/api/todos", func(c *fiber.Ctx) error {
		// Handle the OPTIONS request for preflight
		// Respond with the appropriate CORS headers
		return c.SendStatus(fiber.StatusOK)
	})

	todos := []ToDo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Create a post endpoint
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		// Create unique id
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
	// Create a new ToDo item
	// createToDoItem(1, time.Date(2023, 03, 18, 13, 30, 30, 30, time.UTC), "Finish laying out metadata struct for my TODO list project!", 1)
	// createToDoItem(2, time.Date(2020, 03, 18, 13, 30, 30, 30, time.UTC), "Finish laying out metadata struct for my TODO list project!", 1)
	// readToDoList()
	// updateToDoListDescription(1, "Shower")
	// readToDoList()
}
