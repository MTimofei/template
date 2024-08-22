package router

import "net/http"

type handler interface {
	Get(http.ResponseWriter, *http.Request)
}

func Router(h handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", http.HandlerFunc(h.Get))

	return mux
}
