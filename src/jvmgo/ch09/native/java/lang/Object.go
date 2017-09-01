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
    native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
    this := frame.LocalVars().GetThis()
    class := this.Class().JClass()
    frame.OperandStack().PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
    this := frame.LocalVars().GetThis()
    hash := int32(uintptr(unsafe.Pointer(this)))
    frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
    this := frame.LocalVars().GetThis()
    cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
    if !this.Class().IsImplements(cloneable) {
        panic("java.lang.CloneNotSupportedException")
    }
    frame.OperandStack().PushRef(this.Clone())
}