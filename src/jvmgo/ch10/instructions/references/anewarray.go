package references
/*
    Chap8.3.2
    anewarray指令
    xinxin.shi
    2017-08-18 21:32:01
*/
import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

//Create new array of reference

type ANEW_ARRAY struct { base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
    componentClass := classRef.ResolvedClass()
    stack := frame.OperandStack()
    count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
    }
    arrClass := componentClass.ArrayClass()
    arr := arrClass.NewArray(uint(count))
    stack.PushRef(arr)
}