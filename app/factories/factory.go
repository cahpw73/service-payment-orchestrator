package factories

import (
	"log"
	"service-payment-orchestrator/app/interfaces/handlers"
	"service-payment-orchestrator/app/interfaces/services"
	"service-payment-orchestrator/client/redis"
	"service-payment-orchestrator/config"
)

type Factory struct {
	RedisClient                 redis.RedisServiceInterface
	MiddlewareService           services.MiddlewareServiceInterface
	CityService                 services.CityServiceInterface
	CityHandler                 handlers.CityHandlerInterface
	CategoryService             services.CategoryServiceInterface
	CategoryHandler             handlers.CategoryHandlerInterface
	GetDebtService              services.GetDebtServiceInterface
	GetDebtHandler              handlers.GetDebtHandlerInterface
	PayDebtService              services.PayDebtServiceInterface
	PayDebtHandler              handlers.PayDebtHandlerInterface
	SearchCriteriaService       services.SearchCriteriaInterface
	SearchCriteriaHandler       handlers.SearchCriteriaHandlerInterface
	ServiceBySubCategoryService services.ServicesSubCategoryInterface
	ServiceBySubCategoryHandler handlers.ServicesBySubCategoryHandlerInterface
	SubCategoryService          services.SubCategoryInterface
	SubCategoryHandler          handlers.SubCategoryHandlerInterface
	AffiliationService          services.AffiliationInterface
	AffiliationHanlder          handlers.AffiliationHandlerInterface
}

func NewFactory(redisHost string) (*Factory, error) {

	config.LoadConfig()

	redisClient := redis.NewRedisService(redisHost)
	middlewareService := services.NewMiddlewareService(redisClient)
	cityService := services.NewCityService(redisClient, middlewareService)
	cityHandler := handlers.NewCityHandler(cityService)
	categoryService := services.NewCategoryService(redisClient, middlewareService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	getDebtService := services.NewGetDebtService(middlewareService)
	getDebtHandler := handlers.NewGetDebtHandler(getDebtService)
	payDebtService := services.NewPayDebtService(middlewareService)
	payDebtHandler := handlers.NewPayDebtHandler(payDebtService)
	searchCriteriaService := services.NewSearchCriteria(redisClient, middlewareService)
	searchCriteriaHandler := handlers.NewSearchCriteriaHandler(searchCriteriaService)
	serviceBySubCategoryService := services.NewServicesSubCategoryService(redisClient, middlewareService)
	serviceBySubCategoryHandler := handlers.NewServicesBySubCategoryHandler(serviceBySubCategoryService)
	subCategoryService := services.NewSubCategoryService(redisClient, middlewareService)
	subCategoryHandler := handlers.NewSubCategoryHandler(subCategoryService)
	affiliationService := services.NewAffiliationService(middlewareService)
	affiliationHandler := handlers.NewAffiliationHandler(affiliationService)

	return &Factory{
		RedisClient:                 redisClient,
		MiddlewareService:           middlewareService,
		CityService:                 cityService,
		CityHandler:                 cityHandler,
		CategoryService:             categoryService,
		CategoryHandler:             categoryHandler,
		GetDebtService:              getDebtService,
		GetDebtHandler:              getDebtHandler,
		PayDebtService:              payDebtService,
		PayDebtHandler:              payDebtHandler,
		SearchCriteriaService:       searchCriteriaService,
		SearchCriteriaHandler:       searchCriteriaHandler,
		ServiceBySubCategoryService: serviceBySubCategoryService,
		ServiceBySubCategoryHandler: serviceBySubCategoryHandler,
		SubCategoryService:          subCategoryService,
		SubCategoryHandler:          subCategoryHandler,
		AffiliationService:          affiliationService,
		AffiliationHanlder:          affiliationHandler,
	}, nil
}

// Método para cerrar conexiones, si es necesario
func (f *Factory) Close() {
	if err := f.RedisClient.Close(); err != nil {
		log.Fatalf("Error cerrando la conexión de Redis: %v", err)
	}
}
