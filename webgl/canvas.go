package webgl

import (
	"github.com/andrylavr/webapi/global"
	"syscall/js"
)

type Canvas struct {
	js.Value
}

func NewCanvas() Canvas {
	elem := global.Document.CreateElement("canvas")
	return Canvas{elem}
}

func GetCanvasById(id string) Canvas {
	elem := global.Document.Call("getElementById", id)
	return Canvas{elem}
}

func (c Canvas) SetWidth(w int) {
	c.Set("width", w)
}

func (c Canvas) SetHeight(h int) {
	c.Set("height", h)
}

func (c Canvas) GetWidth() int {
	return c.Get("width").Int()
}

func (c Canvas) GetHeight() int {
	return c.Get("height").Int()
}

func (c Canvas) GetWebGLContext() Context {
	ctx := Context{c.Call("getContext", "webgl")}
	return ctx
}
