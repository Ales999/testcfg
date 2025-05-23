package hello

import (
	"context"
	"fmt"

	"encore.dev/beta/errs"
	"encore.dev/config"
)

// AppConfig represents the application configuration.
type AppConfig struct {
	Appmode config.Bool
}

// Response represents the response structure.
type Response struct {
	Message string
}

var cfg *AppConfig = config.Load[*AppConfig]()

//encore:api public path=/hello/:name
func World(ctx context.Context, name string) (*Response, error) {
	// Check if configuration is initialized
	if cfg == nil {
		return &Response{}, errs.B().Code(errs.Internal).Msg("configuration not initialized").Err()
	}

	// Determine configuration mode
	isConfigMode := cfg.Appmode()

	// Construct message based on configuration status
	var msg string
	if isConfigMode {
		msg = fmt.Sprintf("Config work TRUE, %s!", name)
	} else {
		msg = fmt.Sprintf("Config work FALSE, %s!", name)
	}

	return &Response{Message: msg}, nil
}
