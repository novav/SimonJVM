 package references
import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/class"

/*
    Chap7.5.1
    xinxin.shi
    2017-08-13 13:58:54
*/
// Invoke a class (static) method
type INVOKE_STATIC struct { base.Index16Instrucion }

func (self *INVOKE_STATIC) Execute (frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
    resolvedMethod := methodRef.ResolvedMethod()
    if !resolvedMethod.IsStatic() {
        panic("hava.lang.IncompatibleClassChangeError")
    }
    base.InvokeMethod(frame, resolvedMethod)
}
