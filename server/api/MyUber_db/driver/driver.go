/* 
* Generated by
* 
*      _____ _          __  __      _     _
*     / ____| |        / _|/ _|    | |   | |
*    | (___ | | ____ _| |_| |_ ___ | | __| | ___ _ __
*     \___ \| |/ / _` |  _|  _/ _ \| |/ _` |/ _ \ '__|
*     ____) |   < (_| | | | || (_) | | (_| |  __/ |
*    |_____/|_|\_\__,_|_| |_| \___/|_|\__,_|\___|_|
*
* The code generator that works in many programming languages
*
*			https://www.skaffolder.com
*
*
* You can generate the code from the command-line
*       https://npmjs.com/package/skaffolder-cli
*
*       npm install -g skaffodler-cli
*
*   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *
*
* To remove this comment please upgrade your plan here: 
*      https://app.skaffolder.com/#!/upgrade
*
* Or get up to 70% discount sharing your unique link:
*       https://app.skaffolder.com/#!/register?friend=5d525f5a7b158e50f28eb4af
*
* You will get 10% discount for each one of your friends
* 
*/
package driver

import (
	"skaffolder/MyUber/config"
	"skaffolder/MyUber/security"

	"github.com/go-chi/chi"
)

type Config struct {
	*config.Config
}

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

// Routes
func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()

	// start routing

	router.Group(func(router chi.Router) { 
		router.Use(security.HasRole())
		router.Post("/", config.create)
	})

	router.Group(func(router chi.Router) { 
		router.Use(security.HasRole())
		router.Delete("/{id}", config.delete)
	})

	router.Group(func(router chi.Router) { 
		router.Use(security.HasRole())
		router.Get("/{id}", config.get)
	})

	router.Group(func(router chi.Router) { 
		router.Use(security.HasRole())
		router.Get("/", config.list)
	})

	router.Group(func(router chi.Router) { 
		router.Use(security.HasRole())
		router.Post("/{id}", config.update)
	})

	// end routing

	// Write here your custom APIs
	// EXAMPLE :

	/**
	router.Group(func(router chi.Router) {
		router.Get("/", config.listCustom) // Create the listCustom method in this file
	})
	*/

	return router
}
