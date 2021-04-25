package blockchain

import (
	"errors"
	"time"
)

func (bc *BlockChain) GetBalanceAt(userID string, ts time.Time) (uint64, error) {
	if bc == nil {
		return 0, errors.New("blockchain is not init")
	}

	height := len(bc.blocks)
	return bc.getBalance(userID, ts, 0, height-1)
}

func (bc *BlockChain) getBalance(userID string, ts time.Time, from, to int) (uint64, error) {
	if from >= to {
		t, err := bc.GetTimestamp(to)
		if err != nil {
			return 0, err
		}

		// Not founded balance because of ts > t
		if t.After(ts) {
			return 0, errors.New("balance is not found")
		}

		return bc.GetBalance(userID, to)
	}

	middle := (from + to) / 2
	middleTs, err := bc.GetTimestamp(middle)
	if err != nil {
		return 0, err
	}

	if ts == middleTs {
		return bc.GetBalance(userID, middle)
	}

	if ts.Before(middleTs) {
		return bc.getBalance(userID, ts, from, middle-1)
	}

	if ts.After(middleTs) {
		// If ts is after middle ts there are two cases happen:
		// The ts is in somewhere between middle < ts < next block of middle
		if nextTs, err := bc.GetTimestamp(middle + 1); err == nil && nextTs.After(ts) {
			return bc.GetBalance(userID, middle)
		}

		// Or ts > next block of middle
		return bc.getBalance(userID, ts, middle+1, to)
	}

	return 0, errors.New("balance is not found")
}
