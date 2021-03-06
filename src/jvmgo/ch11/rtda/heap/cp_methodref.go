package heap
import "jvmgo/ch11/classfile"

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
    d := self.cp.class
    c := self.ResolvedClass()
    if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
    }

    method := lookupMethod(c, self.name, self.descriptor)
    if method == nil {
        panic("java.lang.NoSuchMethodError")
    }
    if !method.isAccessibleTo(d) {
        panic("java.lang.IllegalAccessError")
    }
    self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
