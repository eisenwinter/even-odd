package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "./docs"
	"./numbergen"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	Logger *log.Logger
	gen    numbergen.NumberGen
)

func init() {
	Logger = log.New(os.Stdout,
		"LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type numberResponse struct {
	Number int64 `json:"value"`
}

// EvenHandler godoc
// @Summary Returns a even number
// @Description Returns a even number
// @ID even
// @Accept  json
// @Produce  json
// @Success 200 {object} numberResponse
// @Router /rest/even [get]
func evenHandler(w http.ResponseWriter, r *http.Request) {
	res := numberResponse{
		Number: gen.Even(),
	}
	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// OddHandler godoc
// @Summary Returns a odd number
// @Description Returns a odd number
// @ID odd
// @Accept  json
// @Produce  json
// @Success 200 {object} numberResponse
// @Router /rest/odd [get]
func oddHandler(w http.ResponseWriter, r *http.Request) {
	res := numberResponse{
		Number: gen.Odd(),
	}
	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func getPort() int {
	if value, exists := os.LookupEnv("EVEN_ODD_PORT"); exists {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}
	return 8080
}

// @title even-odd API
// @version 1.0
// @description This api supplied even or odd numbers.

// @license.name MIT
// @license.url https://github.com/eisenwinter/even-odd/blob/master/LICENSE

// @host localhost:8080
// @BasePath /

func main() {
	os.Getenv("EVEN_ODD_PORT")
	gen = numbergen.CreateNumberGen()
	addr := fmt.Sprintf(":%d", getPort())
	Logger.Printf("Starting even-odd on %s", addr)
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))
	r.Route("/rest", func(r chi.Router) {
		r.Get("/even", evenHandler)
		r.Get("/odd", oddHandler)
	})
	log.Fatal(http.ListenAndServe(addr, r))
}
