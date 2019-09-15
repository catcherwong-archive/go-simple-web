package handlers

import (
	"encoding/json"
	"fmt"
	"go-simple-web/common"
	"go-simple-web/models"
	"go-simple-web/services"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type StudentDbHandler struct {
	service *services.StudentDbService
}

func NewStudentDbHandler() *StudentDbHandler {
	return &StudentDbHandler{service: new(services.StudentDbService)}
}

func (sh *StudentDbHandler) V1Handler(w http.ResponseWriter, r *http.Request) {

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
		sh.get(w, r)
	} else if m == "POST" {
		sh.post(w, r)
	} else if m == "PUT" {
		sh.put(w, r)
	} else if m == "DELETE" {
		sh.del(w, r)
	} else {
		sh.other(w, r)
	}

}

func (sh *StudentDbHandler) other(w http.ResponseWriter, r *http.Request) {

	resp := models.GetOkResp(true)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func (sh *StudentDbHandler) del(w http.ResponseWriter, r *http.Request) {

	resp := models.GetOkResp(true)

	b, err := json.Marshal(resp)

	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(b))
}

func (sh *StudentDbHandler) put(w http.ResponseWriter, r *http.Request) {

	resp := models.GetOkResp(true)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func (sh *StudentDbHandler) post(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	resp := sh.service.AddStudent(body)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func (sh *StudentDbHandler) get(w http.ResponseWriter, r *http.Request) {

	resp := models.GetErrorResp(common.ERROR, "ERROR")

	u, _ := url.Parse(r.URL.String())
	q, _ := url.ParseQuery(u.RawQuery)

	id, ok := q["id"]

	if ok {
		i, err := strconv.Atoi(id[0])

		if err != nil {
			log.Println(err)
		}

		resp = sh.service.GetStudentById(i)
	} else {
		resp = sh.service.GetAllStudents()
	}

	b, err := json.Marshal(resp)

	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(b))
}
