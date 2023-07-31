package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
)

func main() {
	e := echo.New()
	vld := validator.New()
	handler := handlers{vld}
	e.POST("/verify", handler.verify)
	l, err := net.Listen("tcp", "127.0.0.1:9009")
	if err != nil {
		return
	}

	err = http.Serve(l, e)
	if err != nil {
		panic("failed to serve http server")
	}

}
