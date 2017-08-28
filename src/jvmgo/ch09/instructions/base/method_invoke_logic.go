package base
// import "fmt"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

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

    // 前面的代码不变，下面是 hack
    // if method.IsNative() {
    //     if method.Name() == "registerNatives" {
    //         thread.PopFrame()
    //     } else {
    //         panic(fmt.Sprintf("native method: %v.%v%v\n",
    //             method.Class().Name(), method.Name(), method.Descriptor()))
    //     }
    // }
}

