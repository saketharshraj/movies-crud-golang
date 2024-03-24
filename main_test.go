package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	return r
}

func TestGetMoviesEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/movies", nil)
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteMovieEndpoint(t *testing.T) {
	// Adding a movie to delete
	movies = append(movies, Movie{ID: "2", Title: "To Delete", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	request, _ := http.NewRequest("DELETE", "/movies/2", nil)
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Check if the movie was deleted
	request, _ = http.NewRequest("GET", "/movies/2", nil)
	response = executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
	assert.NotContains(t, response.Body.String(), "To Delete")
}

func TestCreateMovieEndpoint(t *testing.T) {
	var jsonStr = []byte(`{"title":"Test Movie","director":{"firstname":"Jane","lastname":"Doe"}}`)
	request, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(jsonStr))
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateMovieEndpoint(t *testing.T) {
	var jsonStr = []byte(`{"title":"Updated Movie","director":{"firstname":"Jane","lastname":"Doe"}}`)
	request, _ := http.NewRequest("PUT", "/movies/1", bytes.NewBuffer(jsonStr))
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual)
}
