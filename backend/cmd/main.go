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
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "require" // Supabase MUST use SSL
	}

	if host == "" || port == "" || user == "" || dbname == "" {
		panic("missing required database env vars")
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		sslmode,
	)
}


func main() {
	godotenv.Load()
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



func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
