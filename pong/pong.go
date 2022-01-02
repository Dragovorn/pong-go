package pong

import (
	"github.com/dragovorn/go-pong/generated/assets"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/isshoni-soft/roxxy"
	"github.com/isshoni-soft/sakura"
	"github.com/isshoni-soft/sakura/event"
	"github.com/isshoni-soft/sakura/event/events"
	"github.com/isshoni-soft/sakura/input"
	"github.com/isshoni-soft/sakura/render"
	"github.com/isshoni-soft/sakura/window"
)

var pong *Pong

var version = sakura.Version{
	Major:    0,
	Minor:    0,
	Patch:    2,
	Snapshot: true,
}

type Pong struct {
	sakura.Game

	logger        *roxxy.Logger
	shaderProgram *render.ShaderProgram
	vao           uint32
	initialized   bool
}

func Init() (result *Pong) {
	if pong == nil {
		pong = new(Pong)
		result = pong
		result.logger = roxxy.NewLogger("Pong>")
	} else {
		result = pong
	}

	return result
}

func (p *Pong) PreInit() {
	if version.Snapshot {
		p.logger.Log("Snapshot version detected! Enabling debug mode...")
	}

	title := "Pong"

	if version.Snapshot {
		title = title + " v" + version.GetVersion()
	}

	event.RegisterListener(event.Listener{
		IgnoreCancelled: false,
		Async:           true,
		Priority:        event.ASAP,
		Function: func(event *event.Event) {
			eventData := event.Data.(input.KeyEventData)

			if eventData.Action == glfw.Press {
				p.logger.Log("Press: " + eventData.KeyName)
			} else if eventData.Action == glfw.Repeat {
				p.logger.Log("Repeat: " + eventData.KeyName)
			} else if eventData.Action == glfw.Release {
				p.logger.Log("Release: " + eventData.KeyName)
			}
		},
	}, events.INPUT)

	window.SetTitle(title)
	sakura.SetDebug(version.Snapshot)
}

func (p *Pong) Init() {
	p.logger.Log("Initializing Pong v", Version().GetVersion())

	render.ClearColor(0.0, 0.0, 0.0, 1.0)
	render.Enable(gl.DEPTH_TEST)
	render.DepthFunc(gl.LESS)

	p.logger.Log("Compiling shaders...")

	data, _ := assets.Asset("shader/shader.vert")
	vertex := render.ShaderFromString(gl.VERTEX_SHADER, string(data))
	render.CompileShader(vertex)

	data, _ = assets.Asset("shader/shader.frag")
	fragment := render.ShaderFromStrings(gl.FRAGMENT_SHADER, string(data))
	render.CompileShader(fragment)

	p.shaderProgram = render.NewShaderProgram(vertex, fragment)

	render.LinkShaderProgram(p.shaderProgram)

	p.logger.Log("Compiling VBO...")

	var vbo uint32

	points := []float32{
		0.0, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}

	// Configure VBO
	render.GenBuffers(1, &vbo)
	render.BindBuffer(gl.ARRAY_BUFFER, vbo)
	render.BufferData(gl.ARRAY_BUFFER, len(points)*4, gl.Ptr(points), gl.STATIC_DRAW)

	render.GenVertexArrays(1, &p.vao)
	render.BindVertexArray(p.vao)
	render.EnableVertexAttribArray(0)
	render.BindBuffer(gl.ARRAY_BUFFER, vbo)
	render.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	p.SetInitialized(true)
}

func (p *Pong) PostInit() {}

func (p *Pong) Tick() {}

func (p *Pong) Clear() {
	render.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (p *Pong) Draw() {
	render.UseShaderProgram(p.shaderProgram)
	render.BindVertexArray(p.vao)
	render.DrawArrays(gl.TRIANGLES, 0, 3)
}

func (p *Pong) SetInitialized(initialized bool) {
	p.initialized = initialized
}

func (p *Pong) Initialized() bool {
	return p.initialized
}

func Version() *sakura.Version {
	return &version
}
