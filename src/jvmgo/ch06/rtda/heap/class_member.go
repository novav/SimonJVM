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
    self.name   = memberInfo.Name()
    self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
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