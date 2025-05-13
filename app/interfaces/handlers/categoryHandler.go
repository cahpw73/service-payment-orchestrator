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

const logCategoryHandler = "Handler Categories"

type CategoryHandlerInterface interface {
	GetCategories(c *fiber.Ctx) error
}

type CategoryHandler struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryHandler(categoryService services.CategoryServiceInterface) CategoryHandlerInterface {
	return &CategoryHandler{categoryService: categoryService}
}

func (ct *CategoryHandler) GetCategories(c *fiber.Ctx) error {

	utils.LogRegister(c, logCategoryHandler, "INIT")

	categories, err := ct.categoryService.GetCategories()
	if err != nil {

		utils.LogErrorRegister(c, logCategoryHandler, err)
		errCode := errors.GetError("ERR001")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[[]responseMiddleware.CategoryResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    categories,
	}

	logMessage := fmt.Sprintf("Response: %+v", response)
	utils.LogRegister(c, logCategoryHandler, "END - RESPONSE: "+logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
