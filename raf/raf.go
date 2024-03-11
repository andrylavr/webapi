package raf

import (
	"github.com/andrylavr/webapi/global"
	"syscall/js"
)

var Channel = make(chan float64)
var callback js.Func

func cb() {
	global.Window.Call("requestAnimationFrame", callback)
}

func init() {
	callback = js.FuncOf(func(this js.Value, args []js.Value) any {
		cb()
		select {
		case Channel <- args[0].Float():
		default:
		}

		return nil
	})
	cb()
}
