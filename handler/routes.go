package handler

import (
	"github.com/labstack/echo"
)

// InitRoutes intializes the router with its routes and returns
// a pointer to the Echo.
func InitRoutes(env *Env) *echo.Echo {
	e := echo.New()

	// Declare the api subrouter.
	api := e.Group("/api")
	// Register coffee machine routes.
	api.GET("/product/listOfCoffeeMachines", env.coffeeMachines)
	api.GET("/crosssell/coffeeMachines", env.crossSellCoffeeMachines)
	// Register pod routes.
	api.GET("/product/listOfCoffeePods", env.pods)
	api.GET("/crosssell/coffeePods", env.crossSellPods)

	return e
}
