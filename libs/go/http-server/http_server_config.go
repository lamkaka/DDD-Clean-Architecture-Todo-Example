package http_server

type Config struct {
	ServerPort        string
	LogFilter         LogFilter
	ErrorHandler      ErrorHandler
	StreamRequestBody bool
}
