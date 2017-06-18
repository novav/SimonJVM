package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*  
    5.8　类型转换指令
    xinxinshi
    2017-06-18 22:44:11
*/
type D2F struct { base.NoOperandsInstruction }
type D2I struct { base.NoOperandsInstruction }
type D2L struct { base.NoOperandsInstruction }

func (self *D2I) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    d := stack.PopDouble()
    i := int32(d)
    stack.PushInt(i)
}