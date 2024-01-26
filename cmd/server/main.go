package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mycandys/orders/internal/database"
	"github.com/mycandys/orders/internal/env"
	"github.com/mycandys/orders/internal/routes"
	"github.com/mycandys/orders/internal/swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// important only in development
	_ = godotenv.Load(".env")

	port, err := env.GetEnvVar(env.PORT)
	if err != nil {
		panic(err)
	}

	db := database.Connect()
	defer database.Disconnect(db, context.Background())

	//amqp := rabbitmq.Connect()
	//defer rabbitmq.Close(amqp)

	app := routes.InitRouter()
	swagger.InitInfo()

	fmt.Printf("Swagger UI is available on http://localhost:%s/swagger/index.html\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")
}
