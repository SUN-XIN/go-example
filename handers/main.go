package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerSimple(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello %s", name)
}

/////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////    myHandler implements http.Handler    /////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
type myHandler func(http.ResponseWriter, *http.Request) (int, error)

func (fn myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := fn(w, r); err != nil {
		switch status {
		case http.StatusBadRequest:
			http.Error(w, fmt.Sprintf("Please check your request: %v", err), http.StatusBadRequest)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func handlerImpl(w http.ResponseWriter, r *http.Request) (int, error) {
	name := r.FormValue("name")
	if name == "" {
		return http.StatusBadRequest, fmt.Errorf("Name not found from URL")
	}

	fmt.Fprintf(w, "Hello %s", name)
	return http.StatusOK, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////           handler with variable       /////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////
type Setting struct {
	Date   time.Time
	Format string
}

func handlerWithVar(s Setting) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server starts at: %s", s.Date.Format(s.Format))
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	http.HandleFunc("/simple", handlerSimple)
	http.Handle("/impl", myHandler(handlerImpl))

	s := Setting{
		Date:   time.Now(),
		Format: "2006-01-02 15:04",
	}
	http.Handle("/with_var", handlerWithVar(s))
	err := http.ListenAndServe(":8080", nil)

	log.Printf("PROG FAILED: %+v", err)
}
