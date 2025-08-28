package server

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/config"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/handler"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/logger"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/repository/nop"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// InitServer initializes the server.
func InitServer() {
	// Parse config
	path := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	cfg := config.MustLoadConfig(path)
	log.Println("Configuration loaded successfully", cfg)

	// Initialize logger
	err := logger.SetupLogger(&cfg.LoggerConfig)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	log.Debug("Logger initialized successfully")

	// Setup service
	warehouseService := service.NewWarehouseService(&nop.NOP{})

	// Initialize handlers
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)

	// Setup router
	router := setupRouter(warehouseHandler)
	log.Debug("Router setup successfully")

	// Start server
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.WithFields(log.Fields{
		"addr": cfg.Addr,
		"env":  cfg.Env,
	}).Info("Starting server...")

	// Run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Info("Shutting down gracefully, press Ctrl+C again to force")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Info("Server exiting")
}

// setupRouter sets up the Gin router.
func setupRouter(warehouseHandler *handler.WarehouseHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiRoutes := r.Group("/api")

	apiRoutes.GET("/warehouses", warehouseHandler.GetWarehouses)

	apiRoutes.POST("/warehouses", warehouseHandler.CreateWarehouse)

	apiRoutes.GET("/warehouses/:id", warehouseHandler.GetWarehouseByID)

	return r
}
