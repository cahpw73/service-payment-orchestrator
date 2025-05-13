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

const logSubCategoriesHandler = "Handler SubCategories"

type SubCategoryHandlerInterface interface {
	GetSubCategories(c *fiber.Ctx) error
}

type SubCategoryHandler struct {
	subCategoryService services.SubCategoryInterface
}

func NewSubCategoryHandler(subCategoryService services.SubCategoryInterface) SubCategoryHandlerInterface {
	return &SubCategoryHandler{subCategoryService: subCategoryService}
}

func (ct *SubCategoryHandler) GetSubCategories(c *fiber.Ctx) error {

	id := c.Params("id")

	utils.LogRegister(c, logSubCategoriesHandler, "INIT - PARAMS: {idCategory: "+id+"}")

	subCategories, err := ct.subCategoryService.GetSubCategories(id)
	if err != nil {

		utils.LogErrorRegister(c, logSubCategoriesHandler, err)
		errCode := errors.GetError("ERR009")

		return c.Status(fiber.StatusInternalServerError).JSON(models.Response[any]{
			ErrCode: errCode.Code,
			ErrMsg:  errCode.Message,
			Data:    nil,
		})
	}

	response := models.Response[[]responseMiddleware.SubCategoryResponseMiddleware]{
		ErrCode: "",
		ErrMsg:  "",
		Data:    subCategories,
	}

	logMessage := fmt.Sprintf("Response: %+v", response)
	utils.LogRegister(c, logSubCategoriesHandler, "END - RESPONSE: "+logMessage)

	return c.Status(fiber.StatusOK).JSON(response)
}
