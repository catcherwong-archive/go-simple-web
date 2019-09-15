package v1

import (
	"go-simple-web/handlers"
	"net/http"
)

func InitRoute() {
	initStudnetRoute()
	initConfigRoute()
	initStudnetDbRoute()
}

func initStudnetRoute() {
	sh := handlers.NewStudentHandler()
	http.HandleFunc("/v1/student", sh.V1Handler)
}

func initConfigRoute() {
	ch := handlers.NewConfigHandler()
	http.HandleFunc("/v1/config", ch.V1Handler)
}

func initStudnetDbRoute() {
	sh := handlers.NewStudentDbHandler()
	http.HandleFunc("/v1/student/db", sh.V1Handler)
}
