package webgl

import (
	"github.com/andrylavr/webapi/utils"
	"syscall/js"
)

type Context struct {
	js.Value
}

func (c Context) Clear(mask GLbitfield) {
	c.Call("clear", mask)
}

func (c Context) CreateBuffer() Buffer {
	return Buffer{c.Call("createBuffer")}
}

func (c Context) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	c.Call("clearColor", red, green, blue, alpha)
}

func (c Context) ClearDepth(depth GLclampf) {
	c.Call("clearDepth", depth)
}

func (c Context) BindBuffer(target GLenum, buffer Buffer) {
	c.Call("bindBuffer", target, buffer.Value)
}

func (c Context) BufferData(target GLenum, data interface{}, usage GLenum) {
	var jsData = utils.SliceToTypedArray(data)
	c.Call("bufferData", target, jsData, usage)
}

func (c Context) CreateShader(t GLenum) Shader {
	return Shader{c.Call("createShader", t)}
}

func (c Context) ShaderSource(vertShader Shader, vertCode string) {
	c.Call("shaderSource", vertShader.Value, vertCode)
}

func (c Context) CompileShader(vertShader Shader) {
	c.Call("compileShader", vertShader.Value)
}

func (c Context) CreateProgram() Program {
	return Program{c.Call("createProgram")}
}

func (c Context) AttachShader(program Program, shader Shader) {
	c.Call("attachShader", program.Value, shader.Value)
}

func (c Context) LinkProgram(program Program) {
	c.Call("linkProgram", program.Value)
}

func (c Context) GetUniformLocation(program Program, name string) UniformLocation {
	return UniformLocation{c.Call("getUniformLocation", program.Value, name)}
}

func (c Context) GetAttribLocation(program Program, name string) GLint {
	return c.Call("getAttribLocation", program.Value, name).Int()
}

func (c Context) VertexAttribPointer(index GLuint, size GLint, t GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) {
	c.Call("vertexAttribPointer", index, size, t, normalized, stride, offset)
}

func (c Context) EnableVertexAttribArray(index GLuint) {
	c.Call("enableVertexAttribArray", index)
}

func (c Context) UseProgram(program Program) {
	c.Call("useProgram", program.Value)
}

func (c Context) Enable(cap GLenum) {
	c.Call("enable", cap)
}

func (c Context) DepthFunc(f GLenum) {
	c.Call("depthFunc", f)
}

func (c Context) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) {
	c.Call("viewport", x, y, width, height)
}

func (c Context) UniformMatrix4fv(location UniformLocation, transpose GLboolean, value Matrix) {
	c.Call("uniformMatrix4fv", location.Value, transpose, value.ToJS())
}

func (c Context) DrawElements(mode GLenum, count GLsizei, t GLenum, offset GLintptr) {
	c.Call("drawElements", mode, count, t, offset)
}
