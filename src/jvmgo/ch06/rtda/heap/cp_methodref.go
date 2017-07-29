package heap
import "jvmgo/ch06/classfile"

/*
Chap 6.2.3
方法符号引用
xinxin.shi
2017-07-29 22:33:21
*/
type MethodRef struct {
    MemberRef
    method *Method
}

func newMethodRef(cp *ConstantPool,
    refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
    ref := &MethodRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
    return ref
}