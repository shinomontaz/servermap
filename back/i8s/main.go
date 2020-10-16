package i8s

import "net/http"

const TypeWARNINIG string = "warning"
const TypeSUCCESS string = "success"
const TypeDANGER string = "danger"

type IErrorHandler interface {
	HandleError(w http.ResponseWriter, err error)
	Set401(w http.ResponseWriter, str string)
}
