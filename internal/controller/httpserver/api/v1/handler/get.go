package handler

import "net/http"

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get"))
}
