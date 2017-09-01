package lang

import "unsafe"
import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"

/*
    chap9.3.5
    通过反射获取类名
    xinxin.shi
    2017-08-27 14:40:47
*/
const jlObject = "java/lang/Object"

func init() {
    native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
    native.Register(jlObject, "hashCode", "()I", hashCode)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
    this := frame.LocalVars().GetThis()
    class := this.Class().JClass()
    frame.OperandStack().PushRef(class)
}

func hashCode(frame *rtda.Frame) {
    this := frame.LocalVars().GetThis()
    hash := int32(uintptr(unsafe.Pointer(this)))
    frame.OperandStack().PushInt(hash)
}