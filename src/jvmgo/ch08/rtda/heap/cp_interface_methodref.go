package heap
import "jvmgo/ch08/classfile"
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
/*
    7.2.2
    接口方法符号引用的解析
*/
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := self.ResolveClass()
	// todo
    d := self.cp.class
    c := self.ResolvedClass()
    if !c.IsInterface() {
        panic("java.lang.IncompatibleClassChangeError")
    }
    method := lookupInterfaceMethod(c, self.name, self.descriptor)
    if method == nil {
        panic("java.lang.NoSuchMethodError")
    }
    if !method.isAccessibleTo(d) {
        panic("java.lang.IllegalAccessError")
    }
    self.method = method
}


func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
    for _, method := range iface.methods {
        if method.name == name && method.descriptor == descriptor {
            return method
        }
    }
    return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}