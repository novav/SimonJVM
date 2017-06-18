/*
    5.9.4　if_icmp<cond>指令
    xinxin.shi
    2017-06-18 23:32:24
*/
package comparisons
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"


// Branch if int comparison succeeds

type IF_ICMPEQ struct { base.BranchInstruction }
type IF_ICMPNE struct { base.BranchInstruction }
type IF_ICMPLT struct { base.BranchInstruction }
type IF_ICMPLE struct { base.BranchInstruction }
type IF_ICMPGT struct { base.BranchInstruction }
type IF_ICMPGE struct { base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    val2 := stack.PopInt()
    val1 := stack.PopInt()
    if val1 != val2 {
        base.Branch(frame, self.Offset)
    }
}
