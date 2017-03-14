// Package test provides clients for testing Ios's performance.
package test

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/heidi-ann/ios/config"
	"math/rand"
	"time"
)

// Generator generates workloads for the store
// Store has 10 keys
type Generator struct {
	Config config.ConfigAuto
	Keys   []string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Generate(config config.WorkloadConfig) *Generator {
	keys := make([]string, config.Config.Keys)
	for i := range keys {
		keys[i] = RandStringBytes(config.Config.KeySize)
	}
	return &Generator{config.Config, keys}
}

func (g *Generator) Next() (string, bool) {

	//handle termination after n requests
	if g.Config.Requests == 0 {
		return "", false
	}
	g.Config.Requests--

	delay := 0
	if g.Config.Interval > 0 {
		delay = rand.Intn(g.Config.Interval)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// generate key
	key := g.Keys[rand.Intn(g.Config.Keys)]
	glog.V(1).Info("Key is ", key)

	if rand.Intn(100) < g.Config.Reads {
		return fmt.Sprintf("get %s", key), true
	} else {
		value := RandStringBytes(g.Config.ValueSize)
		return fmt.Sprintf("update %s %s", key, value), true
	}
}

func (_ *Generator) Return(_ string) {
	//STUB
}
