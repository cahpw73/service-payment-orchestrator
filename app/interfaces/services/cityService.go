package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/client/redis"
)

type CityServiceInterface interface {
	GetCities(idSubCategory string) ([]responseMiddleware.CityResponseMiddleware, error)
}

type CityService struct {
	middlewareService MiddlewareServiceInterface
	redisClient       redis.RedisServiceInterface
}

func NewCityService(redisClient redis.RedisServiceInterface, middlewareService MiddlewareServiceInterface) CityServiceInterface {
	return &CityService{redisClient: redisClient, middlewareService: middlewareService}
}

func (m *CityService) GetCities(idSubCategory string) ([]responseMiddleware.CityResponseMiddleware, error) {
	var cities []responseMiddleware.CityResponseMiddleware

	citiesRedis, err := m.redisClient.GetData("servicePaymentMiddleware:Cities:" + idSubCategory)
	if err == nil {
		if err := json.Unmarshal([]byte(citiesRedis), &cities); err == nil {
			return cities, nil
		}

		return nil, fmt.Errorf("error parseando JSON de Redis: %v", err)
	}

	_, err = m.middlewareService.MakeRequest(http.MethodGet, "payment-services-manager/bs/v1/sub-categories/"+idSubCategory+"/cities", nil, &cities)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	citiesJSON, err := json.Marshal(cities)

	if err != nil {
		return nil, fmt.Errorf("error al serializar ciudades a JSON: %v", err)
	}

	m.redisClient.SetData("servicePaymentMiddleware:Cities:"+idSubCategory, string(citiesJSON))

	return cities, nil
}
