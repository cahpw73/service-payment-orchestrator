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

const logCitiesHandler = "Handler Cities"

type CityHandlerInterface interface {
	GetCities(c *fiber.Ctx) error
}

type CityHandler struct {
	cityService services.CityServiceInterface
}

func NewCityHandler(cityService services.CityServiceInterface) CityHandlerInterface {
	return &CityHandler{cityService: cityService}
}

func (ct *CityHandler) GetCities(c *fiber.Ctx) error {

	id := c.Params("id")

	utils.LogRegister(c, logCitiesHandler, "INIT - PARAMS: {idSubCategory: "+id+"}")

	cities, err := ct.cityService.GetCities(id)
	if err != nil {

		utils.LogErrorRegister(c, logCitiesHandler, err)
		errCode := errors.GetError("ERR002")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[[]responseMiddleware.CityResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    cities,
	}

	logMessage := fmt.Sprintf("Response: %+v", response)
	utils.LogRegister(c, logCitiesHandler, "END - RESPONSE: "+logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
