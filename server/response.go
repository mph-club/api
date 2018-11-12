package server

type successResponse struct {
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type errorResponse struct {
	Error   map[string]interface{} `json:"error"`
	Message string                 `json:"message"`
}

func generateJSONResponse(success bool, statusCode int, responseMap map[string]interface{}) (int, interface{}) {
	pass := "success"
	fail := "fail"

	if success {
		return statusCode, &successResponse{
			Data:    responseMap,
			Message: pass,
		}
	}

	return statusCode, &errorResponse{
		Error:   responseMap,
		Message: fail,
	}

}
