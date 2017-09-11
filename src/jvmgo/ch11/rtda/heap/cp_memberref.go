package heap
import "jvmgo/ch11/classfile"
/*
    Chap6.2.2 
    字段符号引用1
    xinxin.shi
    2017-07-29 22:23:48
*/
type MemberRef struct {
    SymRef
    name    string
    descriptor  string
}

func (self *MemberRef) copyMemberRefInfo( 
    refInfo *classfile.ConstantMemberrefInfo) {
    self.className = refInfo.ClassName()
    self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
