package util

import (
	"os"
	"strconv"
	"time"
)

type EnvSource interface {
	GetString(name string) string
}

type Env struct {
	EnvSource
}

func (e *Env) GetString(name string) string {
	if nil == e.EnvSource {
		return ""
	}
	return e.EnvSource.GetString(name)
}

func (e *Env) GetBool(name string) bool {
	s := e.GetString(name)
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}

func (e *Env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int(i)
}

func (e *Env) GetFloat(name string) float64 {
	s := e.GetString(name)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func (e *Env) GetDuration(name string) time.Duration {
	s := e.GetString(name)
	d, err := time.ParseDuration(s)
	if err != nil {
		return time.Duration(0)
	}
	return d
}

func (e *Env) GetStringWithDefault(name string, value string) string {
	if nil == e.EnvSource {
		return value
	}
	return e.EnvSource.GetString(name)
}

func (e *Env) GetBoolWithDefault(name string, value bool) bool {
	s := e.GetString(name)
	b, err := strconv.ParseBool(s)
	if err != nil {
		return value
	}
	return b
}

func (e *Env) GetIntWithDefault(name string, value int) int {
	s := e.GetString(name)
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return value
	}
	return int(i)
}

func (e *Env) GetFloatWithDefault(name string, value float64) float64 {
	s := e.GetString(name)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return value
	}
	return f
}

func (e *Env) GetDurationWithDefault(name string, value time.Duration) time.Duration {
	s := e.GetString(name)
	d, err := time.ParseDuration(s)
	if err != nil {
		return value
	}
	return d
}

type EnvGetter struct{}

func (*EnvGetter) GetString(name string) string {
	return os.Getenv(name)
}
