package webhook

import (
	"github.com/OoXoSoO/labxIII/src/webhook/book"
	"github.com/OoXoSoO/labxIII/src/webhook/search"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Handle("/search",
		search.NewSearchHandle(
			&search.SearchService{},
		),
	)
	router.Handle("/book",
		book.NewBookHandle(
			&book.BookService{},
		),
	)

	log.Printf("Running!", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
