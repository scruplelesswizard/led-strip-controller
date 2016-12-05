package api

import (
	"fmt"
	"log"
	"net/http"

	"goji.io"

	"github.com/chaosaffe/led-strip-controller/controller"

	"goji.io/pat"
	"golang.org/x/net/context"
)

var ss *controller.Strips

func Register(r *goji.Mux, s *controller.Strips) {
	ss = s

	r.HandleFuncC(pat.Get("/ping"), ping)

	r.HandleFuncC(pat.Get("/strips"), ping)
	r.HandleFuncC(pat.Get("/strip/:id"), ping)

	r.HandleFuncC(pat.Get("/strip/:id/effects"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/off"), off)
	r.HandleFuncC(pat.Post("/strip/:id/effect/fade"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/fadeBetween"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/fadeOut"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/flashBetween"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/flash"), ping)
	r.HandleFuncC(pat.Post("/strip/:id/effect/rotate"), rotate)
}

func ping(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!")
}

func off(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}
	s.Off()
}

func rotate(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}
	go s.Rotate()
}

func getStrip(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Strip id=%s", pat.Param(c, "id"))
}
