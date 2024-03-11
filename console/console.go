package console

import "syscall/js"

var console = js.Global().Get("console")

func Log(args ...interface{}) {
	console.Call("log", args...)
}
