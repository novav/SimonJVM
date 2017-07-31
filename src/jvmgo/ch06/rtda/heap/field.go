package heap

import "jvmgo/ch06/classfile"
/*
    6.1.2
    字段
    2017-07-25 23:26:14
*/
type Field struct {
    ClassMember
    constValueIndex uint
    slotId  uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field{
    fields := make([]*Field, len(cfFields))
    for i, cfField := range cfFields {
        fields[i] = &Field{}
        fields[i].class = class
        fields[i].copyMemberInfo(cfField)
        fields[i].copyAttributes(cfField)
    }
    return fields
}

/*6.4*/
func (self *Field) isLongOrDouble() bool {
    return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
    if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
        self.constValueIndex = uint(valAttr.ConstantValueIndex())
    }
}