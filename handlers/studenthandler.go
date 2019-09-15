package handlers

import (
	"encoding/json"
	"fmt"
	"go-simple-web/models"
	"go-simple-web/services"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler() *StudentHandler {
	return &StudentHandler{service: new(services.StudentService)}
}

func (sh *StudentHandler) V1Handler(w http.ResponseWriter, r *http.Request) {

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
		get(w, r, sh)
	} else if m == "POST" {
		post(w, r, sh)
	} else if m == "PUT" {
		put(w, r, sh)
	} else if m == "DELETE" {
		del(w, r, sh)
	} else {
		other(w, r, sh)
	}

}

func other(w http.ResponseWriter, r *http.Request, sh *StudentHandler) {

	resp := models.GetOkResp(true)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func del(w http.ResponseWriter, r *http.Request, sh *StudentHandler) {

	resp := models.GetOkResp(true)

	b, err := json.Marshal(resp)

	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(b))
}

func put(w http.ResponseWriter, r *http.Request, sh *StudentHandler) {

	resp := models.GetOkResp(true)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func post(w http.ResponseWriter, r *http.Request, sh *StudentHandler) {

	body, _ := ioutil.ReadAll(r.Body)

	resp := sh.service.AddStudent(body)

	b, _ := json.Marshal(resp)

	fmt.Fprint(w, string(b))
}

func get(w http.ResponseWriter, r *http.Request, sh *StudentHandler) {

	resp := models.GetErrorResp(-1, "ERROR")

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
