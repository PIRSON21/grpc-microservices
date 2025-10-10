package server

import "github.com/PIRSON21/grpc-microservices/microservice-inventory/internal/logger"

// StartServer start an Inventory server.
func StartServer() {
	// TODO: implement the config reading

	// Initialize logger
	log := logger.GetLogger() // TODO: need to initialize logger first, now it returns a default text logger
	log.Info("Logger initialized")

	// TODO: implement the repository logic

	// TODO: implement the service logic

	// TODO: implement the HTTP server logic

	// TODO: implement the gRPC server logic

	// TODO: implement the graceful shutdown logic
}
