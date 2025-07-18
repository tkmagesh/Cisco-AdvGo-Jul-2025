package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Cost  float64 `json:"cost"`
	Units int     `json:"-"`
}

var products = []Product{
	{101, "Pen", 10, 5},
	{102, "Pencil", 5, 20},
	{103, "Marker", 50, 10},
}

type AppServer struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (appServer *AppServer) AddRoute(path string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.routes[path] = handler
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := appServer.routes[r.URL.Path]; handler != nil {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

// Application specific handlers

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	timeout := rand.Intn(5)
	log.Printf("[IndexHandler] processing the request, will take (%d)s....\n", timeout)
	time.Sleep(time.Duration(timeout) * time.Second)
	select {
	case <-r.Context().Done():
		log.Println("[IndexHandler] stop processing the request due to timeout....")
	default:
		fmt.Fprintln(w, "Hello, World!")
	}
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// fmt.Fprintln(w, "List of the products are on the way!")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of the customers are on the way!")
}

// Middlewares
// log middleware
func logMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func timeoutMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rCtx := r.Context()
		timeoutCtx, cancel := context.WithTimeout(rCtx, 3*time.Second)
		defer cancel()
		reqWithTimeout := r.WithContext(timeoutCtx)
		handler(w, reqWithTimeout)
		if reqWithTimeout.Context().Err() == context.DeadlineExceeded {
			http.Error(w, "request timedout", http.StatusRequestTimeout)
		}
	}
}

func main() {
	go func() {
		var q string
		for {
			if fmt.Scanln(&q); q == "list" {
				fmt.Println(products)
			}
		}
	}()
	appServer := NewAppServer()
	appServer.AddRoute("/", logMiddleware(timeoutMiddleware(IndexHandler)))
	appServer.AddRoute("/products", logMiddleware(ProductsHandler))
	appServer.AddRoute("/customers", logMiddleware(CustomersHandler))
	http.ListenAndServe(":8080", appServer)
}
