package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/bagusrexy/technical-test-troben/models"
	"github.com/bagusrexy/technical-test-troben/repositories"
	"github.com/bagusrexy/technical-test-troben/utils"

	"github.com/gin-gonic/gin"
)

type Service struct{}

// var ctx = context.Background()

func (s *Service) GetMoviesByID(c *gin.Context, id string) (models.Movie, error) {
	var movie models.Movie

	cachedMovie, err := repositories.GetCache(id)
	if err == nil {
		fmt.Println("===============DAPAT DATA DARI REDIS============")
		json.Unmarshal(cachedMovie, &movie)
		return movie, nil
	}
	fmt.Println("===============TIDAK DAPAT DATA DARI REDIS============")

	url := fmt.Sprintf("https://www.omdbapi.com/?i=%s&apikey=%s", id, os.Getenv("API_KEY"))
	resp, err := utils.MakeRequest(url)
	if err != nil {
		return models.Movie{}, err
	}

	var temp map[string]interface{}
	err = json.Unmarshal(resp, &temp)
	if err != nil {
		return models.Movie{}, err
	}

	if temp["Response"] == "False" {
		return models.Movie{}, fmt.Errorf("error OMDB API: %v", temp["Error"])
	}

	err = json.Unmarshal(resp, &movie)
	if err != nil {
		return models.Movie{}, err
	}

	movieJSON, _ := json.Marshal(movie)
	repositories.SetCache(id, movieJSON, 10*time.Minute)
	return movie, nil
}

func (s *Service) GetPopularMovie(c *gin.Context, title string) ([]models.MovieSearch, error) {
	var movies []models.MovieSearch

	apiURL := "https://www.omdbapi.com/"

	params := url.Values{}
	params.Add("apikey", os.Getenv("API_KEY"))
	params.Add("y", os.Getenv("YEAR_OF_RELEASE"))
	params.Add("type", os.Getenv("TYPE"))
	if title != "" {
		params.Add("s", title)
	}

	url := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	resp, err := utils.MakeRequest(url)
	if err != nil {
		return nil, err
	}

	type OMDBResponse struct {
		Search       []models.MovieSearch `json:"Search"`
		TotalResults string               `json:"totalResults"`
		Response     string               `json:"Response"`
		Error        string               `json:"Error"`
	}

	var omdbResp OMDBResponse
	if err := json.Unmarshal(resp, &omdbResp); err != nil {
		return nil, err
	}

	if omdbResp.Response != "True" {
		return nil, fmt.Errorf("error OMDB API: %v", omdbResp.Error)
	}

	for _, movie := range omdbResp.Search {
		if movie.Image != "N/A" {
			movies = append(movies, models.MovieSearch{
				ID:    movie.ID,
				Title: movie.Title,
				Year:  movie.Year,
				Image: movie.Image,
			})
		}
	}

	return movies, nil
}

func (s *Service) SearchMovies(c *gin.Context, title, year, movieType string) ([]models.MovieSearch, error) {
	var movies []models.MovieSearch
	apiURL := "https://www.omdbapi.com/"

	params := url.Values{}
	params.Add("apikey", os.Getenv("API_KEY"))
	if title != "" {
		params.Add("s", title)
	}
	if year != "" {
		params.Add("y", year)
	}
	if movieType != "" {
		params.Add("type", movieType)
	}

	url := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	resp, err := utils.MakeRequest(url)
	if err != nil {
		return nil, err
	}

	type OMDBResponse struct {
		Search       []models.MovieSearch `json:"Search"`
		TotalResults string               `json:"totalResults"`
		Response     string               `json:"Response"`
		Error        string               `json:"Error"`
	}

	var omdbResp OMDBResponse
	if err := json.Unmarshal(resp, &omdbResp); err != nil {
		return nil, err
	}

	if omdbResp.Response != "True" {
		return nil, fmt.Errorf("error OMDB API: %v", omdbResp.Error)
	}

	for _, movie := range omdbResp.Search {
		movies = append(movies, models.MovieSearch{
			ID:    movie.ID,
			Title: movie.Title,
			Year:  movie.Year,
			Image: movie.Image,
		})
	}

	return movies, nil
}
