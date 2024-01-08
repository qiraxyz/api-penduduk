package main

import (
	_ "datawarehouse/config/initialize"
	"datawarehouse/config/route"
	"datawarehouse/middleware"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Println("Starting Program ...")
	//migration.RunMigration()
	configs := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	app := fiber.New(configs)
	//middleware
	middleware.FiberMiddleware(app)
	// init route
	//route.SwaggerRoute(app)
	route.RouteInit(app)

	app.Listen(":2000")
}
