package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	"piggy.com/internal/db/repo"
	"piggy.com/internal/db/sqlc"
	"piggy.com/internal/handlers"
	"piggy.com/internal/middleware"
	"piggy.com/internal/piggyservice"
)

// func init() {
// 	net.DefaultResolver.PreferGo = true
// }
func buildDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		func() string {
			if os.Getenv("DB_SSLMODE") == "" {
				return "disable"
			}
			return os.Getenv("DB_SSLMODE")
		}(),
	)
}


func main() {
	godotenv.Load()
	route := gin.Default()

	// Configure Cors
	route.Use(cors.New(cors.Config{
	AllowOrigins: []string{
		"http://localhost:3000",
		"https://piggy-save-kappa.vercel.app",
		"https://piggy-save-git-main-rosesharonanchis-projects.vercel.app",
	},
	AllowMethods: []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
		"OPTIONS",
	},
	AllowHeaders: []string{
		"Origin",
		"Content-Type",
		"Authorization",
		"X-User-ID",
	},
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
 //
 var resolver = &net.Resolver{
	PreferGo: true,
	Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, "udp4", "8.8.8.8:53")
	},
}

net.DefaultResolver = resolver
	// Initialize repo and apply migrations
	ctx := context.Background()
	dbUrl := buildDBUrl()
	// dbConn, err := pgxpool.New(ctx, dbUrl)
	dialer := &net.Dialer{
	Timeout: 5 * time.Second,
}

config, err := pgxpool.ParseConfig(dbUrl)
if err != nil {
	panic(err)
}

// 🔥 FORCE IPv4 ONLY (THIS FIXES YOUR ERROR)
config.ConnConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
	return dialer.DialContext(ctx, "tcp4", addr)
}

dbConn, err := pgxpool.NewWithConfig(ctx, config)
if err != nil {
	panic(err)
}
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established!")
	repostory := repo.NewRepository(dbConn)
	migrationPath := getEnv("MIGRATIONS_PATH", "./internal/db/migrations")
	if err :=repo.MigrateUp(dbUrl, migrationPath, zerolog.Nop().With().Logger());err !=nil{
		panic(err)
	}

	// Initialize service
	appService := piggyservice.NewService(repostory)
	authService := piggyservice.NewAuthService(repostory.Do().(*sqlc.Queries))
	handlers := handlers.NewHandler(appService, authService)
  

	v1 := route.Group("/api/v1")
	// Define application endpoints
	route.POST("/api/v1/signup", handlers.SignUp)
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



func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
