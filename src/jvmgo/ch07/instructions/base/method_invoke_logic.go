package base

/*
    Chap 7.3
    xinxin.shi
    2017-08-12 00:35:38
*/
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
    thread := invokerFrame.Thread()
    newFrame := thread.NewFrame(method)
    thread.PushFrame(newFrame)
    argSlotSlot := int(method.ArgSlotCount())
    if argSlotSlot > 0 {
        for i := argSlotSlot - 1 ; i >= 0; i-- {
            slot := invokerFrame.OperandStack().PopSlot()
            newFrame.LocalVars().SetSlot(uint(i), slot)
        }
    }
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
    self[index] = slot
}

