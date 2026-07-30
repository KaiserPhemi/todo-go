package todo

import "net/http"

func RegisterRoutes(
	mux *http.ServeMux,
	handler *Handler,
) {

	mux.HandleFunc(
		"GET /todos/{id}",
		handler.store.Get()
	)
}