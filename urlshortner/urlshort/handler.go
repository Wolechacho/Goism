package urlshort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

//URLPath contains the path and url
type URLPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"path"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, ok := pathsToUrls[r.URL.String()]
		if ok {
			http.Redirect(w, r, val, 301)

		} else {
			fallback.ServeHTTP(w, r)
		}
	})
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
	// TODO: Implement this...
	urlPaths, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	urlsMaps := buildMap(urlPaths)
	return MapHandler(urlsMaps, fallback), nil
}
func parseYAML(y []byte) ([]URLPath, error) {
	m := []URLPath{}
	err := yaml.Unmarshal(y, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return m, nil
}
func buildMap(urls []URLPath) map[string]string {
	m := make(map[string]string)
	for _, urlPath := range urls {
		m[urlPath.Path] = urlPath.URL
	}
	return m
}
