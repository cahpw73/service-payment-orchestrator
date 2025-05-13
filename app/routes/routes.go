package routes

import (
	"service-payment-orchestrator/app/factories"

	"service-payment-orchestrator/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, appFactory factories.Factory) {

	app.Use(middleware.RequestIDMiddleware)

	app.Get("/servicepayment/v1/categories", middleware.RequestIDMiddleware, appFactory.CategoryHandler.GetCategories)
	app.Get("/servicepayment/v1/subcategories/:id", middleware.RequestIDMiddleware, appFactory.SubCategoryHandler.GetSubCategories)
	app.Get("/servicepayment/v1/cities/:id", middleware.RequestIDMiddleware, appFactory.CityHandler.GetCities)
	app.Get("/servicepayment/v1/services", middleware.RequestIDMiddleware, appFactory.ServiceBySubCategoryHandler.GetServicesBySubCategory)
	app.Get("/servicepayment/v1/search-criteria/:serviceCode", middleware.RequestIDMiddleware, appFactory.SearchCriteriaHandler.GetSearchCriteria)
	app.Get("/servicepayment/v1/affiliation/:accountNumber", middleware.RequestIDMiddleware, appFactory.AffiliationHanlder.GetAffiliationByAccount)

	app.Post("/servicepayment/v1/validate-affiliation", middleware.RequestIDMiddleware, appFactory.AffiliationHanlder.PostValidateAffiliation)
	app.Post("/servicepayment/v1/services/debts", middleware.RequestIDMiddleware, appFactory.GetDebtHandler.PostDebtConsultation)
	app.Post("/servicepayment/v1/pay-debt", middleware.RequestIDMiddleware, appFactory.PayDebtHandler.PostPayDebt)
	app.Post("/servicepayment/v1/services/affiliation", middleware.RequestIDMiddleware, appFactory.AffiliationHanlder.PostSaveAffiliation)
	app.Post("/servicepayment/v1/services/affiliation-update", middleware.RequestIDMiddleware, appFactory.AffiliationHanlder.PostUpdateAffiliation)

	app.Delete("/servicepayment/v1/affiliation/:accountNumber/:affiliationCode", middleware.RequestIDMiddleware, appFactory.AffiliationHanlder.DeleteAffiliation)
}
