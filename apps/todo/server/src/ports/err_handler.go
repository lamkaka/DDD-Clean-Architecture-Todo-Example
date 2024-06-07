package ports

import (
	"libs/http_server"
	"libs/http_server/rest"

	"dario.cat/mergo"
)

var ErrStatusMap = rest.ErrorStatusMap{}

func NewErrorHandler() (http_server.ErrorHandler, error) {
	var restErrStatusMap rest.ErrorStatusMap
	err := mergo.Merge(&restErrStatusMap, ErrStatusMap)
	if err != nil {
		return nil, err
	}

	errHandler, err := rest.NewErrorHandler(restErrStatusMap)
	if err != nil {
		return nil, err
	}

	return errHandler, nil
}
