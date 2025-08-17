package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"_id" bson: "_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello Worlds")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro loading dotenv File")
	}
	//PORT := os.Getenv("PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo")

	collection = client.Database("goland_db").Collection("todos")

	app := fiber.New()

	// Get all todos Endpoint
	app.Get("/api/todos", getTodos)
	// Create a new todo endpoint
	app.Post("/api/todos", createTodo)
	// Update todo endpoint
	app.Patch("/api/todos", updateTodo)
	// Delete todo endpoint
	app.Delete("/api/todos", deleteTodo)

	//*****************************************************
	//*****************************************************
	//*****************************************************
	// In Memory todo app
	//*****************************************************
	//*****************************************************
	//*****************************************************

	// app := fiber.New()

	// todos := []Todo{}

	// app.Get("/api/todos", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(todos)
	// })

	// // create a todo
	// app.Post("/api/todos", func(c *fiber.Ctx) error {
	// 	todo := &Todo{}

	// 	if err := c.BodyParser(todo); err != nil {
	// 		return err
	// 	}

	// 	if todo.Body == "" {
	// 		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
	// 	}

	// 	todo.ID = len(todos) + 1
	// 	todos = append(todos, *todo)

	// 	return c.Status(201).JSON(todo)
	// })

	// //update a todo
	// app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos[i].Completed = true
	// 			return c.Status(200).JSON(fiber.Map{
	// 				"Message": "Todo has been updated",
	// 				"todo":    todos[i],
	// 			})
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	// })

	// // Delete todo
	// app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos = append(todos[:i], todos[i+1:]...)
	// 			return c.Status(200).JSON(fiber.Map{
	// 				"message": "Successfully deleted the todo",
	// 				"Todo":    todo,
	// 			})
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "No Todos found"})
	// })

	// log.Fatal(app.Listen(":" + PORT))

}
