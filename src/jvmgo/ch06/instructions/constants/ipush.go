package constants 
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

/*
    5.3.3
    xinxin.shi
    2017-06-13 23:39:17
*/
// Push byte
type BIPUSH struct { 
    val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
    i := int32(self.val)
    frame.OperandStack().PushInt(i)
}

// Push short
type SIPUSH struct { 
    val int16
}
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
