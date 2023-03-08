package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/routes"
	"github.com/emekarr/coding-test-busha/server_response"
)

func main() {
	defer func() {
		CleanUp()
	}()

	err := godotenv.Load()
	if os.Getenv("GIN_MODE") != "release" {
		if err != nil {
			panic("could not load env variables")
		}
	}

	StartServices()

	server := gin.Default()

	server.Use(cors.Default())

	baseRouter := server.Group("/api")
	{
		routes.InitRouter(baseRouter)
	}

	server.GET("/ping", func(ctx *gin.Context) {
		server_response.Respond(ctx, http.StatusOK, "server is up and running", true, nil, nil)
	})

	server.NoRoute(func(ctx *gin.Context) {
		server_response.Respond(ctx, http.StatusNotFound, "this route does not exist", false, nil, nil)
	})

	gin_mode := os.Getenv("GIN_MODE")
	port := os.Getenv("PORT")
	if gin_mode == "debug" {
		server.Run(port)
	} else if gin_mode == "release" {
		server.Run(":" + port)
	} else {
		panic("invalid gin mode used")
	}
	logger.Info("server is up on port" + port)
}
