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
	if nil != err {
		return false
	}
	return b
}

func (e *Env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.ParseInt(s, 10, 0)
	if nil != err {
		return 0
	}
	return int(i)
}

func (e *Env) GetFloat(name string) float64 {
	s := e.GetString(name)
	f, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return f
}

func (e *Env) GetDuration(name string) time.Duration {
	s := e.GetString(name)
	d, err := time.ParseDuration(s)
	if nil != err {
		return time.Duration(0)
	}
	return d
}

type EnvGetter struct{}

func (r *EnvGetter) GetString(name string) string {
	return os.Getenv(name)
}
