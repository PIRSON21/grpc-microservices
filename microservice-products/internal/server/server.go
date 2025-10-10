package server

import (
	"errors"
	"flag"
	"log"
	"net/http"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/config"
	httpHandler "github.com/PIRSON21/grpc-microservices/microservice-products/internal/handler/http"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/logger"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/repository"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/service"
	"github.com/gin-gonic/gin"
)

// InitServer initializes the server.
func InitServer() {
	// Parse config
	path := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	cfg := config.MustLoadConfig(path)
	log.Println("Configuration loaded successfully", cfg)

	// Initialize logger
	log, err := logger.SetupLogger(&cfg.LoggerConfig)
	if err != nil {
		log.Fatal().Msgf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	log.Debug().Msg("Logger initialized successfully")

	// Setup repository
	productRepo := repository.NewProductRepository(&cfg.DBConfig)
	defer productRepo.Close()

	// Setup service
	productService := service.NewProductService(productRepo, log)

	// Initialize handlers
	warehouseHTTPHandler := httpHandler.NewProductHTTPHandler(productService, log)

	// Setup router
	router := setupRouter(warehouseHTTPHandler)
	log.Debug().Msg("Router setup successfully")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	log.Info().Fields(map[string]any{
		"http_addr": cfg.HTTPAddr,
		"grpc_addr": cfg.GRPCAddr,
		"env":       cfg.Env,
	}).Msg("Starting server...")

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Msgf("Server failed: %v", err)
	}
}

// setupRouter sets up the Gin router.
func setupRouter(warehouseHandler *httpHandler.ProductHTTPHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiRoutes := r.Group("/api")

	apiRoutes.GET("/products", warehouseHandler.GetProducts)
	apiRoutes.GET("/products/:id", warehouseHandler.GetProductByID)
	apiRoutes.POST("/products", warehouseHandler.CreateProduct)
	apiRoutes.PATCH("/products/:id", warehouseHandler.UpdateProduct)
	apiRoutes.DELETE("/products/:id", warehouseHandler.DeleteProduct)

	return r
}
