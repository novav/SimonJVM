/*
    5.9.2　fcmp<op>和dcmp<op>指令
    xinxin.shi 
    2017-06-18 22:58:48
*/
package comparisons 
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
// Compare float    
type FCMPG struct { base.NoOperandsInstruction }
type FCMPL struct { base.NoOperandsInstruction }

func _fcmp(frame *rtda.Frame, gFlag bool) {
    stack := frame.OperandStack()
    v2 := stack.PopFloat()
    v1 := stack.PopFloat()
    if v1 > v2 {
        stack.PushInt(1)       
    } else if v1 == v2 {
        stack.PushInt(0)
    } else if gFlag {
        stack.PushInt(1)
    } else {
        stack.PushInt(-1)
    }
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
    _fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
    _fcmp(frame, false)
}