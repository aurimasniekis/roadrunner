package rpc

import (
	"errors"
	"github.com/spiral/roadrunner/service"
	"net"
	"strings"
	"syscall"
)

// Config defines RPC service config.
type Config struct {
	// Indicates if RPC connection is enabled.
	Enable bool

	// AddListener string
	Listen string
}

// Hydrate must populate Config values using given Config source. Must return error if Config is not valid.
func (c *Config) Hydrate(cfg service.Config) error {
	if err := cfg.Unmarshal(c); err != nil {
		return err
	}

	return c.Valid()
}

// Valid returns nil if config is valid.
func (c *Config) Valid() error {
	if dsn := strings.Split(c.Listen, "://"); len(dsn) != 2 {
		return errors.New("invalid socket DSN (tcp://:6001, unix://rpc.sock)")
	}

	if dsn := strings.Split(c.Listen, "://"); len(dsn) != 2 {
		return errors.New("invalid socket DSN (tcp://:6001, unix://rpc.sock)")
	}

	return nil
}

// Listener creates new rpc socket Listener.
func (c *Config) Listener() (net.Listener, error) {
	dsn := strings.Split(c.Listen, "://")
	if len(dsn) != 2 {
		return nil, errors.New("invalid socket DSN (tcp://:6001, unix://rpc.sock)")
	}

	if dsn[0] == "unix" {
		syscall.Unlink(dsn[1])
	}

	return net.Listen(dsn[0], dsn[1])
}

// Dialer creates rpc socket Dialer.
func (c *Config) Dialer() (net.Conn, error) {
	dsn := strings.Split(c.Listen, "://")
	if len(dsn) != 2 {
		return nil, errors.New("invalid socket DSN (tcp://:6001, unix://rpc.sock)")
	}

	return net.Dial(dsn[0], dsn[1])
}
