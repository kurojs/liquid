package blockchain

import (
	"testing"
	"time"
)

func TestBlockChain_GetBalanceAt(t *testing.T) {
	type args struct {
		userID string
		ts     time.Time
	}

	blocks := []*Block{
		{
			Timestamp: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 100,
				"B": 0,
				"C": 0,
			},
		},
		{
			Timestamp: time.Date(2000, 1, 2, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 100,
				"B": 0,
				"C": 0,
			},
		},
		{
			Timestamp: time.Date(2000, 2, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 90,
				"B": 0,
				"C": 10,
			},
		},
		{
			Timestamp: time.Date(2000, 3, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 90,
				"B": 5,
				"C": 5,
			},
		},
		{
			Timestamp: time.Date(2000, 4, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 40,
				"B": 50,
				"C": 10,
			},
		},
		{
			Timestamp: time.Date(2000, 5, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 30,
				"B": 30,
				"C": 40,
			},
		},
		{
			Timestamp: time.Date(2000, 6, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 100,
				"B": 0,
				"C": 0,
			},
		},
		{
			Timestamp: time.Date(2000, 7, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 10,
				"B": 70,
				"C": 20,
			},
		},
		{
			Timestamp: time.Date(2000, 8, 1, 1, 1, 1, 1, time.UTC),
			Balance: map[string]uint64{
				"A": 10,
				"B": 10,
				"C": 80,
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "get balance success when timestamp is equal with a block",
			args: args{
				userID: "B",
				ts:     time.Date(2000, 7, 1, 1, 1, 1, 1, time.UTC),
			},
			want:    70,
			wantErr: false,
		},
		{
			name: "get balance success when timestamp is between two blocks",

			args: args{
				userID: "A",
				ts:     time.Date(2000, 4, 2, 1, 1, 1, 1, time.UTC),
			},
			want:    40,
			wantErr: false,
		},
		{
			name: "get balance success when timestamp is greater than the last block's ts",
			args: args{
				userID: "A",
				ts:     time.Date(2001, 4, 2, 1, 1, 1, 1, time.UTC),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "get balance failed when timestamp smaller begin block",
			args: args{
				userID: "C",
				ts:     time.Date(1999, 1, 2, 1, 1, 1, 1, time.UTC),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlockChain{
				blocks: blocks,
			}
			got, err := bc.GetBalanceAt(tt.args.userID, tt.args.ts)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalanceAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBalanceAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
