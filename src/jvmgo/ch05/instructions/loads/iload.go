package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
    Chap；5.4
    xinxin.shi
    2017-06-13 23:49:45
*/

// Load init from local variable

type ILOAD struct { base.Index8Instruction }
type ILOAD_0 struct { base.NoOperandsInstruction }
type ILOAD_1 struct { base.NoOperandsInstruction }
type ILOAD_2 struct { base.NoOperandsInstruction }
type ILOAD_3 struct { base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
    val := frame.LocalVars().GetInt(index)
    frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
    _iload(frame, uint(self.index))
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
    _iload(frame, 1)
}
