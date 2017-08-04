package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

/*
    Chap 6.6.2
    putstatic & getstatic指令
    xinxin.shi
    2017-08-04 22:02:50
*/
// Set static field in class

type PUT_STATIC struct { base.Index16Instruction }


func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
    currentMethod := frame.Method()
    currentClass := currentMethod.Class()
    cp := currentClass.ConstantPool()
    fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
    field := fieldRef.ResovledField()
    class := field.Class()
    descriptor := field.Descriptor()
    slotId := field.SlotId()
    slots := class.StaticVars()
    stack := frame.OperandStack()
    switch descriptor[0] {
        case "Z", "B", "C", "S", "I": slots.SetInt(slotId, stack.PopInt())
        case "F": slots.SetFloat(slotId, stack.PopFloat())
        case "J": slots.SetLong(slotId, stack.PopLong())
        case "D": slots.SetDouble(slotId, stack.PopDouble())
        case "L", "[": slots.SetRef(slotId, stack.PopRef())
    }
}