package constants 
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
    5.3.3
    xinxin.shi
    2017-06-13 23:39:17
*/
type BIPUSH struct { val int8} // push byte
type SIPUSH struct { cal int16} // Push short

func (self *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
    self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
    i := int32(self.val)
    frame.OperandStack().PushInt(i)
}

