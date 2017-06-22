/*
    5.11.3　goto_w指令
    xinxin.shi
    2017-06-22 23:22:51
*/
package extended 
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch always (wide index)
type GOTO_W struct {
    offset int
}

func (self *GOTO_W) FectchOperands(reader *base.BytecodeReader) {
    self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtad.Frame) {
    base.Branch(frame, self.offset)
}