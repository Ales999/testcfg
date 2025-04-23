package hello

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/config"
)

type AppConfig struct {
	Appmode config.Bool
}

var cfg *AppConfig = config.Load[*AppConfig]()

//encore:api public path=/hello/:name
func World(ctx context.Context, name string) (*Response, error) {

	var _mode = cfg.Appmode()

	var msg = "Hello" + name + "!"
	var _msg string

	if cfg != nil {
		if _mode {
			_msg = "Config work TRUE," + name + "!"
		} else {
			_msg = "Config work FALSE," + name + "!"
		}

	} else {
		return &Response{}, errs.B().Code(errs.InvalidArgument).Msg("cfg is not initialised").Err()
	}
	if len(_msg) > 0 {
		return &Response{Message: _msg}, nil
	}
	return &Response{Message: msg}, nil

}

type Response struct {
	Message string
}
