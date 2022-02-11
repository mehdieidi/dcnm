package frontend

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/MehdiEidi/dcnm/core"
	"github.com/gorilla/mux"
)

type restFrontEnd struct {
	store *core.KeyValueStore
}

func (f *restFrontEnd) Start(store *core.KeyValueStore) error {
	f.store = store

	r := mux.NewRouter()

	r.Use(f.loggingMiddleware)

	r.HandleFunc("/v1/{key}", f.keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", f.keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", f.keyValueDeleteHandler).Methods("DELETE")

	r.HandleFunc("/v1", f.notAllowedHandler)
	r.HandleFunc("/v1/{key}", f.notAllowedHandler)

	return http.ListenAndServe(":8080", r)
}

func (f *restFrontEnd) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (f *restFrontEnd) notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

func (f *restFrontEnd) keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.store.Put(key, string(value), false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	log.Printf("PUT key=%s value=%s\n", key, string(value))
}

func (f *restFrontEnd) keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := f.store.Get(key)
	if errors.Is(err, core.ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))

	log.Printf("GET key=%s\n", key)
}

func (f *restFrontEnd) keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := f.store.Delete(key, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("DELETE key=%s\n", key)
}
