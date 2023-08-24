package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huybne/api-todo-fiber/controllers" // replace
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}
