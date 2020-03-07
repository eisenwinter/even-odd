package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	_ "github.com/eisenwinter/evenodd/docs"
	service "github.com/eisenwinter/evenodd/grpc"
	"github.com/eisenwinter/evenodd/numbergen"
	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes/empty"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
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

type grpcService struct{}

func (g *grpcService) Even(context.Context, *empty.Empty) (*service.NumberResponse, error) {
	return &service.NumberResponse{
		Value: gen.Even(),
	}, nil
}

func (g *grpcService) Odd(context.Context, *empty.Empty) (*service.NumberResponse, error) {
	return &service.NumberResponse{
		Value: gen.Odd(),
	}, nil
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

func getHTTPPort() int {
	if value, exists := os.LookupEnv("EVEN_ODD_HTTP_PORT"); exists {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}
	return 8080
}

func getGRPCPort() int {
	if value, exists := os.LookupEnv("EVEN_ODD_GRPC_PORT"); exists {
		if value, err := strconv.Atoi(value); err == nil {
			return value
		}
	}
	return 8081
}

// @title even-odd API
// @version 1.0
// @description This api supplied even or odd numbers.

// @license.name MIT
// @license.url https://github.com/eisenwinter/even-odd/blob/master/LICENSE

// @host localhost:8080
// @BasePath /

func runGrpcService() {
	addr := fmt.Sprintf(":%d", getGRPCPort())
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		Logger.Println("Skipping GRPC Setup - unable to open TCP port")
		return
	}
	Logger.Printf("Starting even-odd GRPC on %s", addr)
	grpcServer := grpc.NewServer()
	service.RegisterEvenOddServiceServer(grpcServer, &grpcService{})
	grpcServer.Serve(lis)
}

func main() {
	gen = numbergen.CreateNumberGen()

	go runGrpcService()

	addr := fmt.Sprintf(":%d", getHTTPPort())
	Logger.Printf("Starting even-odd REST on %s", addr)
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	r.Route("/rest", func(r chi.Router) {
		r.Get("/even", evenHandler)
		r.Get("/odd", oddHandler)
	})
	log.Fatal(http.ListenAndServe(addr, r))
}
