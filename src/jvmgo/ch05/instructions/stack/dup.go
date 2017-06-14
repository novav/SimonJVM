package stack
import "jvmgo/ch05/instrctions/base"
import "jvmgo/ch05/rtda"

/*
    chap 5.6.2
    xinxin.shi 
    2017-06-14 23:38:19
*/

type DUP struct { base.NoOperandsInstruction }
type DUP_X1 struct { base.NoOperandsInstruction }
type DUP_X2 struct { base.NoOperandsInstruction }
type DUP2 struct { base.NoOperandsInstruction }
type DUP2_X1 struct { base.NoOperandsInstruction }
type DUP2_X2 struct { base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    slot := stack.PopSlot()
    stack.PushSlot(slot)
    stack.PushSlot(slot)
}
