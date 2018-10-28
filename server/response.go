package server

import (
	"github.com/kataras/iris"
)

func makeResponse(success bool, ctx iris.Context) map[string]interface{} {
	pass := "success"
	fail := "fail"

	var message string

	if success {
		message = pass
	} else {
		message = fail
	}

	return map[string]interface{}{
		"message": message,
	}
}
