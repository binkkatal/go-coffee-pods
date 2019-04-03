package dao

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/binkkatal/go-coffee-pods/models"
)

// DBApi is type which will implement Datastore and will be the receiver
// for our data access functions.
type DBApi struct {
	*sql.DB
}

// Defining assetsDir as a package scoped variable to make sure we can access it before environment is established
var assetsDir string

// NewDB returns an instance of the DB struct.
func NewDB() (*DBApi, error) {
	// gets assetsDir flag.
	assetsDir = flag.Lookup("assetsDir").Value.(flag.Getter).Get().(string)
	// Recreates data on build.
	if err := os.Remove(assetsDir + "/data.db"); err != nil {
		log.Printf("either you are creating this DB for the first time or assetDir is not set properly")
	}

	db, err := sql.Open("sqlite3", assetsDir+"/data.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	newDB := &DBApi{db}
	if err = newDB.initDB(); err != nil {
		return nil, err
	}
	return newDB, nil
}
 
// This function returns an array of pods and it filters the pods based on 
// flavorID and or sizeID

func (db *DBApi) Pods(flavorID int, sizeID int) ([]models.Pod, error) {
	var pods []models.Pod
	var rows *sql.Rows
	var err error

	switch {
	case flavorID != 0 && sizeID != 0:
		query := podQ + " WHERE Pod.FlavorID = ? AND Pod.SizeID = ?"
		rows, err = db.Query(query, flavorID, sizeID)
		if err != nil {
			return nil, err
		}
	case flavorID != 0:
		query := podQ + " WHERE Pod.FlavorID = ?"
		rows, err = db.Query(query, flavorID)
		if err != nil {
			return nil, err
		}
	case sizeID != 0:
		query := podQ + " WHERE Pod.SizeID = ?"
		rows, err = db.Query(query, sizeID)
		if err != nil {
			return nil, err
		}
	default:
		rows, err = db.Query(podQ)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		pod := new(models.Pod)
		err = rows.Scan(&pod.PodID, &pod.FlavorID, &pod.FlavorName,
			&pod.SizeID, &pod.SizeName, &pod.SKU, &pod.Quantity)
		if err != nil {
			return nil, err
		}
		pods = append(pods, *pod)
	}

	return pods, nil
}

// This function returns an array of different pack sizes for each flavor of
// coffee pods to be cross sold on the coffee pod page

func (db *DBApi) CrossSellPods(id int) ([]models.Pod, error) {
	var pods []models.Pod
	var pod models.Pod

	err := db.QueryRow(podQ+" WHERE PodID = ?", id).Scan(&pod.PodID,
		&pod.FlavorID, &pod.FlavorName, &pod.SizeID, &pod.SizeName,
		&pod.SKU, &pod.Quantity)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(podQ+" WHERE Pod.FlavorID = ? AND Pod.SizeID = ? AND PodID != ?",
		pod.FlavorID, pod.SizeID, pod.PodID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pod := new(models.Pod)
		if err = rows.Scan(&pod.PodID, &pod.FlavorID, &pod.FlavorName,
			&pod.SizeID, &pod.SizeName, &pod.SKU, &pod.Quantity); err != nil {
			return nil, err
		}
		pods = append(pods, *pod)
	}

	return pods, nil
}

// This function returns an array of coffee machines based on the value of sizeID passed

func (db *DBApi) Machines(sizeID int) ([]models.CoffeeMachine, error) {
	var machines []models.CoffeeMachine
	var rows *sql.Rows
	var err error

	// Get the corresponding machine.
	if sizeID == 0 {
		rows, err = db.Query(coffeeMachineQ)
	} else {
		rows, err = db.Query(coffeeMachineQ+" WHERE CoffeeMachine.SizeID = ?", sizeID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		m := new(models.CoffeeMachine)
		err = rows.Scan(&m.CoffeeMachineID, &m.SizeID, &m.SizeName,
			&m.SKU, &m.ModelID, &m.ModelName, &m.WaterLine)
		if err != nil {
			return nil, err
		}
		machines = append(machines, *m)
	}

	return machines, nil
}

// This function returns an array of pods to be cross sold on the coffee machine pages.

func (db *DBApi) CrossSellMachines(id int) ([]models.Pod, error) {
	var machine models.CoffeeMachine
	var pods []models.Pod

	// Get the corresponding machine.
	err := db.QueryRow(coffeeMachineQ+" where CoffeeMachineId = ?", id).
		Scan(&machine.CoffeeMachineID, &machine.SizeID, &machine.SizeName,
			&machine.SKU, &machine.ModelID, &machine.ModelName,
			&machine.WaterLine)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(machineCrossQ, machine.SizeID)
	defer rows.Close()

	for rows.Next() {
		pod := new(models.Pod)
		err = rows.Scan(&pod.PodID, &pod.FlavorID, &pod.FlavorName,
			&pod.SizeID, &pod.SizeName, &pod.SKU, &pod.Quantity)
		if err != nil {
			return nil, err
		}
		pods = append(pods, *pod)
	}

	return pods, nil
}
