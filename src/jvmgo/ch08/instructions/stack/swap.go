package stack
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

/*
    chap 5.6.3
    xinxin.shi
    2017-06-14 23:44:50    
*/

// Swap the Top two operand stack values

type SWAP struct { base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame){
    stack := frame.OperandStack()
    slot1 := stack.PopSlot()
    slot2 := stack.PopSlot()
    stack.PushSlot(slot1)
    stack.PushSlot(slot2)
}