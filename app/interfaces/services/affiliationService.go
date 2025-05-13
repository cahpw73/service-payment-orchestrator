package services

import (
	"fmt"
	"log"
	"net/http"
	"service-payment-orchestrator/app/errors"
	requestMiddleware "service-payment-orchestrator/app/models/RequestMiddleware"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/data"
	"strconv"
)

type AffiliationInterface interface {
	GetAffiliationByAccount(account string) ([]responseMiddleware.AffiliationDataResponseMiddleware, error)
	PostValidateAffiliation(payload requestMiddleware.ValidateAffiliationRequestMiddleware) (responseMiddleware.ValidateAffiliationResponseMiddleware, error)
	SaveAffiliation(payload requestMiddleware.RegisterAffiliation) (responseMiddleware.RegisterAffiliationResponseMiddleware, error)
	UpdateAffiliation(data requestMiddleware.UpdateAffiliation) (responseMiddleware.RegisterAffiliationResponseMiddleware, error)
	DeleteAffiliation(accountNumber string, affiliationCode string) (responseMiddleware.RegisterAffiliationResponseMiddleware, error)
}

type AffiliationService struct {
	middlewareService MiddlewareServiceInterface
}

func NewAffiliationService(middlewareService MiddlewareServiceInterface) AffiliationInterface {
	return &AffiliationService{middlewareService: middlewareService}
}

func (m *AffiliationService) GetAffiliationByAccount(account string) ([]responseMiddleware.AffiliationDataResponseMiddleware, error) {
	var affiliations []responseMiddleware.AffiliationDataResponseMiddleware

	log.Printf("get affiliation by account")

	log.Printf("creating new oracle conexion")
	database, err := data.NewOracleService()
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR012").Code,
			Message:    errors.GetError("ERR012").Message,
			ErrorTrace: err,
		}
		return nil, errs
	}

	log.Printf("closing oracle conexion")
	defer database.DB.Close()

	log.Printf("getting nroPerson by account from SP")
	nroPersona, err := database.GetNroPersonaByAccount(account)
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR013").Code,
			Message:    errors.GetError("ERR013").Message,
			ErrorTrace: err,
		}
		return nil, errs
	}

	log.Printf("making request to middleware services")
	_, err = m.middlewareService.MakeRequest(http.MethodGet, "payment-services-manager/bs/v1/affiliations-services/persons/"+strconv.Itoa(nroPersona), nil, &affiliations)
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR014").Code,
			Message:    errors.GetError("ERR014").Message,
			ErrorTrace: err,
		}
		return nil, errs
	}

	log.Printf("returning data response")
	return affiliations, nil
}

func (m *AffiliationService) PostValidateAffiliation(payload requestMiddleware.ValidateAffiliationRequestMiddleware) (responseMiddleware.ValidateAffiliationResponseMiddleware, error) {
	var affiliation responseMiddleware.ValidateAffiliationResponseMiddleware

	_, err := m.middlewareService.MakeRequest(http.MethodPost, "payment-services-manager/bs/v1/affiliations/validate", payload, &affiliation)

	if err != nil {
		return affiliation, fmt.Errorf("error al hacer la solicitud: %v", err)
	}

	return affiliation, nil
}

func (m *AffiliationService) SaveAffiliation(payload requestMiddleware.RegisterAffiliation) (responseMiddleware.RegisterAffiliationResponseMiddleware, error) {
	var affiliation responseMiddleware.RegisterAffiliationResponseMiddleware

	log.Printf("making request to affiliations-services middleware services")
	_, err := m.middlewareService.MakeRequest(http.MethodPost, "payment-services-manager/bs/v1/affiliations-services", payload, &affiliation)

	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR015").Code,
			Message:    errors.GetError("ERR015").Message,
			ErrorTrace: fmt.Errorf("Error al hacer la solictud: %v", err),
		}

		return affiliation, errs
	}

	return affiliation, nil
}

func (m *AffiliationService) UpdateAffiliation(accountNumber requestMiddleware.UpdateAffiliation) (responseMiddleware.RegisterAffiliationResponseMiddleware, error) {
	var affiliation responseMiddleware.RegisterAffiliationResponseMiddleware

	log.Printf("creating new oracle conexion")
	database, err := data.NewOracleService()
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR012").Code,
			Message:    errors.GetError("ERR012").Message,
			ErrorTrace: err,
		}
		return affiliation, errs
	}

	log.Printf("closing oracle conexion")
	defer database.DB.Close()

	log.Printf("getting nroPerson by account from SP")
	nroPersona, err := database.GetNroPersonaByAccount(accountNumber.AccountNumber)
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR013").Code,
			Message:    errors.GetError("ERR013").Message,
			ErrorTrace: err,
		}
		return affiliation, errs
	}

	var payload requestMiddleware.UpdateAffiliationRequestMiddleware
	payload.PersonId = strconv.Itoa(nroPersona)
	payload.AffiliationCode = accountNumber.AffiliationCode
	payload.StateAffiliation = accountNumber.StateAffiliation

	log.Printf("making request to affiliations/update middleware services")
	_, err = m.middlewareService.MakeRequest(http.MethodPost, "payment-services-manager/bs/v1/affiliations/update", payload, &affiliation)

	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR017").Code,
			Message:    errors.GetError("ERR017").Message,
			ErrorTrace: fmt.Errorf("Error al hacer la solictud: %v", err),
		}

		return affiliation, errs
	}

	return affiliation, nil
}

func (m *AffiliationService) DeleteAffiliation(accountNumber string, affiliationCode string) (responseMiddleware.RegisterAffiliationResponseMiddleware, error) {
	log.Printf("Delete affiliation by accountNumber and AffiliationCode")

	var affiliation responseMiddleware.RegisterAffiliationResponseMiddleware

	log.Printf("creating new oracle conexion")
	database, err := data.NewOracleService()
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR012").Code,
			Message:    errors.GetError("ERR012").Message,
			ErrorTrace: err,
		}
		return affiliation, errs
	}

	log.Printf("closing oracle conexion")
	defer database.DB.Close()

	log.Printf("getting nroPerson by account from SP")
	nroPersona, err := database.GetNroPersonaByAccount(accountNumber)
	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR013").Code,
			Message:    errors.GetError("ERR013").Message,
			ErrorTrace: err,
		}
		return affiliation, errs
	}

	var payload requestMiddleware.DataDeleteAffiliationRequestMiddleware
	payload.PersonId = strconv.Itoa(nroPersona)
	payload.AffiliationCode = affiliationCode

	log.Printf(fmt.Sprintf("valor payload %v", payload))

	log.Printf("making request to affiliations/delete middleware services")
	_, err = m.middlewareService.MakeRequest(http.MethodPost, "payment-services-manager/bs/v1/affiliations/delete", payload, &affiliation)

	if err != nil {
		errs := &errors.AppError{
			Code:       errors.GetError("ERR016").Code,
			Message:    errors.GetError("ERR016").Message,
			ErrorTrace: fmt.Errorf("Error al hacer la solictud: %v", err),
		}

		return affiliation, errs
	}

	return affiliation, nil
}
