package ammswap

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestCreateTxPayload(t *testing.T) {
	type args struct {
		params *SwapParams
	}
	tests := []struct {
		name    string
		args    args
		want    *Payload
		wantErr bool
	}{
		{
			name: "case out",
			args: args{
				&SwapParams{
					Script:     "0x123::interface",
					FromCoin:   "0x123::BTC",
					ToCoin:     "0x123::APT",
					FromAmount: big.NewInt(1),
					ToAmount:   big.NewInt(266607),
					Slippage:   decimal.NewFromFloat(0.005),
				},
			},
			want: &Payload{
				Function: "0x123::interface::swap",
				TypeArgs: []string{
					"0x123::BTC",
					"0x123::APT",
				},
				Args: []string{
					"1",
					"265273",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSwapPayload(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTxPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTxPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
