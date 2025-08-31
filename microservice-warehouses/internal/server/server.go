package server

import (
	"context"
	"errors"
	"flag"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/config"
	grpcHandler "github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/handler/grpc"
	httpHandler "github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/handler/http"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/logger"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/repository/gorm"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/service"
	pb "github.com/PIRSON21/grpc-microservices/microservice-warehouses/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	// Setup repository
	dbGorm := gorm.NewDbGorm(&cfg.DBConfig)
	defer dbGorm.Close()

	// Setup service
	warehouseService := service.NewWarehouseService(dbGorm)

	// Initialize handlers
	warehouseHTTPHandler := httpHandler.NewWarehouseHandler(warehouseService)
	warehouseGRPCHandler := grpcHandler.NewWarehouseHandler(warehouseService)

	// Setup router
	router := setupRouter(warehouseHTTPHandler)
	log.Debug("Router setup successfully")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.WithFields(log.Fields{
		"http_addr": cfg.HTTPAddr,
		"grpc_addr": cfg.GRPCAddr,
		"env":       cfg.Env,
	}).Info("Starting server...")

	// Start HTTP server
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	var grpcServer *grpc.Server
	// Start GRPC server
	go func() {
		lis, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			log.Fatalf("Failed to listen on %s: %v", cfg.GRPCAddr, err)
		}
		grpcServer = grpc.NewServer()
		pb.RegisterWarehouseServiceServer(grpcServer, warehouseGRPCHandler)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Info("Shutting down gracefully, press Ctrl+C again to force")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP Server forced to shutdown: %v", err)
	}

	grpcServer.GracefulStop()

	log.Info("Server exiting")
}

// setupRouter sets up the Gin router.
func setupRouter(warehouseHandler *httpHandler.WarehouseHTTPHandler) *gin.Engine {
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
