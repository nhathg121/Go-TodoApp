package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Hello world")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	counter := 0
	// tao list todos de luu tru task vao db
	todos := []Todo{}

	// Healthcheck
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})

	// add todos into list
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
		counter++
		todo.ID = counter

		todos = append(todos, *todo)

		log.Printf("Todo created: %v", todo)
		return c.JSON(todos)

	})

	// Make done todos
	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(404).SendString("Invalid ID")
		}
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break
			}
			// if todo not found return 404
			if i == len(todos)-1 {
				return c.Status(404).SendString("Todo not found")
			}
		}
		return c.JSON(todos)
	})

	// List all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// Delete todo by id
	app.Delete("/api/todos/:id/delete", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(404).SendString("Invalid ID")
		}
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
			// if todo not found return 404
			if i == len(todos)-1 {
				return c.Status(404).SendString("Todo not found")
			}
		}

		return c.JSON(todos)
	})
	// update todo by id
	app.Put("/api/todos/:id/update", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(404).SendString("Invalid ID")
		}
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i].Title = todo.Title
				todos[i].Body = todo.Body
				todos[i].Done = todo.Done
				break
			}
			// if todo not found return 404
			if i == len(todos)-1 {
				return c.Status(404).SendString("Todo not found")
			}
		}
		return c.JSON(todos)

	})

	// SHow all done task
	app.Get("/api/todos/done", func(c *fiber.Ctx) error {
		doneTodos := []Todo{}
		for _, todo := range todos {
			if todo.Done {
				doneTodos = append(doneTodos, todo)
			}
		}
		return c.JSON(doneTodos)

	})

	// Show all undone task
	app.Get("/api/todos/undone", func(c *fiber.Ctx) error {
		undoneTodos := []Todo{}
		for _, todo := range todos {
			if !todo.Done {
				undoneTodos = append(undoneTodos, todo)
			}
		}
		return c.JSON(undoneTodos)
	})

	log.Fatal(app.Listen(":3000"))

}
