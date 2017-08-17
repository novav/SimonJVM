package references 

import "fmt"
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Invoke instance method; dispatch based on class
/*
    Chap 6 & 7.5.3
    invokevirtual指令
    xinxin.shi
    2017-08-13 16:31:33
*/

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

// hack!

func  (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
    currentClass := frame.Method().Class()
    // cp := frame.Method().Class().ConstantPool()
    cp := currentClass.ConstantPool()
    methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
    resolvedMethod := methodRef.ResolvedMethod()
    if resolvedMethod.IsStatic() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)
    if ref == nil {
        // hack!
        if methodRef.Name() == "println" {
            _println(frame.OperandStack(), methodRef.Descriptor())
            return
        }
		panic("java.lang.NullPointerException")
    }

    if resolvedMethod.IsProtected() &&
        resolvedMethod.Class().IsSuperClassOf(currentClass) &&
        resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
        ref.Class() != currentClass &&
        !ref.Class().IsSubClassOf(currentClass) {
            panic("java.lang.IllegalAccessError")
    }

    methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), 
            methodRef.Name(), methodRef.Descriptor())
    if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
        panic("java.lang.AbstractMethodError")
    }
    base.InvokeMethod(frame, methodToBeInvoked)


}
func _println(stack *rtda.OperandStack, descriptor string) {
        switch descriptor {
            case "(Z)V": 
                fmt.Printf("%v\n", stack.PopInt() != 0)
            case "(C)V": 
                fmt.Printf("%c\n", stack.PopInt())
            case "(I)V", "(B)V", "(S)V": 
                fmt.Printf("%v\n", stack.PopInt())
            case "(J)V": 
                fmt.Printf("%v\n", stack.PopLong())
            case "(F)V": 
                fmt.Printf("%v\n", stack.PopFloat())
            case "(D)V": 
                fmt.Printf("%v\n", stack.PopDouble())
            default : 
                panic("println : " + descriptor)
        }
        stack.PopRef()
}