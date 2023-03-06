package controllers

type ResponseFailure struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
