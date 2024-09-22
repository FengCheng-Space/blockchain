package main

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "time"
)

// Block 结构体表示区块 
type Block struct {
    Index        int
    Timestamp    string
    Data         string
    PrevHash     string
    Hash         string
    Nonce        int
}

// CalculateHash 用于计算区块的哈希值
func CalculateHash(index int, timestamp string, data string, prevHash string, nonce int) string {
    record := strconv.Itoa(index) + timestamp + data + prevHash + strconv.Itoa(nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hash := h.Sum(nil)
    return hex.EncodeToString(hash)
}

// NewBlock 创建一个新区块 
func NewBlock(index int, prevHash string, data string) *Block {
    block := &Block{
        Index:     index,
        Timestamp: time.Now().String(),
        Data:      data,
        PrevHash:  prevHash,
        Hash:      "",
        Nonce:     0,
    }
    block.Hash = CalculateHash(block.Index, block.Timestamp, block.Data, block.PrevHash, block.Nonce)
    return block
}

// Blockchain 结构体表示区块链
type Blockchain struct {
    blocks []*Block
}

// NewBlockchain 创建一个新的区块链
func NewBlockchain() *Blockchain {
    return &Blockchain{blocks: []*Block{NewBlock(0, "", "Genesis Block")}}
}

// AddBlock 向区块链中添加一个新区块
func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks)-1]
    newBlock := NewBlock(len(bc.blocks), prevBlock.Hash, data)
    bc.blocks = append(bc.blocks, newBlock)
}

// PrintBlocks 打印区块链中的所有区块
func (bc *Blockchain) PrintBlocks() {
    for _, block := range bc.blocks {
        println("Index:", block.Index)
        println("Timestamp:", block.Timestamp)
        println("Data:", block.Data)
        println("PrevHash:", block.PrevHash)
        println("Hash:", block.Hash)
        println()
    }
}
