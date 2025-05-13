package handlers

import (
	"fmt"
	"service-payment-orchestrator/app/errors"
	"service-payment-orchestrator/app/interfaces/services"
	"service-payment-orchestrator/app/models"
	requestMiddleware "service-payment-orchestrator/app/models/RequestMiddleware"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/utils"

	"github.com/gofiber/fiber/v2"
)

const logAffiliationHandler = "Handler Validate Affiliation"

type AffiliationHandlerInterface interface {
	GetAffiliationByAccount(c *fiber.Ctx) error
	PostValidateAffiliation(c *fiber.Ctx) error
	PostSaveAffiliation(c *fiber.Ctx) error
	PostUpdateAffiliation(c *fiber.Ctx) error
	DeleteAffiliation(c *fiber.Ctx) error
}

type AffiliationHandler struct {
	affiliationService services.AffiliationInterface
}

func NewAffiliationHandler(affiliationService services.AffiliationInterface) AffiliationHandlerInterface {
	return &AffiliationHandler{affiliationService: affiliationService}
}

func (ct *AffiliationHandler) GetAffiliationByAccount(c *fiber.Ctx) error {
	account := c.Params("accountNumber")

	utils.LogRegister(c, logAffiliationHandler, "INIT  - PARAMS: {accountNumber: "+account+"}")

	affiliations, err := ct.affiliationService.GetAffiliationByAccount(account)
	if err != nil {
		utils.LogErrorRegister(c, logAffiliationHandler, err)
		errCode := errors.GetError("ERR011")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[[]responseMiddleware.AffiliationDataResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    affiliations,
	}

	logMessage := fmt.Sprintf("Response: %+v", response)
	utils.LogRegister(c, logAffiliationHandler, "END - RESPONSE: "+logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ct *AffiliationHandler) PostValidateAffiliation(c *fiber.Ctx) error {

	var validationRequest requestMiddleware.ValidateAffiliationRequestMiddleware

	if err := c.BodyParser(&validationRequest); err != nil {

		errCode := errors.GetError("ERR010")

		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	logMessage := fmt.Sprintf("INIT - PARAMS: {request affiliation: %+v }", validationRequest)
	utils.LogRegister(c, logAffiliationHandler, logMessage)

	services, err := ct.affiliationService.PostValidateAffiliation(validationRequest)
	if err != nil {

		utils.LogErrorRegister(c, logAffiliationHandler, err)
		errCode := errors.GetError("ERR011")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.ValidateAffiliationResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    services,
	}

	logMessage = fmt.Sprintf("END - RESPONSE: Response: %+v", response)
	utils.LogRegister(c, logAffiliationHandler, logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ct *AffiliationHandler) PostSaveAffiliation(c *fiber.Ctx) error {
	utils.LogRegister(c, logAffiliationHandler, "INIT  - Save affiliation")

	var affiliationRequest requestMiddleware.RegisterAffiliation

	if err := c.BodyParser(&affiliationRequest); err != nil {
		errCode := errors.GetError("ERR010")

		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	logMessage := fmt.Sprintf("INIT - PARAMS: {request affiliation: %+v }", affiliationRequest)
	utils.LogRegister(c, logAffiliationHandler, logMessage)

	affiliation, err := ct.affiliationService.SaveAffiliation(affiliationRequest)

	if err != nil {
		utils.LogErrorRegister(c, logAffiliationHandler, err)
		errCode := errors.GetError("ERR011")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.RegisterAffiliationResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    affiliation,
	}

	logMessage = fmt.Sprintf("END - RESPONSE: Response: %+v", response)
	utils.LogRegister(c, logAffiliationHandler, logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ct *AffiliationHandler) PostUpdateAffiliation(c *fiber.Ctx) error {
	utils.LogRegister(c, logAffiliationHandler, "INIT  - Save affiliation")

	var affiliationRequest requestMiddleware.UpdateAffiliation

	if err := c.BodyParser(&affiliationRequest); err != nil {
		errCode := errors.GetError("ERR010")

		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	logMessage := fmt.Sprintf("INIT - PARAMS: {request affiliation: %+v }", affiliationRequest)
	utils.LogRegister(c, logAffiliationHandler, logMessage)

	affiliation, err := ct.affiliationService.UpdateAffiliation(affiliationRequest)

	if err != nil {
		utils.LogErrorRegister(c, logAffiliationHandler, err)
		errCode := errors.GetError("ERR011")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.RegisterAffiliationResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    affiliation,
	}

	logMessage = fmt.Sprintf("END - RESPONSE: Response: %+v", response)
	utils.LogRegister(c, logAffiliationHandler, logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (ct *AffiliationHandler) DeleteAffiliation(c *fiber.Ctx) error {
	account := c.Params("accountNumber")
	affiliationCode := c.Params("affiliationCode")

	utils.LogRegister(c, logAffiliationHandler, "INIT  - PARAMS: {accountNumber: "+account+"}")

	affiliation, err := ct.affiliationService.DeleteAffiliation(account, affiliationCode)
	if err != nil {
		utils.LogErrorRegister(c, logAffiliationHandler, err)
		errCode := errors.GetError("ERR011")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.RegisterAffiliationResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    affiliation,
	}

	logMessage := fmt.Sprintf("Response: %+v", response)
	utils.LogRegister(c, logAffiliationHandler, "END - RESPONSE: "+logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
