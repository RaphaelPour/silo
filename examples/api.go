package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RaphaelPour/silo"
)

// $ go run api.go
// $ curl -X POST -d "value" http://localhost:8000/key
// $ curl http://localhost:8000/key
// $ value

func main() {
	store := silo.NewFile("blob.storage")
	api := http.NewServeMux()
	api.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		key := req.URL.Path[1:]

		if req.Method == http.MethodGet {
			value, err := store.Get(key)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, err.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", value)
		} else if req.Method == http.MethodPost {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, err.Error())
				return
			}
			if err := store.Set(key, body); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, err.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	})

	http.ListenAndServe(":8000", api)
}
