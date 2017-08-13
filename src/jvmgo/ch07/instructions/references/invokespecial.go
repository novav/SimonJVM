package references 

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type INVOKE_SPECIAL struct { base.Index16Instruction }

// hack!

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
    // frame.OperandStack().PopRef()
    currentClass := frame.Method().Class()
    cp := currentClass.ConstantPool()
    methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
    resolvedClass := methodRef.ResolvedClass()
    resolvedMethod := methodRef.ResolvedMethod()
    if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
        panic("java.lang.NoSuchMethodError")
    }
    if resolvedMethod.IsStatic() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
    if ref == nil {
        panic("java.lang.NullPointException")
    }
    if resolvedMethod.IsProtected() &&
        resolvedMethod.IsSuperClassOf(currentClass) &&
        resolvedMethod.GetPackageName() != currentClass.GetPackageName() &&
        ref.Class() != currentClass &&
        !ref.Class().IsSubClassOf(currentClass) {
            panic("java.lang.IllegalAccessError")
    }

    methodBeInvoked := resolvedMethod
    if currentClass.IsSuper() && 
        resolvedClass.IsSuperClassOf(currentClass) && 
        resolvedMethod.Name() != "<init>" {
            methdoToInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), 
                methodRef.Name(), methodRef.Descriptor())
    }

    if methdoToInvoked == nil || methdoToInvoked.IsAbstract() {
        panic("java.lang.AbstractMethodError")
    }
    base.InvokeMethod(frame, methdoToInvoked)

    
}