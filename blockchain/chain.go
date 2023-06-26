package blockchain

import (
	"sync"

	"github.com/cumedang/go/db"
	"github.com/cumedang/go/utils"
)

type blockchain struct {
	NewestHash       string `json:"newesthash"`
	Height           int    `json:"height"`
	CurrenDifficulty int    `json:"currendifficulty"`
}

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func persistBlockchain(b *blockchain) {
	db.SaveCheckPoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1, getdifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrenDifficulty = block.Difficulty
	persistBlockchain(b)
}

func Txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Transaction...)
	}
	return txs
}

func Findtx(b *blockchain, targetID string) *Tx {
	for _, tx := range Txs(b) {
		if tx.Id == targetID {
			return tx
		}
	}
	return nil
}

func recalculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
	newestBlock := allBlocks[0]
	lastRecalculateBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.TimeStamp / 60) - (lastRecalculateBlock.TimeStamp / 60)
	expectedtime := difficultyInterval * blockInterval
	if actualTime <= (expectedtime - allowedRange) {
		return b.CurrenDifficulty + 1
	} else if actualTime >= (expectedtime + allowedRange) {
		return b.CurrenDifficulty - 1
	} else {
		return b.CurrenDifficulty
	}
}

func getdifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return recalculateDifficulty(b)
	} else {
		return b.CurrenDifficulty
	}
}

func Blocks(b *blockchain) []*Block {
	var blocks []*Block
	hashCusor := b.NewestHash
	for {
		block, _ := FindBlock(hashCusor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCusor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)
	for _, block := range Blocks(b) {
		for _, tx := range block.Transaction {
			for _, input := range tx.TxIns {
				if input.Signature == "COINBASE" {
					break
				}
				if Findtx(b, input.TxID).TxOuts[input.Index].Address == address {
					creatorTxs[input.TxID] = true
				}
			}
			for index, output := range tx.TxOuts {
				if output.Address == address {
					if _, ok := creatorTxs[tx.Id]; !ok {
						uTxOut := &UTxOut{tx.Id, index, output.Amount}
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}

					}
				}

			}
		}
	}
	return uTxOuts
}

func BalanceByAddress(address string, b *blockchain) int {
	txOuts := UTxOutsByAddress(address, b)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}

		checkpoint := db.Checkpoint()
		if checkpoint == nil {
			b.AddBlock()
		} else {

			b.restore(checkpoint)
		}

	})
	return b
}
