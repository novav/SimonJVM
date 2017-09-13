package reserved

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/native"
// 预加载指令
import _ "jvmgo/ch11/native/java/io"
import _ "jvmgo/ch11/native/java/lang"
import _ "jvmgo/ch11/native/java/security"
import _ "jvmgo/ch11/native/java/util/concurrent/atomic"
import _ "jvmgo/ch11/native/sun/misc"
import _ "jvmgo/ch11/native/sun/reflect"
/*
    Chap9.2 
    调用本地方法
    xinxin.shi
    2017-08-26 21:20:49
*/
type INVOKE_NATIVE struct { base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
    method := frame.Method()
    className := method.Class().Name()
    methodName := method.Name()
    methodDescriptor := method.Descriptor()
    nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
    if nativeMethod == nil {
        methodInfo := className + "." + methodName + methodDescriptor
        panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
    }
    nativeMethod(frame)
}