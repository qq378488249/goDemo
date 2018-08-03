package core

import (
	"log"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.Blocks = append(bc.Blocks, GenerteGenesisBlock())
	return bc
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.GetLastBlock()
	bc.Blocks = append(bc.Blocks, GenerteNewBolck(preBlock, data))
}

//获取最后一个区块
func (bc *BlockChain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

//添加区块
func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
	}
	lastBlock := bc.GetLastBlock()
	if isValid(newBlock, lastBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Println("invalid block")
		return
	}
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Println("index:", block.Index)
		fmt.Println("hash:", block.Hash)
		fmt.Println("PrevBlockHash:", block.PrevBlockHash)
		fmt.Println("data:", block.Data)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println()
	}
}

func isValid(newBlock, oldBlack *Block) bool {
	if newBlock.Index-1 != oldBlack.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlack.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
