package services

import (
	"encoding/json"
	"go-simple-web/models"
	"log"
	"time"
)

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (ss *StudentService) GetAllStudents() (r models.CommonResp) {
	students := models.AllStudnet

	var data []models.StudnetDto

	for _, v := range students {
		data = append(data, v.ConvToDto())
	}

	r = models.GetOkResp(data)

	return r
}

func (ss *StudentService) GetStudentById(id int) (r models.CommonResp) {
	students := models.AllStudnet

	var s models.StudnetDto

	f := false

	for _, v := range students {
		if v.Id == id {
			s = v.ConvToDto()
			f = true
		}
	}

	if f {
		r = models.GetOkResp(s)
	} else {
		r = models.GetOkResp(nil)
	}

	return r
}

func (ss *StudentService) AddStudent(b []byte) (r models.CommonResp) {

	resp := models.GetOkResp(true)

	var s models.Student

	err := json.Unmarshal(b, &s)

	if err != nil {
		log.Printf("json.Unmarshal error, %s", err.Error())
		resp.Code = -1
		resp.Msg = "Can not parse json value"
		resp.Data = false
		return resp
	}

	if s.Name == "" {
		resp.Code = -1
		resp.Msg = "name must not empty"
		resp.Data = false
		return resp
	}

	if s.Gender > 2 || s.Gender < 1 {
		resp.Code = -1
		resp.Msg = "gender must be 1 or 2"
		resp.Data = false
		return resp
	}

	n := time.Now()

	s.Id = n.Hour() + n.Minute() + n.Second()
	s.CreateTime = n.UnixNano() / int64(time.Millisecond)

	models.AllStudnet = append(models.AllStudnet, s)

	return resp
}
