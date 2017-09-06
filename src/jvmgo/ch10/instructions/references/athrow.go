package  references

import "reflect"
import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

// Throw exception or error
type ATHROW struct { base.NoOperandsInstruction }

func (self *ATHROW) Executor(frame *rtda.Frame) {
    ex := frame.OpreandStack().PopRef()
    if ex == nil {
       panic("java.lang.NullPointerException")
    }
    thread := frame.Thread();
    if !findAndGotoExceptionHandler(thread, ex) {
        handleUncaughtException(thread, ex)
    }
}