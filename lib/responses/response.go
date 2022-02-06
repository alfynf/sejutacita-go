package responses

import "net/http"

// fungsi untuk memberikan respon success API
func SuccessResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":   http.StatusOK,
		"status": "Successful Operation",
	}
	return result
}

// fungsi untuk memberikan respon success API dengan mereturn data
func SuccessResponseWithData(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":   http.StatusOK,
		"status": "Successful Operation",
		"data":   data,
	}
	return result
}

func FailedReponse(message string) map[string]interface{} {
	result := map[string]interface{}{
		"code":          http.StatusBadRequest,
		"status":        "Operation Failed",
		"error message": message,
	}
	return result
}
