package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一块区块链的hash值
	Hash          string //当前区块的hash值
	Data          string //区块数据
}

func calculateHash(b *Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInByte := sha256.Sum256([]byte(blockData))
	hashInStr := hex.EncodeToString(hashInByte[:])

	return hashInStr
}

func GenerteNewBolck(preBlock *Block, data string) *Block {
	newBlock := new(Block)
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

//原始区块
func GenerteGenesisBlock() *Block {
	newBlock := new(Block)
	newBlock.Index = - 1
	newBlock.Hash = ""

	return GenerteNewBolck(newBlock, "Genesis Block")
}
