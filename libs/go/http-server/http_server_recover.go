package http_server

import (
	"fmt"
	"libs/errors"

	"github.com/gofiber/fiber/v2"
)

func newRecoverMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(errors.ErrorUnexpected, fmt.Sprintf("Panic: %v", r))
				fmt.Errorf("err: %v", err)
			}
		}()
		return ctx.Next()
	}
}
