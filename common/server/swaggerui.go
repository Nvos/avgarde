package server

import (
	"avgarde/static"
	"io/fs"
	"net/http"
)

func SwaggerUI() http.Handler {
	// render the index template with the proper spec name inserted
	s, _ := fs.Sub(static.Swagger, "embed")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(s)))
	return mux
}