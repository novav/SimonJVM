package math
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

/*
    5.7.4　iinc指令
    xinxin.shi    
    2017-06-16 23:17:57
*/
// Increment local variable by constant 
type IINC struct {
    Index uint 
    Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
    self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
    localVars := frame.LocalVars()
    val := localVars.GetInt(self.Index)
    val += self.Const
    localVars.SetInt(self.Index, val)
}

