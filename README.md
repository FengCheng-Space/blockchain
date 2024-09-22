# Go 语言搭建建议的区块链通信框架

## 文件结构：
``` 
└─blockchain
    └─static
    |    └─index.html
    ├─blockchain.go
    ├─main.go
    ├─web.go
    └─README.md

```

## 运行：
```
$ go run .\main.go .\web.go .\blockchain.go
```

## 可视化显示：
#### 成功运行程序后，web浏览器将在localhost:8080端口上显示区块链信息。

其中包括：

|Index |	Timestamp |	Data |	Previous Hash |	Hash |
|:---:|:---:|:---:|:---:|:---:|
|0	| 2024-09-22 19:59:07.9881302 +0800 CST m=+0.008861501	|Genesis Block	| NULL|	72300bc396bff2ca0e26b764518ddb4240b0f7c0dfdfd28da97aac02ba917e43|
|:---:|:---:|:---:|:---:|:---:|


## Attention:
仅作测试实验，不保证安全性。


