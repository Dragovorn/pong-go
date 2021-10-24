package main

import (
	"github.com/dragovorn/go-pong/pong"
	"github.com/isshoni-soft/sakura"
	"github.com/isshoni-soft/sakura/logging"
)

func main() {
	logging.InitLogfile("pong/logs/")
	sakura.Init(pong.Init())
}
