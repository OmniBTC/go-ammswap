package ammswap

import (
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	EQUAL        = 0
	LESS_THAN    = 1
	GREATER_THAN = 2
)

func AmountMinOut(val *big.Int, slippage decimal.Decimal) *big.Int {
	return withSlippage(val, slippage)
}

func withSlippage(val *big.Int, slippage decimal.Decimal) *big.Int {
	return decimal.NewFromBigInt(val, 0).Sub(decimal.NewFromBigInt(val, 0).Mul(slippage)).BigInt()
}

func GetAmountOut(fromCoin, toCoin Coin, amount *big.Int, pool PoolResource) *big.Int {
	isSorted := IsSortedCoin(fromCoin, toCoin)
	var (
		reserveX *big.Int
		reserveY *big.Int
	)
	reserveX, reserveY = pool.CoinXReserve, pool.CoinYReserve
	if !isSorted {
		reserveX, reserveY = reserveY, reserveX
	}

	return getCoinOutWithFees(amount, reserveX, reserveY)
}

func getCoinOutWithFees(coinIn, reserveIn, reserveOut *big.Int) *big.Int {
	feePct, feeScale := getFee()
	feeMultiplier := big.NewInt(0).Sub(feeScale, feePct)

	coinInAfterFees := big.NewInt(0).Mul(coinIn, feeMultiplier)
	newReservesInSize := big.NewInt(0).Add(big.NewInt(0).Mul(reserveIn, feeScale), coinInAfterFees)
	return big.NewInt(0).Div(big.NewInt(0).Mul(coinInAfterFees, reserveOut), newReservesInSize)
}

func getFee() (*big.Int, *big.Int) {
	return big.NewInt(30), big.NewInt(10000)
}

func IsSortedCoin(coinx, coiny Coin) bool {
	comp := compare(coinx.TokenType.StructName+coinx.TokenType.Module, coiny.TokenType.StructName+coiny.TokenType.Module)
	if comp != EQUAL {
		return comp == LESS_THAN
	}

	return compare(coinx.TokenType.Address, coiny.TokenType.Address) == LESS_THAN
}

func compare(symbolX, symbolY string) int {
	ix := len(symbolX)
	iy := len(symbolY)
	lenCmp := cmp(ix, iy)
	// &bcs::to_bytes(utf8(b"hello")) in Aptos the first bytes contains length of string
	if lenCmp != EQUAL {
		return lenCmp
	}
	i := 0
	for i < ix && i < iy {
		elemCmp := cmp(int(symbolX[i]), int(symbolY[i]))
		if elemCmp != EQUAL {
			return elemCmp
		}
		i++
	}
	return EQUAL
}

func cmp(a, b int) int {
	if a == b {
		return EQUAL
	} else if a < b {
		return LESS_THAN
	} else {
		return GREATER_THAN
	}
}
