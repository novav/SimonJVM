package references 

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"


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
    fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
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
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
        val := stack.PopInt()
        ref := stack.PopRef()
        if ref == nil {
            panic("java.lang.NullPointerException")
        }
        ref.Fields().SetInt(slotId, val)
    case 'F': 
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}