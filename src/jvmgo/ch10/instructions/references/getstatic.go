package references 

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

/*
    Chap 6.2.1
    putstatic
    xinxin.shi
    2017-08-04 22:18:07
    */
// Get static field from class
type GET_STATIC struct { base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
    field := fieldRef.ResolvedField()
    class := field.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

    if !field.IsStatic() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    descriptor := field.Descriptor()
    slotId := field.SlotId()
    slots := class.StaticVars()
    stack := frame.OperandStack()
    switch descriptor[0] {
    case 'Z', 'B', 'C', 'S', 'I': stack.PushInt(slots.GetInt(slotId))
    case 'F': stack.PushFloat(slots.GetFloat(slotId))
    case 'J': stack.PushLong(slots.GetLong(slotId))
    case 'D': stack.PushDouble(slots.GetDouble(slotId))
    case 'L', '[': stack.PushRef(slots.GetRef(slotId))
    }
}