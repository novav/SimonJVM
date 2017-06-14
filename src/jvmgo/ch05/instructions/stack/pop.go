package stack
import "jvmgo/ch05/instruction/base"
import "jvmgo/ch05/rtda"

/*
    chap 5.6.1 poo pop2 
    xinxin.shi
    2017-06-14 23:26:33
*/

type POP struct { base.NoOperandsInstruction }
type POP2 struct { base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.POPSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopSlot()
    stack.PopSlot()
}

