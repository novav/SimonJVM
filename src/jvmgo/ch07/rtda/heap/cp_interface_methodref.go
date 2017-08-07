package heap
import "jvmgo/ch07/classfile"
/*
Chap 6.2.4
接口Method符号引用
xinxin.shi
2017-07-29 22:41:07
*/

type InterfaceMethodRef struct {
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

func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.4
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := self.ResolveClass()
	// todo
}
