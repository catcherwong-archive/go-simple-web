package models

import "go-simple-web/common"

type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetOkResp(d interface{}) (r CommonResp) {
	r.Code = common.SUCCESS
	r.Msg = "OK"
	r.Data = d

	return r
}

func GetErrorResp(c int, m string) (r CommonResp) {
	r.Code = c
	r.Msg = m
	r.Data = nil

	return r
}
