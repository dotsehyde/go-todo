package server

import (
	"context"
	"database/sql"
	"fmt"
	"go-todo/internal/controller"
	"go-todo/internal/database"
	"go-todo/internal/middleware"
	custom_sql "go-todo/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	port       int
	controller *controller.Controller
}

func NewServer() *http.Server {
	// Declare Database config
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal(err)
	}
	// create tables
	if _, err := db.ExecContext(ctx, custom_sql.Schema); err != nil {
		log.Fatal(err.Error())
	}
	// Init SessionManager
	middleware.InitSessionManager(db)
	queries := database.New(db)
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		controller: &controller.Controller{
			Db: queries,
		},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
