package main

import "fmt"
import "strings"
import "jvmgo/ch04/classfile"
import "jvmgo/ch04/classpath"
import "jvmgo/ch04/rtda"

    /*
    xinxin.shi
    2017-06-04 23:45:47
    */
    /*
    介绍了Java虚拟机从哪里搜索class文件，并且实现了类路径功能，已经可以把class文件读取到内存中
    */

    func main() {
        cmd := parseCmd()

        if cmd.versionFlag {
            fmt.Println("version 0.0.1")
        } else if cmd.helpFlag || cmd.class == "" {
            printUsage()
        } else {
            startJVM(cmd)
        }
    }
    
    func startJVM(cmd *Cmd) {
        frame := rtda.NewFrame(100, 100)
        testLocalVars(frame.LocalVars())
        testOperandStack(frame.OperandStack())
    }
    
    func testLocalVars(vars rtda.LocalVars) {
        vars.SetInt(0, 100)
        vars.SetInt(1, -100)
        vars.SetLong(2, 2997924580)
        vars.SetLong(4, -2997924580)
        vars.SetFloat(6, 3.1415926)
        vars.SetDouble(7, 2.71828182845)
        vars.SetRef(9, nil)
        println(vars.GetInt(0))
        println(vars.GetInt(1))
        println(vars.GetLong(2))
        println(vars.GetLong(4))
        println(vars.GetFloat(6))
        println(vars.GetDouble(7))
        println(vars.GetRef(9))
    }

    func testOperandStack(ops *rtda.OperandStack) {
        ops.PushInt(100)
        ops.PushInt(-100)
        ops.PushLong(2997924580)
        ops.PushLong(-2997924580)
        ops.PushFloat(3.1415926)
        ops.PushDouble(2.71828182845)
        ops.PushRef(nil)
        println(ops.PopRef())
        println(ops.PopDouble())
        println(ops.PopFloat())
        println(ops.PopLong())
        println(ops.PopLong())
        println(ops.PopInt())
        println(ops.PopInt())
    }


    func startJVMFFFF(cmd *Cmd) {
        // ch02.exe -Xjre D:\Tools\Java\jdk1.8.0_77\jre\     java.lang.Object
        //                  cmd.XjreOption                    cmd.cpOption
        fmt.Printf("main->startJVM() => cmd.XjreOption:%v, cmd.cpOption:%v \n", cmd.XjreOption, cmd.cpOption)
        // cmd.XjreOption:D:\Tools\Java\jdk1.8.0_77\jre\, cmd.cpOption:cmd.XjreOption:D:\Tools\Java\jdk1.8.0_77\jre\, cmd.cpOption:

        cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

        fmt.Printf("main->startJVM() => classpath:%v class:%v args:%v \n", cp, cmd.class, cmd.args)
        // classpath:D:\Tools\go\workspace\bin class:java.lang.Object args:[]

        className := strings.Replace(cmd.class, ".", "/", -1)

        cf := loadClass(className, cp)
        fmt.Println(cmd.class)
        printClassInfo(cf)

    }

    func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
        classData, _, err := cp.ReadClass(className)
        if err != nil {
            panic(err)
        }
        cf, err := classfile.Parse(classData)
        if err != nil {
            panic(err)
        }
        return cf
    }

    func printClassInfo(cf *classfile.ClassFile) {
        fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
        fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
        fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
        fmt.Printf("this class: %v\n", cf.ClassName())
        fmt.Printf("super class: %v\n", cf.SuperClassName())
        fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
        fmt.Printf("fields count: %v\n", len(cf.Fields()))
        for _, f := range cf.Fields() {
            fmt.Printf("    %s\n", f.Name())
        }
        
        fmt.Printf("methods count: %v\n", len(cf.Methods()))
        for _, m := range cf.Methods() {
            fmt.Printf("    %s\n", m.Name())
        }
    }