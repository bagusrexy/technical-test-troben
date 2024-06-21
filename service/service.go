package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bagusrexy/technical-test-troben/models"
	"github.com/bagusrexy/technical-test-troben/utils"
	"github.com/gin-gonic/gin"
)

type Service struct{}

func (s *Service) GetMoviesByID(c *gin.Context, id string) (models.Movie, error) {
	var movie models.Movie

	url := fmt.Sprintf("https://www.omdbapi.com/?i=%s&apikey=%s", id, os.Getenv("API_KEY"))
	resp, err := utils.MakeRequest(url)
	if err != nil {
		return models.Movie{}, err
	}
	json.Unmarshal(resp, &movie)
	return movie, nil
}

func (s *Service) GetPopularMovie(c *gin.Context) ([]models.Movie, error) {
	var movie []models.Movie

	url := fmt.Sprintf("https://www.omdbapi.com/?y=%s&apikey=%s", os.Getenv("YEAR_OF_RELEASE"), os.Getenv("API_KEY"))
	resp, err := utils.MakeRequest(url)
	if err != nil {
		return []models.Movie{}, err
	}
	json.Unmarshal(resp, &movie)
	return movie, nil
}

func (s *Service) SearchMovies(c *gin.Context, title, year, movieType string) ([]models.Movie, error) {
	var movies []models.Movie
	var url string

	url = fmt.Sprintf("https://www.omdbapi.com/?apikey=%s", os.Getenv("API_KEY"))
	if title != "" {
		url = fmt.Sprintf("%s&s=%s", url, title)
	}
	if year != "" {
		url = fmt.Sprintf("%s&y=%s", url, year)
	}
	if movieType != "" {
		url = fmt.Sprintf("%s&type=%s", url, movieType)
	}

	resp, err := utils.MakeRequest(url)
	if err != nil {
		return []models.Movie{}, err
	}
	json.Unmarshal(resp, &movies)
	return movies, nil

	return movies, nil
}
