package server

func generateJSONResponse(success bool, responseMap map[string]interface{}) map[string]interface{} {
	pass := "success"
	fail := "fail"

	var message string
	mapToIterate := map[string]interface{}{}

	if success {
		message = pass
		mapToIterate["data"] = responseMap
	} else {
		message = fail
		mapToIterate["error"] = responseMap
	}

	mapToIterate["message"] = message

	return mapToIterate
}
