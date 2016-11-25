package api

import (
	"goji.io"
	"goji.io/pat"
	"net/http"
	"fmt"
	"golang.org/x/net/context"
)

func Register(r *goji.Mux) {
	r.HandleFuncC(pat.Get("/ping"), ping)

	r.HandleFuncC(pat.Get("/strips"), ping)
	r.HandleFuncC(pat.Get("/strip/:id"), ping)

	r.HandleFuncC(pat.Get("/strip/:id/effects"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/off"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/fade"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/fadeBetween"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/fadeOut"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/flashBetween"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/flash"), ping)
}

func ping(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!")
}

func getStrip(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Strip id=%s", pat.Param(c, "id"))
}