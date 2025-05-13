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

const logServicesBySubCategoryHandler = "Handler Services By SubCategories"

type ServicesBySubCategoryHandlerInterface interface {
	GetServicesBySubCategory(c *fiber.Ctx) error
}

type ServicesBySubCategoryHandler struct {
	servicesBySubCategoryService services.ServicesSubCategoryInterface
}

func NewServicesBySubCategoryHandler(servicesBySubCategoryService services.ServicesSubCategoryInterface) ServicesBySubCategoryHandlerInterface {
	return &ServicesBySubCategoryHandler{servicesBySubCategoryService: servicesBySubCategoryService}
}

func (ct *ServicesBySubCategoryHandler) GetServicesBySubCategory(c *fiber.Ctx) error {

	subCategoryID := c.Query("subCategory")
	cityID := c.Query("city")

	logMessage := fmt.Sprintf("INIT - PARAMS: {subCategoryID: %s, cityID %s }", subCategoryID, cityID)
	utils.LogRegister(c, logServicesBySubCategoryHandler, logMessage)

	services, err := ct.servicesBySubCategoryService.GetServicesSubCategory(subCategoryID, cityID)
	if err != nil {

		utils.LogErrorRegister(c, logServicesBySubCategoryHandler, err)
		errCode := errors.GetError("ERR008")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[[]responseMiddleware.ServiceBySubCategoryResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    services,
	}

	logMessage = fmt.Sprintf("END - RESPONSE: Response: %+v", response)
	utils.LogRegister(c, logServicesBySubCategoryHandler, logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
