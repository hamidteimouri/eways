package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	e := echo.New()

	e.POST("/verify", handleVerify)
	l, err := net.Listen("tcp", "127.0.0.1:9009")
	if err != nil {
		return
	}

	err = http.Serve(l, e)
	if err != nil {
		panic("failed to serve http server")
	}

}
