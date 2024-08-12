package main

import (
	"fmt"
	"log"
	"net/http"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entrou no middleware")
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v\n", r)
				// debug.PrintStack()
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}

		}()
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic!")
	})

	log.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen server on :3000 %v\n", err)
	}

}
