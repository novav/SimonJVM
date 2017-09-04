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
func init() {
    native.Register("java/lang/Throwable", "fillInStackTrace",
        "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *rtda.Frame) {
    // 10.5 
}