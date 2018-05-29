package urlshort

import (
	"testing"
	"net/http/httptest"
)


func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	handler_func := HelloHandler()
	handler_func(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 code from HelloHandler, got %d instead", w.Code)
	}

	if w.Body.String() != "Hello, world!\n" {
		t.Errorf("Incorrect body for request to HelloHandler")
	}
}


func TestMapHandler(t *testing.T) {
	//pathsToUrls := map[string]string{
	//	"/foo": "http://bar.com",
	//}

	//handler := MapHandler(pathsToUrls, )
}
