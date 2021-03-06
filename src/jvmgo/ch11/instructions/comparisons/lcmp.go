package comparisons

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

/*
    5.9.1　lcmp指令
    xinxin.shi
    2017-06-18 22:57:54
*/
// Compare long

type LCMP struct { base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopLong()
    v1 := stack.PopLong()
    if v1 > v2 {
        stack.PushInt(1)
    } else if v1 == v2 {
        stack.PushInt(0)
    } else {
        stack.PushInt(-1)
    }
}