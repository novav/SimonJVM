package stack
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

/*
    chap 5.6.1 poo pop2 
    xinxin.shi
    2017-06-14 23:26:33
*/
// Pop the top operand stack value
type POP struct { base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopSlot()
    stack.PopSlot()
}

