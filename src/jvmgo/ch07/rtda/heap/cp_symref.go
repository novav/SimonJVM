package heap

/*
    6.2.1
    类符号引用
    xinxin.shi
    2017-07-28 23:43:04
*/
// symbolic reference
type SymRef struct {
    cp          *ConstantPool
    className   string
    class       *Class
}

func (self *SymRef) ResolvedClass() *Class {
    if self.class == nil {
        self.resolveClassRef() 
    }
    return self.class
}

func (self *SymRef) resolveClassRef() {
    d := self.cp.class
    c := d.loader.LoadClass(self.className)
    if !c.isAccessibleTo(d) {
        panic("java.lang.IllegalAccessError")
    }
    self.class = c
}
