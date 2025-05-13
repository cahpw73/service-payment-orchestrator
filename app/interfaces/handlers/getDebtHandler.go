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

const logGetDebtsHandler = "Handler Debt Consultation"

type GetDebtHandlerInterface interface {
	PostDebtConsultation(c *fiber.Ctx) error
}

type GetDebtHandler struct {
	getDebtService services.GetDebtServiceInterface
}

func NewGetDebtHandler(getDebtService services.GetDebtServiceInterface) GetDebtHandlerInterface {
	return &GetDebtHandler{getDebtService: getDebtService}
}

func (ct *GetDebtHandler) PostDebtConsultation(c *fiber.Ctx) error {

	var debtRequest requestMiddleware.GetDebtsRequestMiddleware

	if err := c.BodyParser(&debtRequest); err != nil {

		errCode := errors.GetError("ERR003")

		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	logMessage := fmt.Sprintf("INIT - PARAMS: {request debt: %+v }", debtRequest)
	utils.LogRegister(c, logGetDebtsHandler, logMessage)

	services, err := ct.getDebtService.PostDebtConsultation(debtRequest)
	if err != nil {

		utils.LogErrorRegister(c, logGetDebtsHandler, err)
		errCode := errors.GetError("ERR004")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.GetDebtsResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    services,
	}

	logMessage = fmt.Sprintf("END - RESPONSE: Response: %+v", response)
	utils.LogRegister(c, logGetDebtsHandler, logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
