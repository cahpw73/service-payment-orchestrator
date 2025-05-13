package services

import (
	"fmt"
	"net/http"
	requestMiddleware "service-payment-orchestrator/app/models/RequestMiddleware"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
)

type GetDebtServiceInterface interface {
	PostDebtConsultation(payload requestMiddleware.GetDebtsRequestMiddleware) (responseMiddleware.GetDebtsResponseMiddleware, error)
}

type GetDebtService struct {
	middlewareService MiddlewareServiceInterface
}

func NewGetDebtService(middlewareService MiddlewareServiceInterface) GetDebtServiceInterface {
	return &GetDebtService{middlewareService: middlewareService}
}

func (m *GetDebtService) PostDebtConsultation(payload requestMiddleware.GetDebtsRequestMiddleware) (responseMiddleware.GetDebtsResponseMiddleware, error) {
	var debt responseMiddleware.GetDebtsResponseMiddleware

	_, err := m.middlewareService.MakeRequest(http.MethodPost, "payment-services-manager/bs/v1/debts/consultation", payload, &debt)

	if err != nil {
		return debt, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	return debt, nil
}
