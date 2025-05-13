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

const logPayDebtHandler = "Handler Pay Debts"

type PayDebtHandlerInterface interface {
	PostPayDebt(c *fiber.Ctx) error
}

type PayDebtHandler struct {
	payDebtService services.PayDebtServiceInterface
}

func NewPayDebtHandler(payDebtService services.PayDebtServiceInterface) PayDebtHandlerInterface {
	return &PayDebtHandler{payDebtService: payDebtService}
}

func (ct *PayDebtHandler) PostPayDebt(c *fiber.Ctx) error {

	var payRequest requestMiddleware.PayDebtRequestMiddleware

	if err := c.BodyParser(&payRequest); err != nil {

		errCode := errors.GetError("ERR005")

		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	logMessage := fmt.Sprintf("INIT - PARAMS: {request pay debt: %+v }", payRequest)
	utils.LogRegister(c, logPayDebtHandler, logMessage)

	services, err := ct.payDebtService.PostPayDebt(payRequest)
	if err != nil {

		utils.LogErrorRegister(c, logPayDebtHandler, err)
		errCode := errors.GetError("ERR006")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.TransactionDataResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    services,
	}

	logMessage = fmt.Sprintf("END - RESPONSE: Response: %+v", response)
	utils.LogRegister(c, logPayDebtHandler, logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
