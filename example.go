package main

import (
	"github.com/andrylavr/webapi/console"
	"github.com/andrylavr/webapi/global"
	"github.com/andrylavr/webapi/raf"
	"github.com/andrylavr/webapi/webgl"
	"math"
)

func main() {
	canvas, gl := getCanvas()
	test := &CubeTest1{}
	test.RunCube(canvas, gl)

	select {}
}

func getCanvas() (canvas webgl.Canvas, gl webgl.Context) {
	canvas = webgl.NewCanvas()
	canvas.SetWidth(640)
	canvas.SetHeight(480)
	global.Body.AppendChild(canvas.Value)
	gl = canvas.GetWebGLContext()
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(webgl.COLOR_BUFFER_BIT)

	return
}

type CubeTest1 struct {
	proj_matrix  webgl.Matrix
	mov_matrix   webgl.Matrix
	view_matrix  webgl.Matrix
	canvas       webgl.Canvas
	gl           webgl.Context
	Pmatrix      webgl.UniformLocation
	Vmatrix      webgl.UniformLocation
	Mmatrix      webgl.UniformLocation
	index_buffer webgl.Buffer
	indices      []uint16
}

func (this *CubeTest1) RunCube(canvas webgl.Canvas, gl webgl.Context) {

	this.canvas = canvas
	this.gl = gl

	vertices := []float32{
		-1, -1, -1, 1, -1, -1, 1, 1, -1, -1, 1, -1,
		-1, -1, 1, 1, -1, 1, 1, 1, 1, -1, 1, 1,
		-1, -1, -1, -1, 1, -1, -1, 1, 1, -1, -1, 1,
		1, -1, -1, 1, 1, -1, 1, 1, 1, 1, -1, 1,
		-1, -1, -1, -1, -1, 1, 1, -1, 1, 1, -1, -1,
		-1, 1, -1, -1, 1, 1, 1, 1, 1, 1, 1, -1,
	}

	colors := []float32{
		5, 3, 7, 5, 3, 7, 5, 3, 7, 5, 3, 7,
		1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3,
		0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1,
		1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0,
		1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0,
		0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0,
	}

	indices := []uint16{
		0, 1, 2, 0, 2, 3, 4, 5, 6, 4, 6, 7,
		8, 9, 10, 8, 10, 11, 12, 13, 14, 12, 14, 15,
		16, 17, 18, 16, 18, 19, 20, 21, 22, 20, 22, 23,
	}
	this.indices = indices

	// Create and store data into vertex buffer
	var vertex_buffer = gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, vertex_buffer)
	gl.BufferData(webgl.ARRAY_BUFFER, vertices, webgl.STATIC_DRAW)

	console.Log("ARRAY_BUFFER", webgl.ARRAY_BUFFER)
	console.Log("STATIC_DRAW", webgl.STATIC_DRAW)

	// Create and store data into color buffer
	var color_buffer = gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, color_buffer)
	gl.BufferData(webgl.ARRAY_BUFFER, colors, webgl.STATIC_DRAW)

	// Create and store data into index buffer
	var index_buffer = gl.CreateBuffer()
	this.index_buffer = index_buffer
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, index_buffer)
	gl.BufferData(webgl.ELEMENT_ARRAY_BUFFER, indices, webgl.STATIC_DRAW)

	var vertCode = "attribute vec3 position;" +
		"uniform mat4 Pmatrix;" +
		"uniform mat4 Vmatrix;" +
		"uniform mat4 Mmatrix;" +
		"attribute vec3 color;" + //the color of the point
		"varying vec3 vColor;" +
		"void main(void) { " + //pre-built function
		"gl_Position = Pmatrix*Vmatrix*Mmatrix*vec4(position, 1.);" +
		"vColor = color;" +
		"}"

	var fragCode = "precision mediump float;" +
		"varying vec3 vColor;" +
		"void main(void) {" +
		"gl_FragColor = vec4(vColor, 1.);" +
		"}"

	var vertShader = gl.CreateShader(webgl.VERTEX_SHADER)
	gl.ShaderSource(vertShader, vertCode)
	gl.CompileShader(vertShader)

	var fragShader = gl.CreateShader(webgl.FRAGMENT_SHADER)
	gl.ShaderSource(fragShader, fragCode)
	gl.CompileShader(fragShader)

	var shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertShader)
	gl.AttachShader(shaderProgram, fragShader)
	gl.LinkProgram(shaderProgram)

	/* ====== Associating attributes to vertex shader =====*/
	this.Pmatrix = gl.GetUniformLocation(shaderProgram, "Pmatrix")
	this.Vmatrix = gl.GetUniformLocation(shaderProgram, "Vmatrix")
	this.Mmatrix = gl.GetUniformLocation(shaderProgram, "Mmatrix")

	gl.BindBuffer(webgl.ARRAY_BUFFER, vertex_buffer)
	var position = gl.GetAttribLocation(shaderProgram, "position")
	gl.VertexAttribPointer(position, 3, webgl.FLOAT, false, 0, 0)

	// Position
	gl.EnableVertexAttribArray(position)
	gl.BindBuffer(webgl.ARRAY_BUFFER, color_buffer)
	var color = gl.GetAttribLocation(shaderProgram, "color")
	gl.VertexAttribPointer(color, 3, webgl.FLOAT, false, 0, 0)

	// Color
	gl.EnableVertexAttribArray(color)
	gl.UseProgram(shaderProgram)

	w := float32(canvas.GetWidth())
	h := float32(canvas.GetHeight())
	this.proj_matrix = getProjection(40, w/h, 1, 100)

	this.mov_matrix = webgl.Matrix{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	this.view_matrix = webgl.Matrix{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}

	// translating z
	this.view_matrix[14] = this.view_matrix[14] - 6 //zoom

	this.RunAnimate()
}

var time_old = float32(0)

func (this *CubeTest1) Animate(time float32) {
	//fmt.Println("Animate", time)
	gl := this.gl

	var dt = time - time_old
	rotateZ(this.mov_matrix, dt*0.005) //time
	rotateY(this.mov_matrix, dt*0.002)
	rotateX(this.mov_matrix, dt*0.003)
	time_old = time

	gl.Enable(webgl.DEPTH_TEST)
	gl.DepthFunc(webgl.LEQUAL)
	gl.ClearColor(0.5, 0.5, 0.5, 0.9)
	gl.ClearDepth(1.0)

	gl.Viewport(0.0, 0.0, this.canvas.GetWidth(), this.canvas.GetHeight())
	gl.Clear(webgl.COLOR_BUFFER_BIT | webgl.DEPTH_BUFFER_BIT)
	gl.UniformMatrix4fv(this.Pmatrix, false, this.proj_matrix)
	gl.UniformMatrix4fv(this.Vmatrix, false, this.view_matrix)
	gl.UniformMatrix4fv(this.Mmatrix, false, this.mov_matrix)
	gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, this.index_buffer)
	gl.DrawElements(webgl.TRIANGLES, len(this.indices), webgl.UNSIGNED_SHORT, 0)
}

func getProjection(angle, a, zMin, zMax float32) webgl.Matrix {
	ang := float32(math.Tan(float64((angle * 0.5) * math.Pi / 180)))
	return webgl.Matrix{
		0.5 / ang, 0, 0, 0,
		0, 0.5 * a / ang, 0, 0,
		0, 0, -(zMax + zMin) / (zMax - zMin), -1,
		0, 0, (-2 * zMax * zMin) / (zMax - zMin), 0,
	}
}

func rotateZ(m []float32, angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	mv0, mv4, mv8 := m[0], m[4], m[8]

	m[0] = c*m[0] - s*m[1]
	m[4] = c*m[4] - s*m[5]
	m[8] = c*m[8] - s*m[9]

	m[1] = c*m[1] + s*mv0
	m[5] = c*m[5] + s*mv4
	m[9] = c*m[9] + s*mv8
}

func rotateX(m []float32, angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	mv1, mv5, mv9 := m[1], m[5], m[9]

	m[1] = m[1]*c - m[2]*s
	m[5] = m[5]*c - m[6]*s
	m[9] = m[9]*c - m[10]*s

	m[2] = m[2]*c + mv1*s
	m[6] = m[6]*c + mv5*s
	m[10] = m[10]*c + mv9*s
}

func rotateY(m []float32, angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	mv0, mv4, mv8 := m[0], m[4], m[8]

	m[0] = c*m[0] + s*m[2]
	m[4] = c*m[4] + s*m[6]
	m[8] = c*m[8] + s*m[10]

	m[2] = c*m[2] - s*mv0
	m[6] = c*m[6] - s*mv4
	m[10] = c*m[10] - s*mv8
}

func (cube *CubeTest1) RunAnimate() {
	for t := range raf.Channel {
		cube.Animate(float32(t))
	}
}
