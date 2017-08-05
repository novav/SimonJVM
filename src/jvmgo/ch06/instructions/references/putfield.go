package references 

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"


/*
    Chap6.6.3
    putfield指令
    xinxin.shi
    2017-08-05 21:57:15
*/
// Set field in object

type PUT_FIELD struct { base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
    currentMethod := frame.Method()
    currentClass := currentMethod.Class()
    cp := currentClass.ConstantPool()
    fieldRef := cp.GetConstant(self.index).(*heap.FieldRef)
    field := fieldRef.ResolvedField()
    if field.IsStatic() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    if field.IsFinal() {
        if currentClass != field.Class() || currentMethod.Name() != "<init>" {
            panic("java.lang.IllegalAccessError")
        }
    }
    descriptor := field.Descriptor()
    slotId := field.SlotId()
    stack := frame.OperandStack()
    switch 'Z', 'B', 'C', 'S', 'I':
        val := stack.PopInt()
        ref := stack.PopRef()
        if ref == nil {
            panic("java.lang.NullPointerException")
        }
        ref.Fields().SetInt(slotId, val)
    case 'F': 
    case 'J': 
    case 'D': 
    case 'L', '[': 
}