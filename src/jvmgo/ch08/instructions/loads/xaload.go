package loads 
/*
    Chap 8.3.4
    aload系列指令按索引取数组元素值
    xinxin.shi
    2017-08-19 18:15:11
*/
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type AALOAD struct { base.NoOperandsInstructon }
type BALOAD struct { base.NoOperandsInstructon }
type CALOAD struct { base.NoOperandsInstructon }
type DALOAD struct { base.NoOperandsInstructon }
type FALOAD struct { base.NoOperandsInstructon }
type IALOAD struct { base.NoOperandsInstructon }
type LALOAD struct { base.NoOperandsInstructon }
type SALOAD struct { base.NoOperandsInstructon }

func (self *AALOAD) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    index := stack.PopInt()
    arrRef := stack.PopRef()
    checkNotNil(arrRef)
    refs :=  arrRef.Refs()
    checkIndex(len(refs), index)
    stack.PushRef(refs[index])
}

func checkNotNil(ref *heap.Object) {
    if ref == nil {
        panic("java.lang.NullPointerException")
    }
}

func checkIndex(arrLen int, index int32) {
    if index < 0 || index >= int32(arrLen) {
        panic("ArrayIndexOutOfBoundsException")
    }
}