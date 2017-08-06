package main

import "fmt"
import "strings"
// import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"
import "jvmgo/ch06/rtda/heap"

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
        cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
        classLoader := heap.NewClassLoader(cp)
        className := strings.Replace(cmd.class, ".", "/", -1)
        mainClass := classLoader.LoadClass(className)  // class.GetMainClass(className)
        // cf := loadClass(className, cp)
        mainMethod := mainClass.GetMainMethod()
        if mainMethod != nil {
            interpret(mainMethod)
        } else {
            fmt.Printf("Main method not found in class %s\n", cmd.class)
        }
    }