package exception

import (
	"net/http"
	"rizalarfani/belajar-restful-api/helper"
	"rizalarfani/belajar-restful-api/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// coba map ke 400
				if validationError(c, r) {
					c.Abort()
					return
				}
				// coba map ke 404
				if notFoundError(c, r) {
					c.Abort()
					return
				}
				// default → 500
				internalServerError(c, r)
				c.Abort()
				return
			}
		}()

		c.Next()

		// Jika sudah ada response yang dikirim oleh handler sebelumnya, selesai.
		if c.IsAborted() || len(c.Errors) == 0 {
			return
		}

	}
}

func validationError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(c.Writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(c.Writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, err interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}
