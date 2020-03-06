package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"./numbergen"
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

func main() {
	os.Getenv("EVEN_ODD_PORT")
	gen = numbergen.CreateNumberGen()
	http.HandleFunc("/even", evenHandler)
	http.HandleFunc("/odd", oddHandler)
	addr := fmt.Sprintf(":%d", getPort())
	Logger.Printf("Starting even-odd on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
