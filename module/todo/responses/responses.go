package responses

import (
	"to_do_list/common/errors"
	"to_do_list/common/helpers"
)

// Response struktur
type Response struct {
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

// SuccessResponse struktur
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

// ErrResponse untuk error coy
type ErrResponse struct {
	Errors  []string `json:"errors"`
	Message string   `json:"message" example:"Bad request"`
	Status  string   `json:"status" example:"failed"`
}

//GenerateResponse pakai ini ya bro
func GenerateResponse(resp interface{}, err error) Response {
	response := Response{}
	if err != nil {
		httpErr := errors.CastToHTTPErrResp(err)
		response.Errors = helpers.ParseError(err)

		if httpErr.Status != 0 && httpErr.Body != nil {
			response.Message = httpErr.Body.(string)
		} else {
			response.Message = "request failed"
		}
	} else {
		response.Data = resp
		response.Message = "Request success"
		response.Status = "Success"
	}
	return response
}
