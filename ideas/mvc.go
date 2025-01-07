package main

import (
	"fmt"
	"sync"
)

// Model - Singleton Model
type Model struct {
	data string
}

var instance *Model
var once sync.Once

// GetInstance - Get the singleton instance of the Model
func GetInstance() *Model {
	once.Do(func() {
		instance = &Model{}
	})
	return instance
}

// Controller - Controller for the MVC
type Controller struct {
	model *Model
}

// NewController - Create a new Controller
func NewController() *Controller {
	return &Controller{
		model: GetInstance(),
	}
}

// View - View for the MVC
type View struct{}

// Display - Display the data from the model
func (v *View) Display(data string) {
	fmt.Println("Data:", data)
}

// Main function
func main() {
	controller := NewController()
	view := &View{}

	// Set data in the model
	controller.model.data = "Hello, MVC!"

	// Display data using the view
	view.Display(controller.model.data)
}
