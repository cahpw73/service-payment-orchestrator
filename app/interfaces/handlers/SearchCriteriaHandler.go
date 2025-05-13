package handlers

import (
	"fmt"
	"service-payment-orchestrator/app/errors"
	"service-payment-orchestrator/app/interfaces/services"
	"service-payment-orchestrator/app/models"
	responseMiddleware "service-payment-orchestrator/app/models/ResponseMiddleware"
	"service-payment-orchestrator/utils"

	"github.com/gofiber/fiber/v2"
)

const logSeachCriteriaHandler = "Handler Search Criteria"

type SearchCriteriaHandlerInterface interface {
	GetSearchCriteria(c *fiber.Ctx) error
}

type SearchCriteriaHandler struct {
	searchCriteria services.SearchCriteriaInterface
}

func NewSearchCriteriaHandler(searchCriteria services.SearchCriteriaInterface) SearchCriteriaHandlerInterface {
	return &SearchCriteriaHandler{searchCriteria: searchCriteria}
}

func (ct *SearchCriteriaHandler) GetSearchCriteria(c *fiber.Ctx) error {

	serviceCode := c.Params("serviceCode")

	utils.LogRegister(c, logSeachCriteriaHandler, fmt.Sprintf("INIT - PARAMS: {serviceCode: %s}", serviceCode))

	searchCriteria, err := ct.searchCriteria.GetSearchCriteria(serviceCode)
	if err != nil {

		utils.LogErrorRegister(c, logSeachCriteriaHandler, err)
		errCode := errors.GetError("ERR007")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[responseMiddleware.SearchCriteriaResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    searchCriteria,
	}

	utils.LogRegister(c, logSeachCriteriaHandler, fmt.Sprintf("END - RESPONSE: Response: %+v", response))

	return c.Status(fiber.StatusOK).JSON(response)
}
