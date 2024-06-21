package controller

import (
	"fmt"
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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success",
		"data": movie})
}

func (m *IMDBController) GetPopularMovie(c *gin.Context) {
	var service service.Service

	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "title is required"})
		return
	}

	movie, err := service.GetPopularMovie(c, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success",
		"data": movie})
}

func (m *IMDBController) GetMovieByKeyword(c *gin.Context) {
	var service service.Service

	title := c.DefaultQuery("title", "")
	year := c.DefaultQuery("year", "2024")
	movieType := c.DefaultQuery("movie", "movie")
	validTypes := []string{"movie", "series", "episode"}

	isValidType := false
	for _, t := range validTypes {
		if movieType == t {
			isValidType = true
			break
		}
	}
	if !isValidType {
		msg := fmt.Sprintf("Invalid movie type: %s", movieType)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	movie, err := service.SearchMovies(c, title, year, movieType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success",
		"search": movie})
}
