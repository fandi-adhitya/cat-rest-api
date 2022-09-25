package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		status := http.StatusUnauthorized

		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(status)

		webResponse := web.WebResponse{
			Code:   status,
			Status: http.StatusText(status),
		}

		helper.WriteFromResponseBody(writer, webResponse)
	}
}
