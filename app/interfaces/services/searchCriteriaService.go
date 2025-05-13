package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/client/redis"
)

type SearchCriteriaInterface interface {
	GetSearchCriteria(serviceCode string) (responseMiddleware.SearchCriteriaResponseMiddleware, error)
}

type SearchCriteriaService struct {
	middlewareService MiddlewareServiceInterface
	redisClient       redis.RedisServiceInterface
}

func NewSearchCriteria(redisClient redis.RedisServiceInterface, middlewareService MiddlewareServiceInterface) SearchCriteriaInterface {
	return &SearchCriteriaService{redisClient: redisClient, middlewareService: middlewareService}
}

func (ct *SearchCriteriaService) GetSearchCriteria(serviceCode string) (responseMiddleware.SearchCriteriaResponseMiddleware, error) {
	var criteria responseMiddleware.SearchCriteriaResponseMiddleware

	criteriaRedis, err := ct.redisClient.GetData("servicePaymentMiddleware:SearchCriteria:" + serviceCode)
	if err == nil {
		if err := json.Unmarshal([]byte(criteriaRedis), &criteria); err == nil {
			return criteria, nil
		}

		return criteria, fmt.Errorf("error parseando JSON de Redis: %v", err)
	}

	_, err = ct.middlewareService.MakeRequest(http.MethodGet, "payment-services-manager/bs/v1/search-criteria?serviceCode="+serviceCode, nil, &criteria)

	if err != nil {
		return criteria, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	criteriaJSON, err := json.Marshal(criteria)

	if err != nil {
		return criteria, fmt.Errorf("error al serializar criterios de busqueda a JSON: %v", err)
	}

	ct.redisClient.SetData("servicePaymentMiddleware:SearchCriteria:"+serviceCode, string(criteriaJSON))

	return criteria, nil
}
