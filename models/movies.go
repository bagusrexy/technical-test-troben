package models

type Movie struct {
	ID       string `json:"imdbID"`
	Title    string `json:"title"`
	Year     string `json:"Year"`
	Rated    string `json:"rated"`
	Released string `json:"released"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Actors   string `json:"actors"`
	Plot     string `json:"plot"`
	Image    string `json:"poster"`
}

type MovieSearch struct {
	ID    string `json:"imdbID"`
	Title string `json:"title"`
	Year  string `json:"Year"`
	Image string `json:"poster"`
}
