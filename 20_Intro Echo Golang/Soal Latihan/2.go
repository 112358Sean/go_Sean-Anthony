package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Movie2 struct {
	Title	string `json:"Title"`
	Year	string `json:"Year"`
	imdbID	string `json:imdbID`
	Type	string `json:Type`
	Poster	string `json:Poster`
}

func getMovie2(c echo.Context) error {
	Page := c.Param("Page")
	Search := c.Param("Search")

	resp, err := http.Get(fmt.Sprintf("https://www.omdbapi.com/?apikey=a818034f&page=%s&s=%s", Page, Search))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to fetch movie data",
		})
	}
	defer resp.Body.Close()

	var movie Movie2
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to decode movie data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"title": movie.Title,
		"year":  movie.Year,
		"imdbID": movie.imdbID,
		"type": movie.Type,
		"poster": movie.Poster,
	})
}

func main() {
	e := echo.New()

	e.GET("/search", getMovie2)

	e.Logger.Fatal(e.Start(":8080"))
}
