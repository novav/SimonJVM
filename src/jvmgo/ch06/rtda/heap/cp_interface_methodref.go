package heap
import "jvmgo/ch06/classfile"
/*
Chap 6.2.4
接口Method符号引用
xinxin.shi
2017-07-29 22:41:07
*/

type InterfaceMethodRef {
    MemberRef
    method *Method 
}

func newInterfaceMethodRef(cp *ConstantPool,
    refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
    ref := &InterfaceMethodRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
    return ref
}