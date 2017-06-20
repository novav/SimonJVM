package control

import "jvmgo/ch05/instructions/basr"
import "jvmgo/ch05/rtda"

/*
    5.10.1　goto指令
    xinxin.shi
    2017-06-20 23:39:14
*/
// Branch always

type GOTO struct { base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
    base.Branch(frame, self.Offset)
}

