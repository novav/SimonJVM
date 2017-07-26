package heap

import "jvm/ch06/classfile"
/*
    chap 6.1.3
    Method info
    xinxin.shi
    2017-07-26 23:51:35
*/
type Method struct {
    ClassMember
    maxStack    uint
    maxLocals   uint 
    code        []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
    methods := make([]*Method, len(cfMethods))
    for i, cfMethod = range cfMethods {
        methods[i] = %Method{}
        methods[i].class = class
        methods[i].copyMemberInfo(cfMethod)
        methods[i].copyAttributes(cfMethod)
    }
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
    if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
        self.maxStack = codeAttr.MaxStack()
        self.maxLocals = codeAttr.maxLocals()
        self.code = codeAttr.Code()
    }
}