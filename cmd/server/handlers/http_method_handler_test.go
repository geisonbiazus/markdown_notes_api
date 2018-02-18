package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPMethodHandler(t *testing.T) {
	setup := func() (
		*HandlerSpy,
		*HandlerSpy,
		*HandlerSpy,
		*HandlerSpy,
		*HTTPMethodHandler,
		*httptest.ResponseRecorder,
	) {
		getHandler := NewHandlerSpy()
		postHandler := NewHandlerSpy()
		putHandler := NewHandlerSpy()
		deleteHandler := NewHandlerSpy()
		httpMethodHandler := &HTTPMethodHandler{
			Get:    getHandler,
			Post:   postHandler,
			Put:    putHandler,
			Delete: deleteHandler,
		}

		w := httptest.NewRecorder()

		return getHandler, postHandler, putHandler, deleteHandler, httpMethodHandler, w
	}

	t.Run("Given a GET request, it executes the defined Handler", func(t *testing.T) {
		getHandler, postHandler, putHandler, deleteHandler, httpMethodHandler, w := setup()

		r := httptest.NewRequest(http.MethodGet, "http://example.com", nil)

		httpMethodHandler.ServeHTTP(w, r)

		shouldCallHandler(t, getHandler, w, r)
		shouldNotCallHandler(t, postHandler)
		shouldNotCallHandler(t, putHandler)
		shouldNotCallHandler(t, deleteHandler)
	})

	t.Run("Given a POST request, it executes the defined Handler", func(t *testing.T) {
		getHandler, postHandler, putHandler, deleteHandler, httpMethodHandler, w := setup()

		r := httptest.NewRequest(http.MethodPost, "http://example.com", nil)

		httpMethodHandler.ServeHTTP(w, r)

		shouldNotCallHandler(t, getHandler)
		shouldCallHandler(t, postHandler, w, r)
		shouldNotCallHandler(t, putHandler)
		shouldNotCallHandler(t, deleteHandler)
	})

	t.Run("Given a PUT request, it executes the defined Handler", func(t *testing.T) {
		getHandler, postHandler, putHandler, deleteHandler, httpMethodHandler, w := setup()

		r := httptest.NewRequest(http.MethodPut, "http://example.com", nil)

		httpMethodHandler.ServeHTTP(w, r)

		shouldNotCallHandler(t, getHandler)
		shouldNotCallHandler(t, postHandler)
		shouldCallHandler(t, putHandler, w, r)
		shouldNotCallHandler(t, deleteHandler)
	})

	t.Run("Given a DELETE request, it executes the defined Handler", func(t *testing.T) {
		getHandler, postHandler, putHandler, deleteHandler, httpMethodHandler, w := setup()

		r := httptest.NewRequest(http.MethodDelete, "http://example.com", nil)

		httpMethodHandler.ServeHTTP(w, r)

		shouldNotCallHandler(t, getHandler)
		shouldNotCallHandler(t, postHandler)
		shouldNotCallHandler(t, putHandler)
		shouldCallHandler(t, deleteHandler, w, r)
	})

	t.Run("Given a method that doesn't have a Handler, it returns not found", func(t *testing.T) {
		_, _, _, _, httpMethodHandler, w := setup()

		r := httptest.NewRequest(http.MethodGet, "http://example.com", nil)

		httpMethodHandler.Get = nil

		httpMethodHandler.ServeHTTP(w, r)

		if w.Result().StatusCode != http.StatusNotFound {
			t.Errorf("It didn't return NotFound. Expected: %v. Actual: %v", http.StatusNotFound, w.Result().StatusCode)
		}
	})

	t.Run("Given request with a not supported method, it returns not found", func(t *testing.T) {
		_, _, _, _, httpMethodHandler, w := setup()

		r := httptest.NewRequest("unsupported", "http://example.com", nil)

		httpMethodHandler.ServeHTTP(w, r)

		if w.Result().StatusCode != http.StatusNotFound {
			t.Errorf("It didn't return NotFound. Expected: %v. Actual: %v", http.StatusNotFound, w.Result().StatusCode)
		}
	})
}

func shouldCallHandler(t *testing.T, h *HandlerSpy, w http.ResponseWriter, r *http.Request) {
	if !h.ServeHTTPCalled {
		t.Error("It didn't call the Hanlder")
	}

	if h.ServeHTTPResponseWritterArg != w {
		t.Errorf("Expected: %v. Actual: %v", w, h.ServeHTTPResponseWritterArg)
	}

	if h.ServeHTTPRequestArg != r {
		t.Errorf("Expected: %v. Actual: %v", r, h.ServeHTTPRequestArg)
	}
}

func shouldNotCallHandler(t *testing.T, h *HandlerSpy) {
	if h.ServeHTTPCalled {
		t.Error("It called the Handler")
	}
}

type HandlerSpy struct {
	ServeHTTPCalled             bool
	ServeHTTPResponseWritterArg http.ResponseWriter
	ServeHTTPRequestArg         *http.Request
}

func (s *HandlerSpy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTPCalled = true
	s.ServeHTTPResponseWritterArg = w
	s.ServeHTTPRequestArg = r
}

func NewHandlerSpy() *HandlerSpy {
	return &HandlerSpy{}
}
