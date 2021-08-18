package server

import (
	"avgarde/api/openapi"
	"avgarde/static"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"io/fs"
	"net/http"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddr(":8080", createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	s, _ := fs.Sub(static.Swagger, "swagger-ui")
	fileServer := http.FileServer(http.FS(s))

	rootRouter.Handle("/swagger/spec/account", openapi.AccountHandler())
	rootRouter.Handle("/swagger/spec/item", openapi.ItemHandler())
	rootRouter.Handle("/swagger/*", http.StripPrefix("/swagger", fileServer))
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	chi.Walk(rootRouter, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		println(method, route)
		return nil
	})

	http.ListenAndServe(addr, rootRouter)
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug: true,
	})
	router.Use(corsMiddleware.Handler)
}
