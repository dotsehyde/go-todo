package main

import (
	"fmt"
	"go-todo/internal/logger"
	"go-todo/internal/server"
)

func main() {
	logger.LoggerInit()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
