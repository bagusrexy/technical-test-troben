package controller

import (
	"net/http"

	"github.com/bagusrexy/technical-test-troben/service"
	"github.com/gin-gonic/gin"
)

type IMDBController struct{}

func (m *IMDBController) GetMoviesDetail(c *gin.Context) {
	var service service.Service

	id := c.Query("id")
	movie, err := service.GetMoviesByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (m *IMDBController) GetPopularMovie(c *gin.Context) {
	var service service.Service

	id := c.DefaultQuery("year", "2024")
	if id == "" {
		id = "2024"
	}

	movie, err := service.GetPopularMovie(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
