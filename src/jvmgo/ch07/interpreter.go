package main

import "fmt"
// import "jvmgo/ch07/classfile"
import "jvmgo/ch07/instructions"
import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

// func interpret(methodInfo *classfile.MemberInfo) {
    // codeAttr := methodInfo.CodeAttribute()
    // maxLocals := codeAttr.MaxLocals()
    // maxStack := codeAttr.MaxStack()
    // bytecode := codeAttr.Code()
func interpret(method *heap.Method, logInst bool) {
    thread := rtda.NewThread()
    frame := thread.NewFrame(method)
    thread.PushFrame(frame)
    defer catchErr(frame)
    loop(thread, logInst)
}


func catchErr(frame *rtda.Frame) {
    if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
        panic(r)
    }
}

func loop(thread *rtda.Thread, logInst bool) {
    reader := &base.BytecodeReader{}
    for {
	    frame := thread.PopFrame()
        pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(frame.Mehtod().Code(), pc)
        opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
        inst.FetchOperands(reader)
        frame.SetNextPC(reader.PC())
        if (logInst) {
            logInstruction(frame, inst)
        }
        // execute 
        fmt.Printf("pc:%2d inst: %T %v \n", pc, inst, inst)
        inst.Execute(frame)
        if thread.IsStackEmpty() {
            break
        }
    }
}


func catchErr(thread *rtda.Thread) {
    if r := recover(); r != nil {
        logFrames(thread)
        panic(r)
    }
}

func logFrames(thread *rtda.Thread) {
    for !thread.IsStackEmpty() {
        frame := thread.PopFrame()
        method := frame.Method()
        className := method.Class().Name()
        fmt.Printf(">> pc: %4d %v.%v%v \n", 
            frame.NextPC(), className, method.Name(), method.Descriptor())
    }
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
    method := frame.Method()
    className := method.Class().Name()
    methodName := method.Name()
    pc := frame.Thread().PC()
    fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}