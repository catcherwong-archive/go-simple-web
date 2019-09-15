package handlers

import (
	"encoding/json"
	"fmt"
	"go-simple-web/config"
	"go-simple-web/models"
	"log"
	"net/http"
)

type ConfigHandler struct {
}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{}
}

func (ch *ConfigHandler) V1Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ct := r.Header.Get("Content-Type")

	log.Printf("Content-Type is [%s]", ct)

	if ct != "application/json" {

		resp := models.GetErrorResp(-1, "Content-Type must be application/json")

		b, err := json.Marshal(resp)

		if err != nil {
			log.Println(err)
		}

		fmt.Fprint(w, string(b))
		return
	}

	m := r.Method

	log.Printf("method=%s, url=%s, remoteaddr=%s", r.Method, r.URL, r.RemoteAddr)

	if m == "GET" {
		ch.get(w, r)
	} else {
		ch.other(w, r)
	}

}

func (ch *ConfigHandler) other(w http.ResponseWriter, r *http.Request) {

	resp := models.GetOkResp(true)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func (ch *ConfigHandler) get(w http.ResponseWriter, r *http.Request) {

	resp := models.GetOkResp(config.Cfg)

	b, err := json.Marshal(resp)

	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(b))
}
