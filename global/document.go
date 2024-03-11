package global

import "syscall/js"

type JSDocument struct {
	js.Value
}

var Document = JSDocument{Window.Get("document")}

func (d JSDocument) AppendChild(child js.Value) {
	d.Call("appendChild", child)
}

func (d JSDocument) CreateElement(tag string) js.Value {
	return d.Call("createElement", tag)
}
