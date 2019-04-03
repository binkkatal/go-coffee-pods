package dao

import (
	"github.com/binkkatal/go-coffee-pods/models"
)

// Datastore is an interface with data access methods that will
// allow persistence.
type Datastore interface {
	CrossSellMachines(int) ([]models.Pod, error)
	CrossSellPods(int) ([]models.Pod, error)
	initDB() error
	Machines(int) ([]models.CoffeeMachine, error)
	Pods(int, int) ([]models.Pod, error)
	insertPod(string) error
	insertMachine(string) error
	parseProducts() error
}
