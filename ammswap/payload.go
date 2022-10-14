package ammswap

import (
	"errors"
	"math/big"

	"github.com/shopspring/decimal"
)

type Payload struct {
	Function string
	TypeArgs []string
	Args     []string
}

type SwapParams struct {
	Script     string // eg. 0xd415c5143d4f9752e462ab3476c567fdc0e2f0fb02f779d333e819c0e8624ea8::interface
	FromCoin   string
	ToCoin     string
	FromAmount *big.Int
	ToAmount   *big.Int
	Slippage   decimal.Decimal
}

func CreateSwapPayload(params *SwapParams) (*Payload, error) {
	if nil == params {
		return nil, errors.New("invalid params: nil")
	}

	if params.Slippage.LessThan(decimal.Zero) || params.Slippage.GreaterThan(decimal.New(1, 0)) {
		return nil, errors.New("Invalid slippage value:" + params.Slippage.String())
	}

	f := "swap"
	functionName := params.Script + "::" + f
	typeArgs := []string{
		params.FromCoin,
		params.ToCoin,
	}
	params.ToAmount = withSlippage(params.ToAmount, params.Slippage)
	args := []string{
		params.FromAmount.String(),
		params.ToAmount.String(),
	}
	return &Payload{
		Function: functionName,
		TypeArgs: typeArgs,
		Args:     args,
	}, nil
}
