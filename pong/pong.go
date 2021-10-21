package pong

import (
	"github.com/isshoni-soft/sakura"
	"github.com/isshoni-soft/sakura/logging"
	"github.com/isshoni-soft/sakura/window"
)

var pong *Pong

var version = sakura.Version {
	Major: 0,
	Minor: 0,
	Patch: 1,
	Snapshot: true,
}

type Pong struct {
	logger *logging.Logger
}

func Init() (result *Pong) {
	if pong == nil {
		pong = new(Pong)
		result = pong
		result.logger = logging.NewLogger("Pong", 10)
	}

	return
}

func (p Pong) PreInit() {
	if version.Snapshot {
		p.logger.Log("Snapshot version detected! Enabling debug mode...")
	}

	title := "Pong"

	if version.Snapshot {
		title = title + " v" + version.GetVersion()
	}

	window.SetTitle(title)
	sakura.SetDebug(version.Snapshot)
}

func (p Pong) Init() {
	p.logger.Log("Initializing Pong v", version.GetVersion())

	go p.ticker()
}

func (p Pong) ticker() {
	defer sakura.Shutdown()

	for !window.ShouldClose() {
		window.SwapBuffers()
		window.PollEvents()
	}
}

//func (p Pong) PostInit() {
//
//}

func (p Pong) tick() {

}

func Version() sakura.Version {
	return version
}