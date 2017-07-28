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


