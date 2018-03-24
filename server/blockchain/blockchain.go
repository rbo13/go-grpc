package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block simple block
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain simple blockchain
type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// NewBlock creates new block in the block chain.
func NewBlock(data, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.setHash()
	return block
}

// AddBlock adds block to the blockchain
// using the previous block hash.
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// NewBlockchain creates new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock beginning of block
func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", "")
}
