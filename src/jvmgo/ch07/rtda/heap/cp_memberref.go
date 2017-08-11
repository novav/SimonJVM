package heap
import "jvmgo/ch07/classfile"
/*
    Chap6.2.2 
    字段符号引用1
    xinxin.shi
    2017-07-29 22:23:48
*/
type MemberRef struct {
    SymRef
    name    string
    descriptor  string
}

func (self *MemberRef) copyMemberRefInfo( 
    refInfo *classfile.ConstantMemberrefInfo) {
    self.className = refInfo.ClassName()
    self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}

/*7.2.1*/
func (self *MemberRef) ResolvedMethod() *Method {
    if self.method == nil {
        self.resolvedMethodRef()
    }
    return self.method
}

func (self *MemberRef) resolvedMethodRef() {
    d := self.cp.class
    c := self.ResolvedClass()
    if c.IsInterface() {
        panic("java.lang.IncomptiableClassChangeError");
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
    method := lookupMethodInClass(class, name, descriptor)
    if method == nil {
        method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
    }
    return method
}

func lookupMethodInClass(class *Class, name, descriptor string) *Method {
    for c := class; c != nil ; c = c.superClass {
        for _, method := range c.methods {
            if method.name == name && method.descriptor == descriptor {
                return method
            }
        }
    }
    return nil
}

func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
    for _, iface := range ifaces {
        for _, method := range ifaces.methods {
            if method.name == name && method.descriptor == descriptor {
                return method;
            }
        }
        method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
        if method != nil {
            return method
        }
    }
    return nil
}