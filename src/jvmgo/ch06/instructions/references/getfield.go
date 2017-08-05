package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

/*
    Chap6.6.3
    getfield
    xinxin.shi
    2017-08-05 22:08:28
*/
//Fetch field from object

type GET_FIELD struct { base.Indec16Instruction }

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
    field := fieldRef.ResolvedField()
    if field.IsStatic() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    stack := frame.OperandStack()
    ref := stack.PopRef()
    if ref == nil {
        pannic("java.lang.NullPointerException")
    }
    descriptor := field.Descriptor()
    slotId := field.SoltId()
    slots := ref.Fields()
    case descriptor[0] :
        csae 'Z', 'B', 'C', 'S', 'I' : stack.PushInt(slots.GetInt(slotId))
        case 'F': stack.PushFloat(slots.GetFloat(slotId))
        case 'J': stack.PushLong(slots.GetLong(slotId))
        case 'D': stack.PushDouble(slots.GetDouble(slotId))
        case 'L': stack.PushRef(slots.GetRef(slotId))
}