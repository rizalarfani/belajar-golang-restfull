package middleware

import (
	"net/http"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handle http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handle,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Api key tidak",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
