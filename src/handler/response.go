package handler

import (
	"encoding/json"
	"net/http"
)

//Response api object
type Response struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Error api response object
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

/**
 * [err description]
 * @type {[type]}
 */
func NewHttpResponse(w http.ResponseWriter, status int, success bool, payload interface{}, httpError interface{}) {

	response, err := json.Marshal(

		&Response{
			Success: success,
			Error:   httpError,
			Data:    payload,
		})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

/**
 * [err description]
 * @type {[type]}
 */
func HttpError(w http.ResponseWriter, status int, message string) {
	err := Error{
		Code:    status,
		Message: message,
	}
	NewHttpResponse(w, status, false, nil, err)
}

/**
 *
 */
func HttpResponse(w http.ResponseWriter, status int, payload interface{}) {

	NewHttpResponse(w, status, true, payload, nil)
}
