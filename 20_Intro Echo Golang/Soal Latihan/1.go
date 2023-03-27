package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	Title	string `json:"Title"`
	Year	string `json:"Year"`
	Rated	string `json:"Rated"`
	Released	string `json:"Released"`
	Runtime	string `json:"Runtime"`
	Genre	string `json:"Genre"`
	ImdbID	string `json:"imdbID"`
}

func getMovie(c echo.Context) error {
	imdbID := c.Param("imdbid")

	resp, err := http.Get(fmt.Sprintf("https://www.omdbapi.com/?apikey=a818034f&i=%s", imdbID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to fetch movie data",
		})
	}
	defer resp.Body.Close()

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to decode movie data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"title": movie.Title,
		"year":  movie.Year,
		"rated": movie.Rated,
		"released": movie.Released,
		"runtime": movie.Runtime,
		"genre": movie.Genre,
	})
}

func main() {
	e := echo.New()

	e.GET("/:imdbid", getMovie)

	e.Logger.Fatal(e.Start(":8080"))
}
