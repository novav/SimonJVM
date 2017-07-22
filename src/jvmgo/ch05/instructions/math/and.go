package math 
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
    5.7.3　布尔运算指令
    xinxin.shi
    2017-06-16 23:11:28
*/

type IAND struct { base.NoOperandsInstruction }

func (self *IAND) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    result := v1 & v2
    stack.PushInt(result)
}

// Boolean AND long
type LAND struct { base.NoOperandsInstruction }
 
func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
