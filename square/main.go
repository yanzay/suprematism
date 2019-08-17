//+build js,wasm

package main

import (
	"math"
	"syscall/js"
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
	gl := canvas.Call("getContext", "webgl")

	// Data to draw
	vertices := []float32{
		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	vertexBuffer := gl.Call("createBuffer")
	gl.Call("bindBuffer", gl.Get("ARRAY_BUFFER"), vertexBuffer)
	gl.Call("bufferData", gl.Get("ARRAY_BUFFER"), js.TypedArrayOf(vertices), gl.Get("STATIC_DRAW"))

	// Shaders

	// Vertex shader
	vertexShaderSource := `
	attribute vec3 vCoord;
	void main(void) {
		gl_Position = vec4(vCoord, 1.0);
	}`
	vertexShader := gl.Call("createShader", gl.Get("VERTEX_SHADER"))
	gl.Call("shaderSource", vertexShader, vertexShaderSource)
	gl.Call("compileShader", vertexShader)

	// Fragment shader
	fragmentShaderCode := `
	void main(void) {
		gl_FragColor = vec4(0.0, 0.0, 0.0, 1.0);
	}`
	fragmentShader := gl.Call("createShader", gl.Get("FRAGMENT_SHADER"))
	gl.Call("shaderSource", fragmentShader, fragmentShaderCode)
	gl.Call("compileShader", fragmentShader)

	// Shader program
	program := gl.Call("createProgram")
	gl.Call("attachShader", program, vertexShader)
	gl.Call("attachShader", program, fragmentShader)
	gl.Call("linkProgram", program)
	gl.Call("useProgram", program)

	// Bind data to shader
	var coord = gl.Call("getAttribLocation", program, "vCoord")
	gl.Call("vertexAttribPointer", coord, 3, gl.Get("FLOAT"), false, 0, 0)
	gl.Call("enableVertexAttribArray", coord)

	// Clear
	gl.Call("clearColor", 1, 1, 1, 1)
	gl.Call("clear", gl.Get("COLOR_BUFFER_BIT"))

	// Set viewport
	gl.Call("viewport", 0, 0, canvasWidth, canvasHeight)

	// Draw the triangles
	gl.Call("drawArrays", gl.Get("TRIANGLE_STRIP"), 0, 4)
}
