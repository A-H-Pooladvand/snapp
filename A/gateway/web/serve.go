package web

import (
	"fmt"
	"gateway/configs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Serve() {
	app := configs.NewApp()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	Routes(r)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", app.Port),
		r,
	)

	if err != nil {
		panic(err)
	}
}
