package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"fmt"

	"github.com/binkkatal/go-coffee-pods/models"
	"github.com/labstack/echo"
)

// This function will return an array of pods which can be filtered 
// based on size and /or flavor provided as query params
func (env *Env) pods(c echo.Context) error {
	var flavorID int
	var sizeID int
	var err error
	var pods []models.Pod

	q := c.QueryParams()
	if q["flavor_id"] != nil {
		flavorID, err = strconv.Atoi(q["flavor_id"][0])
		if err != nil {
			invalidIDErr := fmt.Errorf("(%s) flavor_id is in an invalid format please provide an integer :description -> (%+v)", q["flavor_id"], err)
			log.Printf("Error: %v", invalidIDErr)
			return handleErr(c, http.StatusBadRequest, invalidIDErr)
		}
	}
	if q["size_id"] != nil {
		sizeID, err = strconv.Atoi(q["size_id"][0])
		if err != nil {
			invalidIDErr := fmt.Errorf("(%s) size_id is in an invalid format please provide an integer :description -> (%+v)", q["size_id"], err)
			log.Printf("Error: %v", invalidIDErr)
			return handleErr(c, http.StatusBadRequest, invalidIDErr)
		}
	}
	pods, err = env.DS.Pods(flavorID, sizeID)
	if err != nil {
		dbErr := fmt.Errorf("Unable to fetch Pods for flavor_id (%d) , size_id (%d) description-> (%+v)", flavorID, sizeID, err)
		log.Printf("Error %v", dbErr)
		return handleErr(c, http.StatusInternalServerError, dbErr)
	}

	data, err := json.Marshal(pods)
	if err != nil {
		if err != nil {
			parseError := fmt.Errorf("Error marshaling Pods description -> (%+v)", err)
			log.Printf("Error %v", parseError)
			return handleErr(c, http.StatusInternalServerError, parseError)
		}
	}
	return c.JSON(http.StatusOK, string(data))
}

// This function will return the cross sell pods based on pod id provided as query param

func (env *Env) crossSellPods(c echo.Context) error {
	var pods []models.Pod
	var podID int
	var err error

	q := c.QueryParams()
	if q["pod_id"] == nil {
		return handleErr(c, http.StatusBadRequest, errors.New("Please Provide a pod_id"))
	}
	podID, err = strconv.Atoi(q["pod_id"][0])
	if err != nil {
		invalidIDErr := fmt.Errorf("(%s) pod_id is in an invalid format please provide an integer :description -> (%+v)", q["pod_id"], err)
		log.Printf("Error: %v", invalidIDErr)
		return handleErr(c, http.StatusBadRequest, invalidIDErr)
	}
	pods, err = env.DS.CrossSellPods(podID)
	if err != nil {
		dbErr := fmt.Errorf("Unable to fetch CrossSellPods for pod_id (%d) description-> (%+v)", podID, err)
		log.Printf("Error %v", dbErr)
		return handleErr(c, http.StatusInternalServerError, dbErr)
	}

	data, err := json.Marshal(pods)
	if err != nil {
		parseError := fmt.Errorf("Error marshaling Pods description -> (%+v)", err)
		log.Printf("Error %v", parseError)
		return handleErr(c, http.StatusInternalServerError, parseError)
	}
	return c.JSON(http.StatusOK, string(data))
}
