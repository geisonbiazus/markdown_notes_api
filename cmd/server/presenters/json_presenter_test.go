package presenters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

func TestJSONPresenter(t *testing.T) {
	setup := func() (
		*httptest.ResponseRecorder,
		*JSONPresenter,
	) {
		w := httptest.NewRecorder()
		presenter := &JSONPresenter{w}

		return w, presenter
	}

	type toSerialize struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	t.Run("RenderJSON", func(t *testing.T) {
		t.Run("Given a struct and a status code, it renders the data as a JSON", func(t *testing.T) {
			w, presenter := setup()

			s := toSerialize{
				ID:    1,
				Title: "Title",
			}

			presenter.RenderJSON(http.StatusOK, s)

			expectedBody := []byte(fmt.Sprintf(`{"id":%d,"title":"%s"}`, s.ID, s.Title))

			assertResponse(t, w, expectedBody, http.StatusOK)
		})
	})

	t.Run("PresentError", func(t *testing.T) {
		t.Run("Given an error list, it renders the erros as JSON", func(t *testing.T) {
			w, presenter := setup()

			errs := []markdownnotes.ValidationError{
				markdownnotes.ValidationError{Field: "firstField", Message: "Error message", Type: "ERROR_TYPE"},
				markdownnotes.ValidationError{Field: "secondField", Message: "Error message", Type: "ERROR_TYPE"},
			}

			presenter.PresentValidationErrors(errs)

			expectedBody := []byte(
				`{"errors":[` +
					fmt.Sprintf(`{"field":"%s","type":"%s","message":"%s"},`, errs[0].Field, errs[0].Type, errs[0].Message) +
					fmt.Sprintf(`{"field":"%s","type":"%s","message":"%s"}`, errs[1].Field, errs[1].Type, errs[1].Message) +
					`]}`)

			assertResponse(t, w, expectedBody, http.StatusUnprocessableEntity)
		})
	})

	t.Run("ServiceUnavailable", func(t *testing.T) {
		t.Run("It renders a JSON saying that the service is unavailable", func(t *testing.T) {
			w, presenter := setup()

			presenter.PresentError(nil)

			expectedBody := []byte(
				`{"errors":[` +
					`{"field":"","type":"SERVICE_UNAVAILABLE","message":"Service Unavailable"}` +
					`]}`)

			assertResponse(t, w, expectedBody, http.StatusServiceUnavailable)
		})
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Run("It renders NotFound", func(t *testing.T) {
			w, presenter := setup()
			presenter.PresentNotFound()

			if w.Code != http.StatusNotFound {
				t.Errorf("Expected: %v. Actual: %v", http.StatusNotFound, w.Code)
			}
		})
	})
}

func responseBody(w *httptest.ResponseRecorder) []byte {
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func assertResponse(t *testing.T, w *httptest.ResponseRecorder, expectedBody []byte, expectedStatus int) {
	body := responseBody(w)

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
	}

	if w.Result().StatusCode != expectedStatus {
		t.Errorf("Expected: %d. Actual: %d", expectedStatus, w.Result().StatusCode)
	}

	expectedContentType := "application/json"
	contentType := w.Result().Header.Get("Content-Type")

	if contentType != expectedContentType {
		t.Errorf("Expected: %s. Actual: %s", expectedContentType, contentType)
	}
}
