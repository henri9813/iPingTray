package main

import (
	"runtime"

	probing "github.com/prometheus-community/pro-bing"
)

var pinger *probing.Pinger

func init() {
	var err error
	pinger, err = probing.NewPinger("8.8.8.8")
	if err != nil {
		panic(err)
	}

	pinger.SetLogger(nil)

	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
}
