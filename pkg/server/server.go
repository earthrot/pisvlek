package server

import (
	"context"
	"log"
	"time"

	"github.com/earthrot/pisvlek/pkg/config"
	"github.com/earthrot/pisvlek/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/bbolt/v2"
	"github.com/gofiber/template/django/v3"
)

var cfg *config.Config
var store *session.Store
var q *db.Queries
var ctx context.Context

func Run(Configuration *config.Config, Queries *db.Queries) {
	cfg = Configuration
	q = Queries
	ctx = context.Background()

	views := django.New("./templates", ".html")

	storage := bbolt.New(bbolt.Config{
		Database: "./assets/sessions.bbolt",
	})

	store = session.New(session.Config{
		Storage:    storage,
		Expiration: 90 * time.Hour * 24,
		KeyLookup:  "cookie:pisvlek/session",
	})

	app := fiber.New(fiber.Config{
		Views:             views,
		ViewsLayout:       "main",
		PassLocalsToViews: true,
	})

	// door blijven draaien ipv crashen bij een panic
	app.Use(recover.New())

	// static content (css, js, etc)
	app.Static("/", "./public")

	app.Get("/login", GetLogin).Name("login")
	app.Post("/login", PostLogin)

	// alles onder deze route werkt alleen voor ingelogde users
	app.Use(firewall)

	app.Get("/", GetWelcome)
	app.Get("/logout", GetLogout)
	app.Get("/users", GetUsers)

	log.Fatal(app.Listen(":54321"))
}
