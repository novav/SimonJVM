/*
    5.10.3　lookupswitch指令
    xinxin.shi    
    2017-06-21 00:01:38
*/
package control 

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type LOOKUP_SWITCH struct {
    defaultOffset   int32
    npairs          int32
    matchOffsets    []int32
}

func (self *LOOKUP_SWITCH) FetchOperands (reader *base.BytecodeReader) {
    reader.SkipPadding()
    self.defaultOffset = reader.ReadInt32()
    self.npairs = reader.ReadInt32()
    self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
    key := frame.OperandStack().PopInt()
    for i := int32(0); i < self.npairs * 2 ; i += 2{
        offset := self.matchOffsets[i+1]
        base.Branch(frame, int(offset))
        return 
    }
    base.Branch(frame, int(self.defaultOffset))
}