package route

import (
	"assesement-test-MicroServices/src/auth-service/delivery/http"
	"assesement-test-MicroServices/src/auth-service/delivery/http/middleware"

	"github.com/gorilla/mux"
)

type ExposeRoute struct {
	Router       *mux.Router
	AccountRoute *AccountRoute
	OfficeRoute  *OfficeRoute
	ScreenRoute  *ScreenRoute
	RoleRoute    *RoleRoute
}

func NewExposeRoute(
	router *mux.Router,
	accountRoute *AccountRoute,
	officeRoute *OfficeRoute,
	screenRoute *ScreenRoute,
	roleRoute *RoleRoute,

) *ExposeRoute {
	rootRoute := &ExposeRoute{
		Router:       router,
		AccountRoute: accountRoute,
		OfficeRoute:  officeRoute,
		ScreenRoute:  screenRoute,
		RoleRoute:    roleRoute,
	}
	return rootRoute
}

func (exposeRoute *ExposeRoute) Register() {
	exposeRoute.AccountRoute.Register()
	exposeRoute.OfficeRoute.Register()
	exposeRoute.RoleRoute.Register()
	exposeRoute.ScreenRoute.Register()
}

// account route

type AccountRoute struct {
	Middleware        *middleware.AuthMiddleware
	Router            *mux.Router
	AccountController *http.ExposeController
}

func NewAccountRoute(router *mux.Router, accountController *http.ExposeController, middleware *middleware.AuthMiddleware) *AccountRoute {
	accountRoute := &AccountRoute{
		Router:            router.PathPrefix("/accounts").Subrouter(),
		AccountController: accountController,
		Middleware:        middleware,
	}
	return accountRoute
}

func (accountRoute *AccountRoute) Register() {
	accountRoute.Router.Use(accountRoute.Middleware.Middleware)
	accountRoute.Router.HandleFunc("/{id}", accountRoute.AccountController.DetailAccount).Methods("GET")
	accountRoute.Router.HandleFunc("/accountName/{accountName}", accountRoute.AccountController.GetOneByAccountName).Methods("GET")
	accountRoute.Router.HandleFunc("", accountRoute.AccountController.ListAccount).Methods("GET")
	accountRoute.Router.HandleFunc("/{id}", accountRoute.AccountController.UpdateAccount).Methods("PATCH")
	accountRoute.Router.HandleFunc("/{id}", accountRoute.AccountController.DeleteAccount).Methods("DELETE")
}

// office route

type OfficeRoute struct {
	Middleware       *middleware.AuthMiddleware
	Router           *mux.Router
	OfficeController *http.ExposeController
}

func NewOfficeRoute(router *mux.Router, officeController *http.ExposeController, middleware *middleware.AuthMiddleware) *OfficeRoute {
	officeRoute := &OfficeRoute{
		Router:           router.PathPrefix("/offices").Subrouter(),
		OfficeController: officeController,
		Middleware:       middleware,
	}
	return officeRoute
}

func (officeRoute *OfficeRoute) Register() {
	officeRoute.Router.Use(officeRoute.Middleware.Middleware)
	officeRoute.Router.HandleFunc("", officeRoute.OfficeController.CreateOffice).Methods("POST")
	officeRoute.Router.HandleFunc("", officeRoute.OfficeController.ListOffices).Methods("GET")
	officeRoute.Router.HandleFunc("/{id}", officeRoute.OfficeController.DetailOffice).Methods("GET")
	officeRoute.Router.HandleFunc("/{id}", officeRoute.OfficeController.DeleteOffice).Methods("DELETE")
	officeRoute.Router.HandleFunc("/{id}", officeRoute.OfficeController.UpdateOffice).Methods("PATCH")
}

// role route

type RoleRoute struct {
	Middleware     *middleware.AuthMiddleware
	Router         *mux.Router
	RoleController *http.ExposeController
}

func NewRoleRoute(router *mux.Router, roleController *http.ExposeController, middleware *middleware.AuthMiddleware) *RoleRoute {
	roleRoute := &RoleRoute{
		Router:         router.PathPrefix("/roles").Subrouter(),
		RoleController: roleController,
		Middleware:     middleware,
	}
	return roleRoute
}
func (roleRoute *RoleRoute) Register() {
	roleRoute.Router.Use(roleRoute.Middleware.Middleware)
	roleRoute.Router.HandleFunc("", roleRoute.RoleController.CreateRole).Methods("POST")
	roleRoute.Router.HandleFunc("", roleRoute.RoleController.ListRoles).Methods("GET")
	roleRoute.Router.HandleFunc("/{id}", roleRoute.RoleController.DetailRole).Methods("GET")
	roleRoute.Router.HandleFunc("/{id}", roleRoute.RoleController.DeleteRole).Methods("DELETE")
	roleRoute.Router.HandleFunc("/{id}", roleRoute.RoleController.UpdateRole).Methods("PATCH")
}

// screen route
type ScreenRoute struct {
	Middleware       *middleware.AuthMiddleware
	Router           *mux.Router
	ScreenController *http.ExposeController
}

func NewScreenRoute(router *mux.Router, ScreenController *http.ExposeController, middleware *middleware.AuthMiddleware) *ScreenRoute {
	ScreenRoute := &ScreenRoute{
		Router:           router.PathPrefix("/screens").Subrouter(),
		ScreenController: ScreenController,
		Middleware:       middleware,
	}
	return ScreenRoute
}

func (screenRoute *ScreenRoute) Register() {
	screenRoute.Router.Use(screenRoute.Middleware.Middleware)
	screenRoute.Router.HandleFunc("", screenRoute.ScreenController.CreateScreen).Methods("POST")
	screenRoute.Router.HandleFunc("", screenRoute.ScreenController.ListScreens).Methods("GET")
	screenRoute.Router.HandleFunc("/{id}", screenRoute.ScreenController.DetailScreen).Methods("GET")
	screenRoute.Router.HandleFunc("/{id}", screenRoute.ScreenController.DeleteScreen).Methods("DELETE")
	screenRoute.Router.HandleFunc("/{id}", screenRoute.ScreenController.UpdateScreen).Methods("PATCH")
}
