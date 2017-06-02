package classfile

/*  xinxin.shi
    2017-06-02
*/

type MemberInfo struct {
    cp                  ConstantPool    // 常量池指针
    accessFlags         uint16          // 访问标志
    nameIndex           uint16          // 常量池索引--字段/方法名
    descriptorIndex     uint16          // 常量池索引--字段/方法描述符
    attributes          []attributes    // 属性表
}

func readMembers(reader *ClassReader, cp ConstantPool) *MemberInfo {
    memberCount := reader.readUint16()
    members := make([]*MemberInfo, memberCount)
    for i := range members {
        members[i] = readMember(reader, cp)
    }
    return members
}

func readMember (reader *ClassReader, cp ConstantPool) []*MemberInfo {
    return &MemberInfo{
        cp:                 cp,
        accessFlags:        reader.readUint16(),
        nameIndex:          reader.readUint16(),
        descriptorIndex:    reader.readUint16(),
        attributes:         readAttributes(reader, cp), //see: 3.4
    }
}

func (self *MemberInfo)AccessFlags() uint16 {
    
}

func (self *MemberInfo) Name() string {
    return self.cp.getUtf8(self.nameIndex)    
}


func (self *MemberInfo) Descriptor() string {
    return self.cp.getUtf8(self.descriptorIndex)
}

