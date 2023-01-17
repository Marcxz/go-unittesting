package controller

import (
	"net/http"
	"unitesttraining/src/api/services"

	"github.com/gin-gonic/gin"
)

func GetCountry(c *gin.Context) {
	ls := services.NewLocationService()

	country, err := ls.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, country)
}
