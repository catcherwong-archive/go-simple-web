package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-simple-web/common"
	"go-simple-web/config"
	"go-simple-web/models"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type StudentDbService struct {
}

func NewStudentDbService() *StudentDbService {
	return &StudentDbService{}
}

func (ss *StudentDbService) GetAllStudents() (r models.CommonResp) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.PostgreSql.Host,
		config.Cfg.PostgreSql.Port,
		config.Cfg.PostgreSql.User,
		config.Cfg.PostgreSql.Pwd,
		config.Cfg.PostgreSql.DbName,
		"disable")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	sql := "select id, name, gender, create_time from t1"

	fmt.Println(db.Ping())

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var data []models.StudnetDto

	for rows.Next() {

		s := models.Student{}

		err = rows.Scan(&s.Id, &s.Name, &s.Gender, &s.CreateTime)
		if err != nil {
			panic(err)
		}

		data = append(data, s.ConvToDto())
	}

	r = models.GetOkResp(data)

	return r
}

func (ss *StudentDbService) GetStudentById(id int) (r models.CommonResp) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.PostgreSql.Host,
		config.Cfg.PostgreSql.Port,
		config.Cfg.PostgreSql.User,
		config.Cfg.PostgreSql.Pwd,
		config.Cfg.PostgreSql.DbName,
		"disable")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	sql := "select id, name, gender, create_time from t1 where id = $1 limit 1"

	s := models.Student{}
	err = db.QueryRow(sql, id).Scan(&s.Id, &s.Name, &s.Gender, &s.CreateTime)

	if err != nil {
		log.Println(err.Error())
		r = models.GetErrorResp(common.ERROR, "record is empty")
	} else {
		r = models.GetOkResp(s.ConvToDto())
	}

	return r
}

func (ss *StudentDbService) AddStudent(b []byte) (r models.CommonResp) {

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

	s.CreateTime = n.UnixNano() / int64(time.Millisecond)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.PostgreSql.Host,
		config.Cfg.PostgreSql.Port,
		config.Cfg.PostgreSql.User,
		config.Cfg.PostgreSql.Pwd,
		config.Cfg.PostgreSql.DbName,
		"disable")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())

		resp.Code = common.ERROR
		resp.Msg = "can not open database"
		resp.Data = false
		return resp
	}

	sql := " insert into t1(name, gender, create_time) values($1,$2,$3) "

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(s.Name, s.Gender, s.CreateTime)

	if err != nil {
		log.Println(err.Error())

		resp.Code = common.ERROR
		resp.Msg = "exec sql error"
		resp.Data = false
		return resp
	}

	return resp
}
