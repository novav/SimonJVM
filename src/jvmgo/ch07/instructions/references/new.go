package references 

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

// Create new Object
/*
    Chap 6.6.1
    6.6.1　new指令
    xinxin.shi
    2017-08-03 23:04:15
*/
type NEW  struct { base.Index16Instruction }


func (self *NEW) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
    class := classRef.ResolvedClass()
    if class.IsInterface() || class.IsAbstract() {
        panic("java.lang.InstantiationError")
    }
    ref := class.NewObject()
    frame.OperandStack().PushRef(ref)
}
