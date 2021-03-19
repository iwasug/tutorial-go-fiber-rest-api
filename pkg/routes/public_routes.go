package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/books", controllers.GetBooks)   // get list of all books
	route.Get("/book/:id", controllers.GetBook) // get one book by ID

	// Routes for POST method:
	route.Post("/user/login", controllers.UserLogin)     // auth user and return JWT
	route.Post("/user/register", controllers.CreateUser) // register a new user

	// Route for generate JWT token.
	// route.Get("/token", func(c *fiber.Ctx) error {
	// 	credentials := []string{"book:delete"}
	// 	token, err := utils.GenerateNewJWTAccessToken(uuid.NewString(), credentials)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(fiber.Map{
	// 		"error": false,
	// 		"msg":   token,
	// 	})
	// })
}
