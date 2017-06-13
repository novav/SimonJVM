package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
    5.3.2
    xinxin.shi
    2017-06-13 23:31:38
*/

type ACONST_NULL struct { base.NoOperandsInstruction }
type DCONST_0 struct { base.NoOperandsInstruction }
type DCONST_1 struct { base.NoOperandsInstruction }
type FCONST_0 struct { base.NoOperandsInstruction }
type FCONST_1 struct { base.NoOperandsInstruction }
type FCONST_2 struct { base.NoOperandsInstruction }
type ICONST_M1 struct { base.NoOperandsInstruction }
type ICONST_0 struct { base.NoOperandsInstruction }
type ICONST_1 struct { base.NoOperandsInstruction }
type ICONST_2 struct { base.NoOperandsInstruction }
type ICONST_3 struct { base.NoOperandsInstruction }
type ICONST_4 struct { base.NoOperandsInstruction }
type ICONST_5 struct { base.NoOperandsInstruction }
type LCONST_0 struct { base.NoOperandsInstruction }
type LCONST_1 struct { base.NoOperandsInstruction }

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
    frame.OperandStack().PushRef(nil)
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
    frame.OperandStack().PushDouble(0.0)
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
    frame.OperandStack().PushInt(-1)
}
