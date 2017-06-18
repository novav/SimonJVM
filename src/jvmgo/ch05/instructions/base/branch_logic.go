/*
    5.9.3　if<cond>指令
    xinxin.shi
    2017-06-18 23:21:22
*/
package base
import "jvmgo/ch05/rtda"

func Branch(frame *rtda.Franme, offset int) {
    pc := frame.Thread().PC()
    nextPC := pc + offset
    frame.SetNextPc(nextPC)
}
