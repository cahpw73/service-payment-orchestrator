package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/client/redis"
)

type SubCategoryInterface interface {
	GetSubCategories(idCategory string) ([]responseMiddleware.SubCategoryResponseMiddleware, error)
}

type SubCategoryService struct {
	middlewareService MiddlewareServiceInterface
	redisClient       redis.RedisServiceInterface
}

func NewSubCategoryService(redisClient redis.RedisServiceInterface, middlewareService MiddlewareServiceInterface) SubCategoryInterface {
	return &SubCategoryService{redisClient: redisClient, middlewareService: middlewareService}
}

func (m *SubCategoryService) GetSubCategories(idCategory string) ([]responseMiddleware.SubCategoryResponseMiddleware, error) {
	var subCategories []responseMiddleware.SubCategoryResponseMiddleware

	subCategoriesRedis, err := m.redisClient.GetData("servicePaymentMiddleware:SubCategories:" + idCategory)
	if err == nil {
		if err := json.Unmarshal([]byte(subCategoriesRedis), &subCategories); err == nil {
			return subCategories, nil
		}

		return nil, fmt.Errorf("error parseando JSON de Redis: %v", err)
	}

	_, err = m.middlewareService.MakeRequest(http.MethodGet, "payment-services-manager/bs/v1/categories/"+idCategory+"/sub-categories", nil, &subCategories)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	subCategoriesJSON, err := json.Marshal(subCategories)

	if err != nil {
		return nil, fmt.Errorf("error al serializar categorías a JSON: %v", err)
	}

	m.redisClient.SetData("servicePaymentMiddleware:SubCategories:"+idCategory, string(subCategoriesJSON))

	return subCategories, nil
}
