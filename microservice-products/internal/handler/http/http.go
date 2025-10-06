package http

import (
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/service"
	"github.com/rs/zerolog"
)

type ProductHTTPHandler struct {
	service *service.ProductService
	logger  *zerolog.Logger
}

func NewProductHTTPHandler(service *service.ProductService, logger *zerolog.Logger) *ProductHTTPHandler {
	return &ProductHTTPHandler{service: service, logger: logger}
}
