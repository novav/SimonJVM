package stores 
/*
    Chap 8.3.5
    astore系列指令按索引取数组元素值
    xinxin.shi
    2017-08-19 18:15:11
*/
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type AASTORE struct { base.NoOperandsInstructon }
type BASTORE struct { base.NoOperandsInstructon }
type CASTORE struct { base.NoOperandsInstructon }
type DASTORE struct { base.NoOperandsInstructon }
type FASTORE struct { base.NoOperandsInstructon }
type IASTORE struct { base.NoOperandsInstructon }
type LASTORE struct { base.NoOperandsInstructon }
type SASTORE struct { base.NoOperandsInstructon }

func (self *IASTORE) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    val := stack.PopInt()
    index := stack.PopInt()
    arrRef := stack.PopRef()
    checkNotNil(arrRef)
    ints := arrRef.Ints()
    checkIndex(len(ints), index)
    ints[index] = int32(val)
}