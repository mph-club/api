package server

func generateJSONResponse(success bool, responseMap map[string]interface{}) map[string]interface{} {
	pass := "success"
	fail := "fail"

	var message string
	mapToIterate := map[string]interface{}{}

	if success {
		message = pass
	} else {
		message = fail
	}

	for k, v := range responseMap {
		mapToIterate[k] = v
	}

	mapToIterate["message"] = message

	return mapToIterate
}
