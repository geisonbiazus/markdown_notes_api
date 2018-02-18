package handlers

import "net/http"

type HTTPMethodHandler struct {
	Get    http.Handler
	Post   http.Handler
	Put    http.Handler
	Delete http.Handler
}

func (h *HTTPMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.getHandler(r.Method).ServeHTTP(w, r)
}

func (h *HTTPMethodHandler) getHandler(method string) http.Handler {
	handler := h.getHandlerBasedOnMethod(method)
	if handler != nil {
		return handler
	}
	return http.NotFoundHandler()
}

func (h *HTTPMethodHandler) getHandlerBasedOnMethod(method string) http.Handler {
	switch method {
	case http.MethodGet:
		return h.Get
	case http.MethodPost:
		return h.Post
	case http.MethodPut:
		return h.Put
	case http.MethodDelete:
		return h.Delete
	}
	return nil
}
