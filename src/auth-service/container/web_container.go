package container

import (
	"assesement-test-MicroServices/src/auth-service/config"
	"assesement-test-MicroServices/src/auth-service/delivery/grpc/client"
	httpdelivery "assesement-test-MicroServices/src/auth-service/delivery/http"
	"assesement-test-MicroServices/src/auth-service/delivery/http/middleware"
	"assesement-test-MicroServices/src/auth-service/delivery/http/route"
	"assesement-test-MicroServices/src/auth-service/repository"
	"assesement-test-MicroServices/src/auth-service/use_case"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type WebContainer struct {
	Env        *config.EnvConfig
	AuthDB     *config.DatabaseConfig
	Repository *RepositoryContainer
	UseCase    *UseCaseContainer
	Controller *ControllerContainer
	Route      *route.RootRoute
}

func NewWebContainer() *WebContainer {
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	authDBConfig := config.NewDBConfig(envConfig)

	authRepository := repository.NewAuthRepository()
	repositoryContainer := NewRepositoryContainer(authRepository)

	accountUrl := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.AccountPort,
	)
	roleUrl := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.RolePort,
	)
	officeUrl := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.OfficePort,
	)
	screenUrl := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.ScreenPort,
	)

	initAccountClient := client.InitAccountServiceClient(accountUrl)
	initRoleClient := client.InitRoleServiceClient(roleUrl)
	initOfficeClient := client.InitOfficeServiceClient(officeUrl)
	initScreenClient := client.InitScreenServiceClient(screenUrl)
	authUseCase := use_case.NewAuthUseCase(authDBConfig, authRepository, envConfig, &initAccountClient)
	exposeUseCase := use_case.NewExposeUseCase(authDBConfig, authRepository, envConfig, &initAccountClient, &initRoleClient, &initOfficeClient, &initScreenClient)

	useCaseContainer := NewUseCaseContainer(authUseCase, exposeUseCase)

	authController := httpdelivery.NewAuthController(authUseCase, exposeUseCase)
	exposeController := httpdelivery.NewExposeController(exposeUseCase)

	controllerContainer := NewControllerContainer(authController, exposeController)

	router := mux.NewRouter()
	authMiddleware := middleware.NewAuthMiddleware(*authRepository, authDBConfig)
	authRoute := route.NewAuthRoute(router, authController)
	// expose route
	accountRoute := route.NewUserRoute(router, exposeController, authMiddleware)
	accountRoute := route.NewAccountRoute(router, exposeController, authMiddleware)
	categoryRoute := route.NewCategoryRoute(router, exposeController, authMiddleware)
	roleRoute := route.NewRoleRoute(router, exposeController, authMiddleware)

	rootRoute := route.NewRootRoute(
		router,
		authRoute,
	)
	exposeRoute := route.NewExposeRoute(
		router,
		accountRoute,
		accountRoute,
		categoryRoute,
		roleRoute,
	)

	rootRoute.Register()
	exposeRoute.Register()

	webContainer := &WebContainer{
		Env:        envConfig,
		AuthDB:     authDBConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Controller: controllerContainer,
		Route:      rootRoute,
	}

	return webContainer
}
