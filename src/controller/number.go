package controller

import (
	"encoding/json"
	"fmt"
	hc "go-rest-api/src/constant"
	"go-rest-api/src/handler"
	"net/http"
	"strconv"
)

type Data struct {
	Number string `json:"number"`
}

type NumberPostRequest struct {
	Number string `json:"number"`
	Date   Time   `json:"date"`
}

type Time struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetNumber(w http.ResponseWriter, r *http.Request) {
	//get the request body param

	number := r.URL.Query().Get("number")
	if number == "" {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST)
		return
	}
	//check if given number is string
	if _, err := strconv.ParseInt(number, 10, 64); err != nil {
		handler.HttpError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		res := Data{
			Number: number,
		}
		handler.HttpResponse(w, http.StatusOK, res)
	}
}

/**
 * [Body description]
 * @type {[type]}
 */
func PostNumber(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST)
		return
	}
	var req NumberPostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST)
		return
	}
	if req.Number == "" || req.Date.CreatedAt == "" || req.Date.UpdatedAt == "" {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST)
		return
	}
	res := Data{
		Number: req.Number,
	}

	fmt.Println(r.Header.Get("x-api-key"))
	handler.HttpResponse(w, http.StatusOK, res)
}
