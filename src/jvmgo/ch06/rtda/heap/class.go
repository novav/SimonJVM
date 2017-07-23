package heap

import "jvmgo/ch06/classfile"

/*
    Chap 6.1.1
    xinxin.shi
    2017-07-23 21:22:47
*/
type Class struct {
    accessFlags     uint16
    name            string // thisClassName
    supperClassName string 
    interfaceNames  []string
    constantPool    *ConstantPool
    fields          []*Field
    methods         []*Method 
    loader          *ClassLoader
    supperClass     *Class
    interfaces      []*Class 
    instanceSlotCount   uint 
    staticSlotCount     uint
    staticVars          *Slot 
}

// ClassFile -> Class 结构体

func newClass(cf *classfile.ClassFile) *Class {
    class := &Class{}
    class.accessFlags = cf.AccessFlags()
    class.name = cf.ClassName()
    class.supperClassName = cf.SupperClassName()
    class.interfaceNames = cf.InterfaceNames()
    class.constantPool = newConstantPool(class, cf.ConstantPool()) // see 6.2小节
    class.fields = newFields(class, cf.Field()) // see 6.1.2
    class.methods = newMathods(class, cf.Methods()) // 6.1.3
    return class
}


func (self *Class) IsPublic() bool {
    return 0 != self.accessFlags&ACC_PUBLIC
}
