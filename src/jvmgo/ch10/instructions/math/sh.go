package math 
import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

/*
    chap 5.7.2 位移指令
    xinxin.shi
    2017-06-16 22:36:57
*/

type ISHL struct { base.NoOperandsInstruction } // int 左移位

// type ISHR struct { base.NoOperandsInstruction } // int 算术右移

// type IUSHR struct { base.NoOperandsInstruction } // int 逻辑右位移

// type LSHL struct { base.NoOperandsInstruction }  // long 左位移

// type LSHR struct { base.NoOperandsInstruction }  // long 算术右位移

// type LUSHR struct { base.NoOperandsInstruction } // long 逻辑右位移


func (self *ISHL) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    s := uint32(v2) & 0x1f  // 5个bit 32位
    result := v1 << s
    stack.PushInt(result)
}

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
	v1 := stack.PopInt()
    s := uint32(v2) & 0x1f  // 64位 取前6个bits
    result := v1 >> s
	stack.PushInt(result)
}

// Logical shift right int
type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := int32(uint32(v1) >> s)
    stack.PushInt(result)
}

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
