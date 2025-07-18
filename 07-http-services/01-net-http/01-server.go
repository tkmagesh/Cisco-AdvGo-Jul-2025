package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!\n"))
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
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
		}
	case "/customers":
		fmt.Fprintln(w, "List of the customers are on the way!")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
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
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
