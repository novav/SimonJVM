package lang

import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"
/*
    chap9.3.5
    通过反射获取类名
    xinxin.shi
    2017-08-27 14:40:47
*/

func init() {
    native.Register("java/lang/Class", "getPrimitiveClass", 
        "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
    native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
    native.Register("java/lang/Class", "desiredAssertionStatus0",
         "(Ljava/lang/String;)Z", desiredAssertionStatus0)
}

// static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtda.Frame) {
    nameObj := frame.LocalVars().GetRef(0)
    name := heap.GoString(nameObj)
    loader := frame.Method().Class().Loader()
    calss := loader.LoadClass(name).JClass()
    frame.OperandStack().PushRef(calss)
}

// private native String getName0();
func getName0(frame *heap.Frame) {
    this := frame.LocalVars().GetThis()
    class := this.Extra().(*heap.Class)
    name := class.JavaName()    //java.lang.Object这样的类名
    nameObj := heap.JString(class.Loader(), name)
    frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus0(frame *rtda.Frame) {
    frame.OperandStack().PushBoolean(false)
}