package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		redirectUrl, ok := pathsToUrls[request.URL.Path]
		if ok {
			println("Redirecting from {v} to {v}", request.URL, redirectUrl)
			http.Redirect(writer, request, redirectUrl, http.StatusFound)
		} else {
			println("Not found")
			fallback.ServeHTTP(writer, request)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	redirects, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildMap(redirects)
	return MapHandler(pathsToUrls, fallback), nil
}

type redirect struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYAML(yml []byte) (redirects []redirect, err error) {
	err = yaml.Unmarshal(yml, &redirects)
	if err != nil {
		return nil, err
	}
	return redirects, nil
}

func buildMap(redirects []redirect) map[string]string {
	output := make(map[string]string, len(redirects))
	for _, redirect := range redirects {
		output[redirect.Path] = redirect.Url
	}
	return output
}
