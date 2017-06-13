package constants 

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Do nathing

/*
    Chap 5.3.1
    xinxin.shi
    2017-06-12 21:28:09
*/

type NOP struct {
    base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
    // 什么也不用
}