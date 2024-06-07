// Read app configurations from environment variables.
package env

import (
	"context"
	"fmt"
	"libs/errors"

	"github.com/caarlos0/env/v6"
)

// Read environment variables and marshal them into golang structs.
type Parser interface {
	// Takes a struct pointer with field tags `env:"ENV_NAME"` and fills the fields with environment variables.
	Parse(ctx context.Context, cfg any) error
}

type parser struct {
}

func NewParser() Parser {
	return parser{}
}

func (p parser) Parse(ctx context.Context, cfg any) error {
	err := env.Parse(cfg)
	if err != nil {
		err = errors.New(errors.ErrorUnexpected, err.Error())
		fmt.Errorf("failed to parse, err: ", err)
		return err
	}
	fmt.Printf("parsed config:%+v", cfg)
	return nil
}
