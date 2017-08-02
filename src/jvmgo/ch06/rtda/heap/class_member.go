package heap

import "jvmgo/ch06/classfile"

/*
    Chap 6.1.2
    字段信息 ： 字段 方法
*/
type ClassMember struct {
    accessFlags     uint16
    name            string
    descriptor      string
    class           *Class // 存放class结构体指针
}

func (self *ClassMember) copyMemberInfo( memberInfo *classfile.MemberInfo) {
    self.accessFlags = memberInfo.AccessFlags()
    self.name   = memberInfo.Nmae()
    self.descriptor = memberInfo.Descriptor()
}

/* 6.5.2 */
func (self *ClassMember) isAccessibleTo(d *Class) bool {
    if self.IsPublic() {
        return true
    }
    c := self.class 
    if self.IsProtected() {
        return d == c || d.isSubClassOf(c) ||
            c.getPackageName() == d.getPackageName()
    }

    if !self.IsPrivate() {
        return c.getPackageName() == d.getPackageName()
    }

    return d == c
}