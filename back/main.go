package main

import (
	"net/http"

	"servermap/config"
	"servermap/internal/errors"
	"servermap/internal/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"
)

var env *config.Env

var serv *web.Service

func init() {
	env = config.NewEnv("./config")
}

func main() {
	erh := errors.New()

	serv = web.New(erh)
	serv.InitData(env.Config.HostsFile, env.Config.VmsFile)

	r := routes(serv)

	log.Debug("start server on port: ", env.Config.ListenPort)
	log.Fatal(http.ListenAndServe(":"+env.Config.ListenPort, r))
}

func routes(h *web.Service) *chi.Mux {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		cors.Handler,
	)
	//	router.Use(middleware.Logger)

	router.Get("/", h.Index)
	router.Get("/health-check", h.Alive)

	return router
}
