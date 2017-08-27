/*
    Chap 6.6.5
    xinxin.shi
    2017-08-05 23:21:54
*/

package constants

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

type LDC struct { base.Index8Instruction }
type LDC_W struct { base.Index16Instruction }
type LDC2_W struct { base.Index16Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}
func (self *LDC_W) Execute(frame *rtda.Frame) {
    _ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
    stack := frame.OperandStack()
    class := frame.Method().Class()
    c := class.ConstantPool().GetConstant(index)
    switch c.(type) {
        case int32: 
            stack.PushInt(c.(int32))
        case float32: 
            stack.PushFloat(c.(float32))
        case string:    //   --> chap8
            internedStr := heap.JString(class.Loader(), c.(string))
            stack.PushRef(internedStr)
        case *heap.ClassRef:    //  --> chap9
            classRef := c.(*heap.ClassRef)
            classObj := classRef.ResolvedClass().JClass()
            stack.PushRef(classObj)
	    // case MethodType, MethodHandle
        default: panic("todo: ldc!")
    }
}


func (self *LDC2_W) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(self.Index)
    switch c.(type) {
        case int64: stack.PushLong(c.(int64))
        case float64: stack.PushDouble(c.(float64))
        default: panic ("java.lang.ClassFormatError")
    }
}