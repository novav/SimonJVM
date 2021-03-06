package heap

import "strings"
import "jvmgo/ch06/classfile"

/*
    Chap 6.1.1
    xinxin.shi
    2017-07-23 21:22:47
*/
type Class struct {
    accessFlags     uint16
    name            string // thisClassName
    superClassName string 
    interfaceNames  []string
    constantPool    *ConstantPool
    fields          []*Field
    methods         []*Method 
    loader          *ClassLoader
    superClass     *Class
    interfaces      []*Class 
    instanceSlotCount   uint 
    staticSlotCount     uint
    staticVars          Slots
}

// ClassFile -> Class 结构体

func newClass(cf *classfile.ClassFile) *Class {
    class := &Class{}
    class.accessFlags = cf.AccessFlags()
    class.name = cf.ClassName()
    class.superClassName = cf.SuperClassName()
    class.interfaceNames = cf.InterfaceNames()
    class.constantPool = newConstantPool(class, cf.ConstantPool()) // see 6.2小节
    class.fields = newFields(class, cf.Fields()) // see 6.1.2
    class.methods = newMethods(class, cf.Methods()) // 6.1.3
    return class
}

func (self *Class) IsPublic() bool {
    return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// jvms 5.4.4
func (self *Class) isAccessibleTo(other *Class) bool {
    return self.IsPublic() || 
        self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
    if i := strings.LastIndex(self.name, "/"); i >= 0 {
        return self.name[:i]
    }
    return ""
}


func (self *Class) GetMainMethod() *Method {
    return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
    for _, method := range self.methods {
        if method.IsStatic() && 
            method.name == name && 
            method.descriptor == descriptor {
                return method
        }
    }
    return nil
}
/* 6.6.1 */
func (self *Class) NewObject() *Object {
    return newObject(self)
}
