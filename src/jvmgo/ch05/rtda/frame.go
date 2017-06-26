package rtda 

// stack frame
type Frame struct {
	lower           *Frame // stack is implemented as linked list
    localVars       LocalVars
    operandStack    *OperandStack
	// todo
    thread          *Thread
    nextPC          int
}

func NewFrame(thread *Therad, maxLocals, maxStack uint) *Frame {
    return &Frame {
        thread:         thread,
        localVars:      newLocalVars(maxLocals),
        operandStack:   newOperandStack(maxStack),
    }
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
