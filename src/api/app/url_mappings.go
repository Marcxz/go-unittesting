package app

import controller "unitesttraining/src/api/controllers"

func mapUrls() {
	router.GET("/locations/countries/:country_id", controller.GetCountry)
}
