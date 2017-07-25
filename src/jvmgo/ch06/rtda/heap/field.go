package heap

import "jvmgo/ch06/classfile"
/*
    6.1.2
    字段
    2017-07-25 23:26:14
*/
type Field struct {
    ClassMember
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field{
    fields := make([]*Field, len(cfFields))
    for i, cfField := range cfFields {
        fields[i] = &Field{}
        fields[i].class = class
        fields[i].copyMemberInfo(cfField)
    }
    return fields
}