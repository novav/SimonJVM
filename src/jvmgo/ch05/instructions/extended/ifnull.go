/*
    5.11.2　ifnull和ifnonnull指令
    xinxin.shi
    2017-06-22 23:13:37
*/

package extended 
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IFNULL struct { base.BranchInstruction } // Branch if reference is null 
type IFNONNULL struct { base.BranchInstruction } // Branch if reference not null

func (self *IFNULL) Execute(frame *rtda.Frame) {
    ref := frame.OperandStack().PopRef()
    if ref == nil {
        base.Branch(frame, self.Offset)
    }
}