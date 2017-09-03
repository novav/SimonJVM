package heap
import "jvmgo/ch09/classfile"

/*
    6.2.1 
    类符号引用
    xinxin.shi
    2017-07-28 23:42:30
*/

type ClassRef struct {
    SymRef
}

func newClassRef(cp *ConstantPool,
       classInfo *classfile.ConstantClassInfo) *ClassRef {
    ref := &ClassRef{}
    ref.cp = cp
    ref.className = classInfo.Name()
    return ref
}