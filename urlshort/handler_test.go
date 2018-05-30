package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandler(t *testing.T) {
	pathsToUrls := map[string]string{
		"/foo": "http://bar.com",
	}

	handler_func := MapHandler(pathsToUrls, http.NotFoundHandler())

	req := httptest.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()
	handler_func(w, req)

	if w.Code != 302 {
		t.Errorf("Expected 302 code from MapHandler, got %d instead", w.Code)
	}

	req = httptest.NewRequest("GET", "/not_a_url", nil)
	w = httptest.NewRecorder()
	handler_func(w, req)

	if w.Code != 404 {
		t.Errorf("Expected 404 code from MapHandler, got %d instead", w.Code)
	}
}

func TestYAMLHandler(t *testing.T) {
	yaml := `
- path: /baz
  url: http://foo.co.uk
`
	handler_func, err := YAMLHandler([]byte(yaml), http.NotFoundHandler())
	if err != nil {
		t.Errorf("Error when trying to parse valid YAML into a handler: %v", err)
	}

	req := httptest.NewRequest("GET", "/baz", nil)
	w := httptest.NewRecorder()
	handler_func(w, req)

	if w.Code != 302 {
		t.Errorf("Expected 302 code from YAMLHandler, got %d instead", w.Code)
	}
}

func TestYAMLHandlerInvalidYAML(t *testing.T) {
	yaml := "I am not YAML"

	_, err := YAMLHandler([]byte(yaml), http.NotFoundHandler())
	if err == nil {
		t.Errorf("expected Error when parsing invalid YAML")
	}
}
