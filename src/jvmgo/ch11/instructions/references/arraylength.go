package references

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

/*
    8.3.3　
    arraylength指令
    xinxin.shi
    2017-08-19 17:54:40
*/
// Get length of array

type ARRAY_LENGTH struct { base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    arrRef := stack.PopRef()
    if arrRef == nil {
        panic("java.lang.NullPointerException")
    }
    arrLen := arrRef.ArrayLength()
    stack.PushInt(arrLen)
}