package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/client/redis"
)

type ServicesSubCategoryInterface interface {
	GetServicesSubCategory(subCategoryID string, cityCode string) ([]responseMiddleware.ServiceBySubCategoryResponseMiddleware, error)
}

type ServicesSubCategoryService struct {
	middlewareService MiddlewareServiceInterface
	redisClient       redis.RedisServiceInterface
}

func NewServicesSubCategoryService(redisClient redis.RedisServiceInterface, middlewareService MiddlewareServiceInterface) ServicesSubCategoryInterface {
	return &ServicesSubCategoryService{redisClient: redisClient, middlewareService: middlewareService}
}

func (ct *ServicesSubCategoryService) GetServicesSubCategory(subCategoryID string, cityCode string) ([]responseMiddleware.ServiceBySubCategoryResponseMiddleware, error) {
	var servicesBySubCategory []responseMiddleware.ServiceBySubCategoryResponseMiddleware

	servicesBySubCategoryRedis, err := ct.redisClient.GetData("servicePaymentMiddleware:ServicesSubCategory:" + subCategoryID + cityCode)
	if err == nil {
		if err := json.Unmarshal([]byte(servicesBySubCategoryRedis), &servicesBySubCategory); err == nil {
			return servicesBySubCategory, nil
		}

		return nil, fmt.Errorf("error parseando JSON de Redis: %v", err)
	}

	_, err = ct.middlewareService.MakeRequest(http.MethodGet, "payment-services-manager/bs/v1/sub-categories/"+subCategoryID+"/cities/"+cityCode, nil, &servicesBySubCategory)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	servicesBySubCategoryJSON, err := json.Marshal(servicesBySubCategory)

	if err != nil {
		return nil, fmt.Errorf("error al serializar servicios por subcategoria a JSON: %v", err)
	}

	ct.redisClient.SetData("servicePaymentMiddleware:ServicesSubCategory:"+subCategoryID+cityCode, string(servicesBySubCategoryJSON))

	return servicesBySubCategory, nil
}
