package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	"github.com/utronframework/booking/controllers"
	"github.com/utronframework/booking/models"
)

func main() {
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}
	// resgister models
	app.Model.Register(&models.Account{})
	app.Model.LogMode(true)
	app.Model.AutoMigrateAll()

	// Register Controllers
	app.AddController(controllers.NewAccount)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
