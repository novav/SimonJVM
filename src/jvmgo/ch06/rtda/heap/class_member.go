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

