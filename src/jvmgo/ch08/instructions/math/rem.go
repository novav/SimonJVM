package math 
import "math"
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

/* 
    chap 5.7.1 rem 求余
    xinxin.shi
    2017-06-16 22:22:42 
*/
type DREM struct { base.NoOperandsInstruction }

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct { base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder int
type IREM struct { base.NoOperandsInstruction }

func (self *IREM) Execute (frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    if v2 == 0 {
        panic("java.lang.AritmeticException: / by zero")
    }
    result := v1 % v2
    stack.PushInt(result)
}

// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}