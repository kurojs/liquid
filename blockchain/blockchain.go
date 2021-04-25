package blockchain

import (
	"errors"
	"fmt"
	"time"
)

type Block struct {
	Timestamp time.Time
	Balance   map[string]uint64
}

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) GetTimestamp(blockID int) (time.Time, error) {
	if err := bc.isValidBlock(blockID); err != nil {
		return time.Time{}, err
	}

	return bc.blocks[blockID].Timestamp, nil
}

func (bc *BlockChain) GetBalance(userID string, blockID int) (uint64, error) {
	if err := bc.isValidBlock(blockID); err != nil {
		return 0, err
	}

	return bc.blocks[blockID].getUserBalance(userID)
}

func (bc *BlockChain) isValidBlock(blockID int) error {
	if bc == nil {
		return errors.New("block is not init")
	}

	height := len(bc.blocks)
	if height == 0 || blockID > height-1 {
		return errors.New("invalid block id")
	}

	return nil
}

func (b Block) getUserBalance(userID string) (uint64, error) {
	if b.Balance == nil {
		return 0, fmt.Errorf("balance is not init")
	}

	balance, found := b.Balance[userID]
	if !found {
		return 0, fmt.Errorf("user %s not found", userID)
	}

	return balance, nil
}
