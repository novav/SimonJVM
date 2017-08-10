package classfile

/*  
    3.2.8 MemberInfo
    xinxin.shi
    2017-06-02

field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
    cp                  ConstantPool    // 常量池指针
    accessFlags         uint16          // 访问标志
    nameIndex           uint16          // 常量池索引--字段/方法名
    descriptorIndex     uint16          // 常量池索引--字段/方法描述符
    attributes          []AttributeInfo    // 属性表
}

// read field or method table
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
    memberCount := reader.readUint16()
    members := make([]*MemberInfo, memberCount)
    for i := range members {
        members[i] = readMember(reader, cp)
    }
    return members
}

func readMember (reader *ClassReader, cp ConstantPool) *MemberInfo {
    return &MemberInfo{
        cp:                 cp,
        accessFlags:        reader.readUint16(),
        nameIndex:          reader.readUint16(),
        descriptorIndex:    reader.readUint16(),
        attributes:         readAttributes(reader, cp), //see: 3.4
    }
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
    return self.cp.getUtf8(self.nameIndex)    
}


func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
    for _, attrinfo := range self.attributes {
        switch attrinfo.(type) {
        case *ConstantValueAttribute:
            return attrinfo.(*ConstantValueAttribute)
        }
    }
    return nil
}