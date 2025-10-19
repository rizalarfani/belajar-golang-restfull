package middleware

import (
	"net/http"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/model/web"

	"github.com/gin-gonic/gin"
)

func ValidateApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if "RAHASIA" == c.Request.Header.Get("X-API-KEY") {
			c.Next()
		} else {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Api key tidak",
			}

			helper.WriteToResponseBody(c.Writer, webResponse)
			c.Abort()
			return
		}
	}
}
