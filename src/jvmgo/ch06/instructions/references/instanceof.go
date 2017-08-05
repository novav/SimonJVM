package  references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Determine if object is of javen type 
type INSTANCE_OF struct { base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    ref := stackPopRef()
    if ref == nil {
        stack.PushInt(0)
        return 
    }
    cp := frame.Method().Class().ConstantPool()
    classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
    if ref.InInstanceOf(class) {
        stack.PushInt(1)
    } else {     
        stack.PushInt(0)
    }
}