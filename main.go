package main

import (
	"github.com/dragovorn/go-pong/pong"
	"github.com/isshoni-soft/sakura/sakura"
)

func main() {
	sakura.Init(pong.Init())
}
