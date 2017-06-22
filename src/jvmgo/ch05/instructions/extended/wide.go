
/*
    5.11.1　wide指令
    xinxin.shi
    2017-06-22 23:12:11
*/
package extended 

import "jvmgo/ch05/instrctions/base"
import "jvmgo/ch05/instrctions/loads"
import "jvmgo/ch05/instrctions/math"
import "jvmgo/ch05/instrctions/stores"
import "jvmgo/ch05/rtda"

type WIDE struct {
    modifiedInstruction base.Instruction
}

func (self *WIDE) FectchUint8(reader *base.BytecodeReader) {
    opcode := reader.ReaderUint8()
    switch opcode {
        case 0x15: // iload
            inst := &loads.ILOAD{}
            inst.Index = uint(reader.ReaderUint16())
            self.modifiedInstruction = inst
        case 0x16: ... // lload
        case 0x17: ... // fload
        case 0x18: ... // dload
        case 0x19: ... // aload
        case 0x36: ... // istore
        case 0x37: ... // lstore
        case 0x38: ... // fstore
        case 0x39: ... // dstore
        case 0x3a: ... // astore
        case 0x84: ... // iinc
            inst := &math.IINC{}
            inst.Index = uint(reader.ReaderUint16())
            inst.Const = int32(reader.ReaderUint16())
            self.modifiedInstruction = inst
        case 0xa9: // ret
            panic("Unsupported opcode 0xa9!")
    }
}

func (self *WIDE) Execute(frame *rtda.Frame) {
    self.modifiedInstruction.Execute(frame)
}