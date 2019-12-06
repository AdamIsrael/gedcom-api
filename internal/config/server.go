package config

import "time"

type ServerConfiguration struct {
	Hostname     string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
