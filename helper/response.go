package helper

import "strings"

type APIResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func FormatResponse(status string, code int, message string, data interface{}) APIResponse {
	metaResponse := Meta{
		Status:  status,
		Code:    code,
		Message: message,
	}

	apiResponse := APIResponse{
		Meta: metaResponse,
		Data: data,
	}

	return apiResponse
}

func FormatError(err error) []string {
	errMsg := strings.Split(err.Error(), "\n")

	return errMsg
}
