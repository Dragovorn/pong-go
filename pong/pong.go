package pong

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/isshoni-soft/sakura"
	"github.com/isshoni-soft/sakura/logging"
	"github.com/isshoni-soft/sakura/render"
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

	render.ClearColor(1.0, 1.0, 1.0, 1.0)
	render.Enable(gl.DEPTH_TEST)
	render.DepthFunc(gl.LESS)
}

func (p Pong) PostInit() { }

func (p Pong) Tick() {
	//p.logger.Log("Tick")
}

func (p Pong) Clear() {
	render.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (p Pong) Draw() {
	points := []float32 {
		0.0, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}

	var vbo uint32
	var vao uint32
	//var vertex_shader TODO - SHADERS

	// Configure VBO
	render.GenBuffers(1, &vbo)
	render.BindBuffer(gl.ARRAY_BUFFER, vbo)
	render.BufferData(gl.ARRAY_BUFFER, len(points) * 4, gl.Ptr(points), gl.STATIC_DRAW)

	render.GenVertexArrays(1, &vao)
	render.BindVertexArray(vao)
	render.EnableVertexAttribArray(0)
	render.BindBuffer(gl.ARRAY_BUFFER, vbo)
	render.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func Version() sakura.Version {
	return version
}
