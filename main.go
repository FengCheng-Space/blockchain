package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    bc = NewBlockchain() // 初始化区块链

    // 启动Web服务器
    go StartServer()

    // 创建一个通道来监听系统信号，以便在用户中断时优雅地退出
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    // 用户输入循环
    for {
        fmt.Println("Enter block data (or type 'exit' to quit):")
        var input string
        fmt.Scanln(&input)
        
        if input == "exit" {
            fmt.Println("Exiting...")
            break
        }
        
        bc.AddBlock(input)
        fmt.Println("Block added!")
        fmt.Println("Current blockchain:")
        bc.PrintBlocks()
    }

    // 等待系统信号
    <-signalChan
    fmt.Println("Server shutting down.")
}
// 全局变量
var bc *Blockchain
