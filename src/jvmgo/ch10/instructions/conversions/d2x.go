package conversions

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

/*  
    5.8　类型转换指令
    xinxinshi
    2017-06-18 22:44:11
*/
// Convert double to float
type D2F struct { base.NoOperandsInstruction }
func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct { base.NoOperandsInstruction }

func (self *D2I) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    d := stack.PopDouble()
    i := int32(d)
    stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
