package services

import (
	"fmt"
	"net/http"
	requestMiddleware "service-payment-orchestrator/app/models/RequestMiddleware"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
)

type PayDebtServiceInterface interface {
	PostPayDebt(payload requestMiddleware.PayDebtRequestMiddleware) (responseMiddleware.TransactionDataResponseMiddleware, error)
}

type PayDebtService struct {
	middlewareService MiddlewareServiceInterface
}

func NewPayDebtService(middlewareService MiddlewareServiceInterface) PayDebtServiceInterface {
	return &PayDebtService{middlewareService: middlewareService}
}

func (m *PayDebtService) PostPayDebt(payload requestMiddleware.PayDebtRequestMiddleware) (responseMiddleware.TransactionDataResponseMiddleware, error) {
	var payDebt responseMiddleware.TransactionDataResponseMiddleware

	_, err := m.middlewareService.MakeRequest(http.MethodPost, "payment-services-manager/bs/v1/payments/debts", payload, &payDebt)

	if err != nil {
		return payDebt, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	return payDebt, nil
}
