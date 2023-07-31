package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	httpClientGoLib "github.com/trustwallet/go-libs/client"
	"net/http"
)

const (
	KycCheckerBaseUrl   = "https://coding-challenge.xeptore.me"
	KycCheckerSecretKey = "mZjMs7-Ci3wqXaFtI5FdhEqAb8Z8YkeYOOmmorinEHVf0bZHn_DCnM7oItT"
)

func handleVerify(c echo.Context) error {
	r := &VerifyRequest{}
	err := c.Bind(r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	err = validate.Struct(r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	f := &FullName{FullName: fmt.Sprintf("%s %s", r.Name, r.Family)}
	err = validate.Struct(f)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "maximum of character is 20")
	}

	resp, err := checkKyc(c.Request().Context(), f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to check KYC")
	}

	return c.JSON(http.StatusOK, resp)

}

type VerifyRequest struct {
	Name   string `json:"firstName" form:"firstName" validate:"required"`
	Family string `json:"lastName" form:"lastName" validate:"required"`
}
type FullName struct {
	FullName string `json:"name" validate:"max=20"`
}

func checkKyc(ctx context.Context, data *FullName) (*KycCheckerResponse, error) {

	req := httpClientGoLib.InitJSONClient(KycCheckerBaseUrl, nil)

	nb := httpClientGoLib.NewReqBuilder()
	nb.Headers(map[string]string{
		"Secret": KycCheckerSecretKey,
	}).Body(data).Method("POST").PathStatic("/verify")

	body, err := req.Execute(ctx, nb.Build())
	if err != nil {
		return nil, err
	}
	var response KycCheckerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type KycCheckerResponse struct {
}
