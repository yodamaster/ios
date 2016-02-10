package config

import (
	"github.com/golang/glog"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Addresses struct {
		Address []string
	}
}

func Parse(filename string) Config {
	var config Config
	err := gcfg.ReadFileInto(&config, filename)
	if err != nil {
		glog.Fatalf("Failed to parse gcfg data: %s", err)
	}
	return config
}