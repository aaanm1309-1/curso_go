package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
	// })
	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler)
	mux.HandleFunc("GET /bookslist/{$}", BooksHandler) // exatidão do caminho com / ou nao no final

	mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)
	mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)
	// mux.HandleFunc("GET /booksid/{x}", BooksPrecedence3Handler)
	// mux.HandleFunc("GET /{x}/latest", BooksPrecedence4Handler)
	http.ListenAndServe(":9000", mux)
}

// Examples: http://localhost:9000/books/2
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Write([]byte("Book: " + id + "\n"))
	// fmt.Fprintf(w, "Book ID: %s", id)

}

// Examples: http://localhost:9000/books/dir/sinqia/pod/server
func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirPath := r.PathValue("d")

	w.Write([]byte("Directory: " + dirPath + "\n"))
	// fmt.Fprintf(w, "Book ID: %s", id)

}

// Examples: http://localhost:9000/bookslist/
func BooksHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("BooksList"))

}

// Examples: http://localhost:9000/books/precedence/latest
func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Books Precedence"))

}

// Examples: http://localhost:9000/books/precedence/xxxx
func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Books Precedence 2"))

}

// Examples: http://localhost:9000/booksid/23
func BooksPrecedence3Handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Books Precedence 3"))

}

// Examples: http://localhost:9000/books/precedence/xxxx
func BooksPrecedence4Handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Books Precedence 4"))

}
