package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cumedang/go/db"
	"github.com/cumedang/go/utils"
)

type Block struct {
	Hash        string `json:"hash"`
	PrevHash    string `json:"prevHash,omitempty"`
	Height      int    `json:"height"`
	Difficulty  int    `json:"difficulty"`
	Nonce       int    `json:"nonce"`
	TimeStamp   int    `json:"timestamp"`
	Transaction []*Tx  `json:"ransaction"`
}

const (
	minerReward int = 50
)

func (b *Block) persist() {

	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("블록을 찾을수 없습니당")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil

}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.TimeStamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int, diff int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: diff,
		Nonce:      0,
	}
	block.mine()
	block.Transaction = Mempool.TxToConfirm()
	block.persist()
	return block
}
