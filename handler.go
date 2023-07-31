package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	httpClientGoLib "github.com/trustwallet/go-libs/client"
	"net/http"
	"time"
)

const (
	KycCheckerBaseUrl     = "https://coding-challenge.xeptore.me"
	KycCheckerSecretKey   = "mZjMs7-Ci3wqXaFtI5FdhEqAb8Z8YkeYOOmmorinEHVf0bZHn_DCnM7oItT"
	KycCheckerOkResponse  = 360
	KycCheckerNOkResponse = 431
)

type handlers struct {
	validator *validator.Validate
}

func (h *handlers) verify(c echo.Context) error {
	r := VerifyRequest{}
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	err = h.validateData(r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	f := &FullName{FullName: fmt.Sprintf("%s %s", r.Name, r.Family)}
	err = h.validator.Struct(f)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "maximum of character is 20")
	}

	result, code := checkKyc(c.Request().Context(), f)
	return c.JSON(code, result)

}

type VerifyRequest struct {
	Name   string `json:"firstName" form:"firstName" validate:"required"`
	Family string `json:"lastName" form:"lastName" validate:"required"`
}

type FullName struct {
	FullName string `json:"name" validate:"max=20"`
}

func (h *handlers) validateData(r VerifyRequest) error {
	err := h.validator.Struct(&r)
	if err != nil {
		return err
	}
	return nil
}

func checkKyc(ctx context.Context, data *FullName) (interface{}, int) {

	req := httpClientGoLib.InitJSONClient(KycCheckerBaseUrl, nil)

	nb := httpClientGoLib.NewReqBuilder()
	nb.Headers(map[string]string{
		"Secret": KycCheckerSecretKey,
	}).Body(data).Method("POST").PathStatic("/verify")

	body, err := req.Execute(ctx, nb.Build())
	if err != nil {
		resp := RespNotOk{
			Ok: false,
		}
		return resp, http.StatusInternalServerError
	}

	var response1 KycCheckerResponse1
	err = json.Unmarshal(body, &response1)
	if err == nil {
		resp := RespError{
			Error: "ABC",
		}
		return resp, http.StatusInternalServerError
	}

	var response2 []interface{}
	err = json.Unmarshal(body, &response2)
	if err == nil {

		if len(response2) == 4 && response2[3] == float64(KycCheckerNOkResponse) {
			resp := RespError{
				Error: "You are not allowed",
			}
			return resp, http.StatusUnauthorized
		}
		if len(response2) == 6 && response2[3] == float64(KycCheckerOkResponse) {
			t := int64(0)
			if val, ok := response2[4].(int64); ok {
				// The interface value is an int64, convert it
				t = val
			}
			resp := RespOk{
				Message: fmt.Sprintf("Request ID %s is usable since %s", response2[4], time.UnixMilli(t).Format(time.RFC3339)),
			}
			return resp, http.StatusOK
		}

		resp := RespError{
			Error: "ABC",
		}
		return resp, http.StatusUnauthorized
	}

	resp := RespNotOk{
		Ok: false,
	}
	return resp, http.StatusInternalServerError

}

type KycCheckerResponse1 struct {
	Message string `json:"message"`
}

type RespError struct {
	Error string `json:"error"`
}

type RespOk struct {
	Message string `json:"message"`
}

type RespNotOk struct {
	Ok bool `json:"ok"`
}
