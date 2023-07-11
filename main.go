package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller"
	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller/routes"
	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	"github.com/foliveiracamara/elixir-api-go/src/domain/service"
	database "github.com/foliveiracamara/elixir-api-go/src/port/database"
	repository "github.com/foliveiracamara/elixir-api-go/src/port/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.DatabaseConn()
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      initDependencies(db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	// Graceful shutdown
	go func() {
		logger.Info("Listening and serving",
			zap.String("address", srv.Addr),
		)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		logger.Info("Stopped serving new connections.",
			zap.String("journey", "server"))
	}()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}

	logger.Info("Graceful shutdown complete.")
}

func initDependencies(db *sql.DB) *gin.Engine {

	donorRepo := repository.InitDonorRepository(db)
	donorService := service.NewDonorDomainService(donorRepo.DonorRepository)
	donorController := controller.NewDonorController(donorService)

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup, donorController)
	return r
}
