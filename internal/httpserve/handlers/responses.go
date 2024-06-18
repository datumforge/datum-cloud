package handlers

import (
	"errors"
	"fmt"
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/rout"
)

var (
	// ErrInvalidInput is returned when the input is invalid
	ErrInvalidInput = errors.New("invalid input")

	// ErrConflict is returned when the request cannot be processed due to a conflict
	ErrConflict = errors.New("conflict")
)

var (
	// InvalidInputErrCode is returned when the input is invalid
	InvalidInputErrCode rout.ErrorCode = "INVALID_INPUT"
)

// InternalServerError returns a 500 Internal Server Error response with the error message.
func (h *Handler) InternalServerError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err)); err != nil {
		return err
	}

	return err
}

// Unauthorized returns a 401 Unauthorized response with the error message.
func (h *Handler) Unauthorized(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusUnauthorized, rout.ErrorResponse(err)); err != nil {
		return err
	}

	return err
}

// NotFound returns a 404 Not Found response with the error message.
func (h *Handler) NotFound(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusNotFound, rout.ErrorResponse(err)); err != nil {
		return err
	}

	return err
}

// BadRequest returns a 400 Bad Request response with the error message.
func (h *Handler) BadRequest(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err)); err != nil {
		return err
	}

	return err
}

// BadRequest returns a 400 Bad Request response with the error message.
func (h *Handler) BadRequestWithCode(ctx echo.Context, err error, code rout.ErrorCode) error {
	if err := ctx.JSON(http.StatusBadRequest, rout.ErrorResponseWithCode(err, code)); err != nil {
		return err
	}

	return err
}

// InvalidInput returns a 400 Bad Request response with the error message.
func (h *Handler) InvalidInput(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, rout.ErrorResponseWithCode(err, InvalidInputErrCode)); err != nil {
		return err
	}

	return err
}

// Conflict returns a 409 Conflict response with the error message.
func (h *Handler) Conflict(ctx echo.Context, err string, code rout.ErrorCode) error {
	if err := ctx.JSON(http.StatusConflict, rout.ErrorResponseWithCode(err, code)); err != nil {
		return err
	}

	return fmt.Errorf("%w: %v", ErrConflict, err)
}

// TooManyRequests returns a 429 Too Many Requests response with the error message.
func (h *Handler) TooManyRequests(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusTooManyRequests, rout.ErrorResponse(err)); err != nil {
		return err
	}

	return err
}

// Success returns a 200 OK response with the response object.
func (h *Handler) Success(ctx echo.Context, rep interface{}) error {
	return ctx.JSON(http.StatusOK, rep)
}

// Created returns a 201 Created response with the response object.
func (h *Handler) Created(ctx echo.Context, rep interface{}) error {
	return ctx.JSON(http.StatusCreated, rep)
}
