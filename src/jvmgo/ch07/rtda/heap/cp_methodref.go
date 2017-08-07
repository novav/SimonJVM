package heap
import "jvmgo/ch07/classfile"

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

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	//class := self.Class()
	// todo
}
