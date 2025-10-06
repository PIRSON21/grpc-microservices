package http

import "github.com/PIRSON21/grpc-microservices/microservice-products/internal/service"

type ProductHTTPHandler struct {
	service *service.ProductService
}

func NewProductHTTPHandler(service *service.ProductService) *ProductHTTPHandler {
	return &ProductHTTPHandler{service: service}
}
