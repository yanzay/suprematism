//+build js,wasm

package main

import (
	"math"
	"syscall/js"

	"github.com/yanzay/suprematism/bindings/webgl"
)

func main() {
	doc := js.Global().Get("document")

	// Prepare canvas
	canvas := doc.Call("createElement", "canvas")
	canvas.Call("setAttribute", "style", `width: 800px; height: 800px;`)
	devicePixelRatio := js.Global().Get("devicePixelRatio").Float()
	canvasWidth := int(math.Round(800 * devicePixelRatio))
	canvasHeight := int(math.Round(800 * devicePixelRatio))
	canvas.Set("width", canvasWidth)
	canvas.Set("height", canvasHeight)
	doc.Get("body").Call("appendChild", canvas)

	// Get WebGL context
	gl, err := webgl.NewContext(canvas, webgl.DefaultAttributes())
	if err != nil {
		panic(err)
	}

	// Data to draw
	vertices := []float32{
		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, js.TypedArrayOf(vertices), gl.STATIC_DRAW)

	// Shaders

	// Vertex shader
	vertexShaderSource := `
	attribute vec3 vCoord;
	void main(void) {
		gl_Position = vec4(vCoord, 1.0);
	}`
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader, vertexShaderSource)
	gl.CompileShader(vertexShader)

	// Fragment shader
	fragmentShaderCode := `
	void main(void) {
		gl_FragColor = vec4(0.0, 0.0, 0.0, 1.0);
	}`
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragmentShader, fragmentShaderCode)
	gl.CompileShader(fragmentShader)

	// Shader program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	gl.UseProgram(program)

	// Bind data to shader
	var coord = gl.GetAttribLocation(program, "vCoord")
	gl.VertexAttribPointer(coord, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(coord)

	// Clear
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Set viewport
	gl.Viewport(0, 0, canvasWidth, canvasHeight)

	// Draw the triangles
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}
