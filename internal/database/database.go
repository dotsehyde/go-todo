package database

import (
	"context"
	"fmt"
	"go-todo/prisma/db"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var Client *db.PrismaClient

type Service interface {
	Health() map[string]string
}

type service struct {
	client *db.PrismaClient
}

func New() Service {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	Client = client
	s := &service{client: client}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := s.client.Todo.FindUnique(
		db.Todo.ID.Equals(""),
	).Exec(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
