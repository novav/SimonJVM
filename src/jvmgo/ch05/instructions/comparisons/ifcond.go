/*
    5.9.3　if<cond>指令
    xinxin.shi
    2017-06-18 23:08:30
*/
package comparisons
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch if int comparison with zero succeeds
type IFEQ struct { base.BranchInstruction }
type IFNE struct { base.BranchInstruction }    
type IFLT struct { base.BranchInstruction }    
type IFLE struct { base.BranchInstruction }    
type IFGT struct { base.BranchInstruction }    
type IFGE struct { base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
    val := frame.OperandStack().PopInt()
    if val == 0 {
        base.Branch(farme, self.Offset)
    }
}

