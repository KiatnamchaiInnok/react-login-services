package main

import (
	"log"
	"os"
	"strings"

	"backend/configs"
	"backend/modules/servers"

	databases "backend/pkg/databases/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.MySQL.Host = os.Getenv("DB_HOST")
	cfg.MySQL.Port = os.Getenv("DB_PORT")
	cfg.MySQL.Username = os.Getenv("DB_USERNAME")
	cfg.MySQL.Password = os.Getenv("DB_PASSWORD")
	cfg.MySQL.Database = os.Getenv("DB_DATABASE")

	// New Database
	db, err := databases.NewMySQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// defer db.Close()

	s := servers.NewServer(cfg, db)

	s.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Content-Type, Authorization",
	}))

	s.Start()
}
