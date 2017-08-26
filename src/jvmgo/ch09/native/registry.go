package native

import "jvmgo/ch09/rtda"

/*
    Chap9.1
    注册和查找本地方法
    xinxin.shi
    2017-08-26 20:17:45
*/
type NativeMtheod func {frame *rtda.Frame }

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
    key := className + "~" + methodName + "~" + methodDescriptor
    registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
    key := className + "~" + methodName + "~" + methodDescriptor
    if method, ok := registry[key]; ok {
        return method
    }
    if methodDescriptor == "()V" && methodName == "registerNatives" {
        return emptyNativeMathod
    }
    return nil
}

func emptyNativeMathod(frame *rtda.Frame) {
    //do nothing
}