package lang

import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"

/*
    chap9.3.5
    通过反射获取类名
    xinxin.shi
    2017-08-27 14:40:47
*/
func init() {
    native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
    this := frame.LocalVars().GetThis()
    class := this.Class().JClass()
    frame.OperandStack().PushRef(class)
}