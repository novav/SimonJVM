package stores

import "jvmgo/ch05/instructions/base/"
import "jvmgo/ch05/rtda"


/*
    chap 5.6 存储指令
    xinxin.shi
    2017-06-14 23:09:44
*/

// Storae long into local variable

type LSTORE struct { base.Index8Instruction }
type LSTORE_0 struct { base.NoOperandsInstruction }
type LSTORE_1 struct { base.NoOperandsInstruction }
type LSTORE_2 struct { base.NoOperandsInstruction }
type LSTORE_3 struct { base.NoOperandsInstruction }


func _lstore(frame *rtda.Frame, index uint) {
    val := frame.OperandStack().PopLong()
    frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE)Execute(frame *rtda.Frame) {
    _lstore(frame, uint(self.Index))
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
    _lstore(frame, 2)
}