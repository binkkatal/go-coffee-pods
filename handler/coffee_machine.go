package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// This function return the list of coffee machines based on the 
// size id provided as query param

func (env *Env) coffeeMachines(c echo.Context) error {
	var id int
	var err error
	q := c.QueryParams()
	if q["size_id"] != nil {
		id, err = strconv.Atoi(q["size_id"][0])
		if err != nil {
			invalidIDErr := fmt.Errorf("(%s) id is in an invalid format please provide an integer :description -> (%+v)", q["size_id"], err)
			return handleErr(c, http.StatusBadRequest, invalidIDErr)
		}
	}

	machines, err := env.DS.Machines(id)
	if err != nil {
		dbErr := fmt.Errorf("Unable to fetch machines for id (%d) description-> (%+v)", id, err)
		return handleErr(c, http.StatusBadRequest, dbErr)
	}
	data, err := json.Marshal(machines)
	if err != nil {
		parseError := fmt.Errorf("Error marshaling machines description -> (%+v)", err)
		return handleErr(c, http.StatusBadRequest, parseError)
	}
	return c.JSON(http.StatusOK, string(data))
}


// This function will return the list of pods, 
// that can be cross sold based on the coffee machine id provided as a query param

func (env *Env) crossSellCoffeeMachines(c echo.Context) error {
	q := c.QueryParams()
	fmt.Println("q",q);
	if q["coffee_machine_id"] == nil || len(q["coffee_machine_id"]) > 1 {
		return handleErr(c, http.StatusBadRequest, errors.New("please provide one machine id"))
	}
	id, err := strconv.Atoi(q["coffee_machine_id"][0])
	if err != nil {
		invalidIDErr := fmt.Errorf("(%s) id is in an invalid format please provide an integer :description -> (%+v)", q["coffee_machine_id"], err)
		return handleErr(c, http.StatusBadRequest, invalidIDErr)
	}
	pods, err := env.DS.CrossSellMachines(id)
	if err != nil {
		dbErr := fmt.Errorf("Unable to fetch Pods for coffee_machine_id (%d) description-> (%+v)", id, err)
		return handleErr(c, http.StatusInternalServerError, dbErr)
	}
	data, err := json.Marshal(pods)
	if err != nil {
		parseError := fmt.Errorf("Error marshaling Pods description -> (%+v)", err)
		return handleErr(c, http.StatusInternalServerError, parseError)
	}

	return c.JSON(http.StatusOK, string(data))
}

func handleErr(c echo.Context, statusCode int, err error) error {
	log.Printf("Error: %v", err)
	return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
}
