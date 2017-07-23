package base

import "jvmgo/ch05/rtda"

type Instruction interface {
    FetchOperands(reader *BytecodeReader) // 提取操作数
    Execute(Frame *rtda.Frame) // 执行命令逻辑
}



type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction)FetchOperands(reader *BytecodeReader) {
    // nothing to do
}


type BranchInstruction struct { // 跳转指令
    Offset int
}

func (self *BranchInstruction)FetchOperands(reader *BytecodeReader) {
    self.Offset = int(reader.ReadInt16())
}


type Index8Instruction struct {
    Index uint
}

func (self *Index8Instruction)FetchOperands(reader *BytecodeReader) {
    self.Index = uint(reader.ReadUint8())
}


type Index16Instruction struct {
    Index uint
}

func (self *Index16Instruction)FetchOperands(reader *BytecodeReader) {
    self.Index = uint(reader.ReadUint16())
}