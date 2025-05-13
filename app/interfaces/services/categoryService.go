package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/client/redis"
)

type CategoryServiceInterface interface {
	GetCategories() ([]responseMiddleware.CategoryResponseMiddleware, error)
}

type CategoryService struct {
	middlewareService MiddlewareServiceInterface
	redisClient       redis.RedisServiceInterface
}

func NewCategoryService(redisClient redis.RedisServiceInterface, middlewareService MiddlewareServiceInterface) CategoryServiceInterface {
	return &CategoryService{redisClient: redisClient, middlewareService: middlewareService}
}

func (m *CategoryService) GetCategories() ([]responseMiddleware.CategoryResponseMiddleware, error) {
	var categories []responseMiddleware.CategoryResponseMiddleware

	categoriesRedis, err := m.redisClient.GetData("servicePaymentMiddleware:Categories")
	if err == nil {
		if err := json.Unmarshal([]byte(categoriesRedis), &categories); err == nil {
			return categories, nil
		}

		return nil, fmt.Errorf("error parseando JSON de Redis: %v", err)
	}

	_, err = m.middlewareService.MakeRequest(http.MethodGet, "payment-services-manager/bs/v1/categories", nil, &categories)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	categoriesJSON, err := json.Marshal(categories)
	if err != nil {
		return nil, fmt.Errorf("error al serializar categorías a JSON: %v", err)
	}

	m.redisClient.SetData("servicePaymentMiddleware:Categories", string(categoriesJSON))

	return categories, nil
}
