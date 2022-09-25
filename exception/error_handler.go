package exception

import (
	"github.com/go-playground/validator/v10"
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		status := http.StatusBadRequest

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)

		webResponse := web.WebResponse{
			Code:   status,
			Status: http.StatusText(status),
			Data:   exception.Error(),
		}

		helper.WriteFromResponseBody(w, webResponse)
		return true
	} else {
		return false
	}

}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		status := http.StatusNotFound

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)

		webResponse := web.WebResponse{
			Code:   status,
			Status: http.StatusText(status),
			Data:   exception.Error,
		}

		helper.WriteFromResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(500),
		Data:   err,
	}

	helper.WriteFromResponseBody(w, webResponse)
}
