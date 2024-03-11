package global

import "syscall/js"

type JSBody struct {
	js.Value
}

var Body = JSBody{Document.Get("body")}

func (b JSBody) AppendChild(child js.Value) {
	b.Call("appendChild", child)
}
