package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"goji.io"

	"github.com/chaosaffe/led-strip-controller/controller"

	"goji.io/pat"
	"golang.org/x/net/context"
)

var ss *controller.Strips

type SingleColorRequest struct {
	HSI             controller.HSI `json:"hsi"`
	DurationSeconds int            `json:"durationSeconds"`
}

type MultiColorRequest struct {
	HSI             []controller.HSI `json:"hsi"`
	DurationSeconds int              `json:"durationSeconds"`
}

type DurationOnlyRequest struct {
	DurationSeconds int `json:"durationSeconds"`
}

func Register(r *goji.Mux, s *controller.Strips) {

	ss = s

	r.HandleFuncC(pat.Get("/strips"), listStrips)
	r.HandleFuncC(pat.Get("/strips/:id"), getStrip)
	r.HandleFuncC(pat.Get("/strips/:id/effects"), listEffects)
	r.HandleFuncC(pat.Post("/strips/:id/effect/off"), off)
	r.HandleFuncC(pat.Post("/strips/:id/effect/fade"), fade)
	r.HandleFuncC(pat.Post("/strips/:id/effect/fade-between"), fadeBetween)
	r.HandleFuncC(pat.Post("/strips/:id/effect/fade-out"), fadeOut)
	r.HandleFuncC(pat.Post("/strips/:id/effect/flash-between"), flashBetween)
	r.HandleFuncC(pat.Post("/strips/:id/effect/flash"), flash)
	r.HandleFuncC(pat.Post("/strips/:id/effect/rotate"), rotate)
	r.HandleFuncC(pat.Post("/strips/:id/effect/pulse"), pulse)
}

func listStrips(c context.Context, w http.ResponseWriter, r *http.Request) {

	sJSON, err := json.Marshal(ss)
	if err != nil {
		r.Response.StatusCode = 500
		return
	}

	fmt.Fprint(w, string(sJSON))

}

func getStrip(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	sJSON, err := json.Marshal(s)
	if err != nil {
		r.Response.StatusCode = 500
		return
	}

	fmt.Fprint(w, string(sJSON))

}

func listEffects(c context.Context, w http.ResponseWriter, r *http.Request) {
	// TODO: List avaliable effects
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

	var req DurationOnlyRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.Rotate(d)
}

func fade(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	var req SingleColorRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.Fade(req.HSI, d)
}

func fadeOut(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	var req DurationOnlyRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.FadeOut(d)
}

func flash(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	var req SingleColorRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.Flash(req.HSI, d)
}

func pulse(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	var req SingleColorRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.Pulse(req.HSI, d)
}

func fadeBetween(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	var req MultiColorRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.FadeBetween(req.HSI, d)
}

func flashBetween(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(c, "id")
	s, err := ss.GetStrip(name)
	if err != nil {
		log.Printf("Could not use strip '%s': %s", name, err)
		return
	}

	var req MultiColorRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err = decoder.Decode(&req)
	if err != nil {
		panic(err)
	}

	d := time.Duration(req.DurationSeconds) * time.Second

	go s.FlashBetween(req.HSI, d)
}
