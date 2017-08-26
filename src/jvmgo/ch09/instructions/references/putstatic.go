package references

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

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
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}