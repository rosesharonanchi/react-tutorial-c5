package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"piggy.com/internal/db/repo"
	"piggy.com/internal/db/sqlc"
	"piggy.com/internal/handlers"
	"piggy.com/internal/middleware"
	"piggy.com/internal/piggyservice"
)

func main() {
	route := gin.Default()

	// Configure Cors
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Healthcheck
	route.GET("/api/v1/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Healthy!",
		})
	})

	// Initialize repo and apply migrations
	ctx := context.Background()
	dbUrl := "postgres://piggy:secret@127.0.0.1:5432/piggydb?sslmode=disable"
	dbConn, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established!")
	repostory := repo.NewRepository(dbConn)
	if err :=repo.MigrateUp(dbUrl, "./internal/db/migrations", zerolog.Nop().With().Logger());err !=nil{
		panic(err)
	}

	// Initialize service
	appService := piggyservice.NewService(repostory)
	authService := piggyservice.NewAuthService(repostory.Do().(*sqlc.Queries))
	handlers := handlers.NewHandler(appService, authService)
  

	v1 := route.Group("/api/v1")
	// Define application endpoints
	route.POST("//api/v/1signup", handlers.SignUp)
    route.POST("/api/v1/login", handlers.Login)

	//protected routes
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware())
    
	
	{
		protected.GET("/transactions", handlers.GetTransactions) 
		protected.POST("/transactions", handlers.CreateTransaction)
		
	}
	fmt.Println("Server running on port 8080")
	route.Run()
}