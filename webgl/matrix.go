package webgl

import (
	"github.com/andrylavr/webapi/utils"
	"syscall/js"
)

type Matrix []float32

func (m Matrix) ToJS() js.Value {
	v := []float32(m)
	return utils.SliceToTypedArray(v)
}
