package handlers

import "errors"

var (
	ErrNoOrderUserIDProvided = errors.New("no user id provided")
)

type FailResponse struct {
	ErrorReason string `json:"error_reason"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
