package control

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
/*
    Chap 7.4
    返回指令
    xinxin.shi
    2017-08-12 00:59:43
*/
type RETURN struct { base.NoOperandsInstrucion }    // Return void from method
type ARETURN struct { base.NoOperandsInstrucion }   // Return reference from method
type DRETURN struct { base.NoOperandsInstrucion }   // Return doubke from method
type FRETURN struct { base.NoOperandsInstrucion }   // Return float from method
type IRETURN struct { base.NoOperandsInstrucion }   // Return int from method
type LRETURN struct { base.NoOperandsInstrucion }   // Return long from method

func (self *RETURN) Execute(frame *rtda.Frame) {
    frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    currentFrame := thread.PopFrame()
    invokerFrame := thread.TopFrame()
    retVal := currentFrame.OperandStack().PopInt()
    invokerFrame.OperandStack().PushInt(retVal)
}