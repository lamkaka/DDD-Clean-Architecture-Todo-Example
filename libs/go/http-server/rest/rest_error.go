// Defines unified formats for restful requests and responses in demo projects
package rest

import (
	"libs/errors"

	"dario.cat/mergo"
	"github.com/gofiber/fiber/v2"
)

type ErrorStatusMap map[errors.ErrorCode]int

// Creates a ErrorHandler that catches demo errors and respond with them as json body.
// It will map the common demo error codes to http response status.
// In order to also map custom-defined error codes, pass in a ErrorStatusMap.
// e.g. ErrorStatusMap{ErrorNotPaid: 402}
func NewErrorHandler(errStatusMap ErrorStatusMap) (fiber.ErrorHandler, error) {
	mergedErrStatusMap := defaultErrStatusMap
	err := mergo.Merge(&mergedErrStatusMap, errStatusMap)
	if err != nil {
		return nil, err
	}

	handler :=
		func(ctx *fiber.Ctx, err error) error {
			httpErr, isHttpErr := err.(*fiber.Error)
			if isHttpErr {
				return ctx.Status(httpErr.Code).SendString(httpErr.Message)
			}

			demoErr, isSilicaErr := err.(errors.Error)
			if isSilicaErr {
				status := mergedErrStatusMap[demoErr.Code()]
				return ctx.Status(status).JSON(errToResponse(demoErr))
			}

			return ctx.Status(500).JSON(errResponse{
				Code:     errors.ErrorUnexpected,
				Messages: []string{err.Error()},
			})
		}

	return handler, nil
}

type errResponse struct {
	Code     errors.ErrorCode `json:"code"`
	Messages []string         `json:"messages"`
}

func errToResponse(err errors.Error) errResponse {
	return errResponse{Code: err.Code(), Messages: err.Messages()}
}

var defaultErrStatusMap = ErrorStatusMap{
	errors.ErrorUnauthorized:     401,
	errors.ErrorEntityNotFound:   404,
	errors.ErrorEntityConflict:   409,
	errors.ErrorEntityValidation: 422,
	errors.ErrorUnexpected:       500,
}

// a reverse map of defaultErrStatusMap.
// It is initialized in init() and used in StatusToError()
var defaultStatusErrMap map[int]errors.ErrorCode

func init() {
	defaultStatusErrMap = make(map[int]errors.ErrorCode)
	for code, status := range defaultErrStatusMap {
		defaultStatusErrMap[status] = code
	}
}

func StatusToError(status int) errors.ErrorCode {
	err, ok := defaultStatusErrMap[status]
	if !ok {
		return errors.ErrorUnexpected
	}
	return err
}
