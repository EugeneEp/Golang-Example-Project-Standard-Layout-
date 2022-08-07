package registry

import (
	"github.com/sarulabs/di/v2"
	"go.uber.org/zap"
	"golang.org/x/sys/windows/registry"
	"projectname/internal/project/infrastructure/logger"
)

func Ctx(ctn di.Container) (*Context, error) {
	var log *zap.Logger

	if err := ctn.Fill(logger.BaseServiceName, &log); err != nil {
		return nil, err
	}

	return &Context{
		Log: log,
	}, nil
}

type Context struct {
	Log *zap.Logger
}

func (c *Context) GetStringValue(key registry.Key, path string, name string, def string) (string, error) {
	v, _, err := key.GetStringValue(name)

	if err == registry.ErrNotExist {
		return def, nil
	}

	if err != nil {
		c.Log.Warn(path, zap.NamedError(name, err))
		return def, nil
	}

	if v == "" {
		v = def
	}

	return v, nil
}

func (c *Context) SetStringValue(key registry.Key, path string, name, value string) error {
	if err := key.SetStringValue(name, value); err != nil {
		c.Log.Warn(path, zap.NamedError(name, err))
		return err
	}

	return nil
}

func (c *Context) GetIntValue(k registry.Key, path string, name string, def int) (int, error) {
	v, _, err := k.GetIntegerValue(name)

	if err == registry.ErrNotExist {
		return def, nil
	}

	if err != nil {
		c.Log.Warn(path, zap.NamedError(name, err))
		return def, nil
	}

	return int(v), nil
}

func (c *Context) GetBinaryValue(k registry.Key, path string, name string, def []byte) ([]byte, error) {
	v, _, err := k.GetBinaryValue(name)

	if err == registry.ErrNotExist {
		return def, nil
	}

	if err != nil {
		c.Log.Warn(path, zap.NamedError(name, err))
		return def, nil
	}

	return v, nil
}
