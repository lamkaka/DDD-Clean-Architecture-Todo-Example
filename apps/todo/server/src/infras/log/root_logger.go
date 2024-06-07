package log

type RootLogger Logger

func NewRootLogger(config Config) (RootLogger, error) {
	return New(Config(config), "root")
}
