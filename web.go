package main

import (
    "encoding/json"
    "net/http"
)

// handleBlockchain 返回区块链的JSON数据
func handleBlockchain(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(bc.blocks) // 将区块链数据编码为JSON格式并返回
}

// serveStaticFile 提供静态文件（index.html）
func serveStaticFile(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
}

// StartServer 启动Web服务器
func StartServer() {
    // 处理区块链数据的请求
    http.HandleFunc("/blockchain", handleBlockchain)

    // 处理静态页面的请求
    http.HandleFunc("/", serveStaticFile)

    // 启动服务器，监听在8080端口
    http.ListenAndServe(":8080", nil)
}
