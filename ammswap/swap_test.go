package ammswap

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestGetAmountOut(t *testing.T) {
	type args struct {
		fromCoin Coin
		toCoin   Coin
		amount   *big.Int
		pool     PoolResource
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "case out",
			args: args{
				fromCoin: Coin{
					TokenType: TokenType{
						Address:    "",
						Module:     "",
						StructName: "",
					},
				},
				toCoin: Coin{
					TokenType: TokenType{
						Address:    "",
						Module:     "",
						StructName: "",
					},
				},
				amount: big.NewInt(1000000),
				pool: PoolResource{
					CoinXReserve: big.NewInt(10415880990),
					CoinYReserve: big.NewInt(3004784231600),
				},
			},
			want: big.NewInt(3456),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAmountOut(tt.args.fromCoin, tt.args.toCoin, tt.args.amount, tt.args.pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSlippage(t *testing.T) {
	type args struct {
		val      *big.Int
		slippage decimal.Decimal
		mod      int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "case out(from)",
			args: args{
				val:      big.NewInt(266607),
				slippage: decimal.NewFromFloat(0.005),
				mod:      -1,
			},
			want: big.NewInt(265273),
		},
		{
			name: "case in(to)",
			args: args{
				val:      big.NewInt(750174),
				slippage: decimal.NewFromFloat(0.005),
				mod:      1,
			},
			want: big.NewInt(753924),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withSlippage(tt.args.val, tt.args.slippage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSlippage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmountMinOut(t *testing.T) {
	type args struct {
		val      *big.Int
		slippage decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "case out(from)",
			args: args{
				val:      big.NewInt(266607),
				slippage: decimal.NewFromFloat(0.005),
			},
			want: big.NewInt(265273),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AmountMinOut(tt.args.val, tt.args.slippage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AmountMinOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compare(t *testing.T) {
	type args struct {
		symbolX string
		symbolY string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "APT BTC",
			args: args{"APT", "BTC"},
			want: LESS_THAN,
		},
		{
			name: "BTC APT",
			args: args{"BTC", "APT"},
			want: GREATER_THAN,
		},
		{
			name: "BTC USDT",
			args: args{"BTC", "USDT"},
			want: LESS_THAN,
		},
		{
			name: "USDT BTC",
			args: args{"USDT", "BTC"},
			want: GREATER_THAN,
		},
		{
			name: "USDC USDT",
			args: args{"USDC", "USDT"},
			want: LESS_THAN,
		},
		{
			name: "APT USDT",
			args: args{"APT", "USDT"},
			want: LESS_THAN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.symbolX, tt.args.symbolY); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
