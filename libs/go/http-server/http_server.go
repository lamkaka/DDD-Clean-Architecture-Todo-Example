// Provides a server that process http requests
package http_server

import (
	"context"
	"fmt"
	"libs/http_server/rest"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// A Router receives all http requests and route them to different Controllers to handle.
type Router = *fiber.App
type RequestContext = fiber.Ctx
type ResponseStatus = int
type Error = fiber.Error
type ErrorHandler = fiber.ErrorHandler

var NewError = fiber.NewError

type Server interface {
	Start(context.Context) error
}

type server struct {
	config Config
	router Router
}

func NewServer(config Config, controllers ...Controller) (Server, error) {
	if config.ErrorHandler == nil {
		restErrHandler, err := rest.NewErrorHandler(nil)
		if err != nil {
			return nil, err
		}
		config.ErrorHandler = restErrHandler
	}

	if config.LogFilter == nil {
		config.LogFilter = func(*RequestContext) bool { return true }
	}

	router := fiber.New(fiber.Config{ErrorHandler: config.ErrorHandler, StreamRequestBody: config.StreamRequestBody, DisablePreParseMultipartForm: config.StreamRequestBody})
	router.Use(cors.New())
	router.Use(newRecoverMiddleware())
	controllers = append([]Controller{newLogMiddleware(config)}, controllers...)
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
	return server{config, router}, nil
}

func (s server) Start(ctx context.Context) error {
	address := getAddress(s.config)
	fmt.Printf("Starting on %s", address)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-ch
		fmt.Printf("Shutting down from %s", sig)
		s.router.Shutdown()
	}()

	return s.router.Listen(address)
}
