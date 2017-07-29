package heap
import "jvmgo/ch06/classfile"
/*
    Chap 6.2.2
    字段符号引用2
    xinxin.shi
    2017-07-29 22:25:06
    */

type FieldRef struct {
    MemberRef
    field *Field
}

func newFieldRef (cp *ConstantPool, 
      refInfo *classfile.ConstantFieldrefInfo) {
    ref := &FieldRef{}
    ref.cp = cp
    ref.copyMemberRefInfo(&refInfo.COnstantMemberrefInfo)
    return ref
}    
