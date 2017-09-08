package lang

import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

/*
    Chap10.2
    Throwable
    xinxin.shi
    2017-09-04 22:36:45
*/

type StackTraceElement struct {
    fileName        string
    className       string
    methodName      string
    lineNumber      int
}

func init() {
    native.Register("java/lang/Throwable", "fillInStackTrace",
        "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *rtda.Frame) {
    // 10.5 
    this := frame.LocalVars().GetThis()
    frame.OperandStack().PushRef(this)
    stes := createStackTraceElements(this, frame.Thread())
    this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) {
    skip := distanceToObject(tObj.Class()) + 2
    frames := thread.GetFrames()[skip:]
    stes := make([]*StackTraceElement, len(frames))
    for i, frame := range frames {
        stes[i] = createStackTraceElements(frame)
    }
    return stes
}