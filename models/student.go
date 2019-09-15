package models

import (
	"go-simple-web/common"
	"time"
)

var AllStudnet []Student = []Student{Student{Id: 1, Name: "catcher", Gender: 1, CreateTime: 1568507970744}, Student{Id: 2, Name: "lisa", Gender: 2, CreateTime: 1568507970745}}

type Student struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Gender     int    `json:"gender"`
	CreateTime int64  `json:"create_time"`
}

type StudnetDto struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	CreateTime string `json:"create_time"`
}

func (s Student) ConvToDto() (d StudnetDto) {

	d.Id = s.Id
	d.Name = s.Name

	if s.Gender == 1 {
		d.Gender = "male"
	} else {
		d.Gender = "female"
	}

	d.CreateTime = time.Unix(0, s.CreateTime*int64(time.Millisecond)).Format(common.TimeMilFormat)

	return d
}
