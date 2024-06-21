package router

import (
	"github.com/bagusrexy/technical-test-troben/controller"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	healthCheck := new(controller.HealthController)
	router.GET("/", healthCheck.HealthCheck)

	imdb := new(controller.IMDBController)
	router.GET("/movie/detail", imdb.GetMoviesDetail)
	router.GET("/movie/search", imdb.GetMovieByKeyword)
	router.GET("/movie/popular", imdb.GetPopularMovie)
}
