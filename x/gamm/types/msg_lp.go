package types

type LiquidityChangeType int

const (
	AddLiquidity LiquidityChangeType = iota
	RemoveLiquidity
)

// SwapMsg defines a simple interface for getting the token denoms on a swap message route.
type LiquidityChangeMsg interface {
	LiquidityChangeType() LiquidityChangeType
}

var _ LiquidityChangeMsg = MsgExitPool{}
var _ LiquidityChangeMsg = MsgExitSwapShareAmountIn{}
var _ LiquidityChangeMsg = MsgExitSwapExternAmountOut{}

var _ LiquidityChangeMsg = MsgJoinPool{}
var _ LiquidityChangeMsg = MsgJoinSwapExternAmountIn{}
var _ LiquidityChangeMsg = MsgJoinSwapShareAmountOut{}

func (msg MsgExitPool) LiquidityChangeType() LiquidityChangeType {
	return RemoveLiquidity
}
func (msg MsgExitSwapShareAmountIn) LiquidityChangeType() LiquidityChangeType {
	return RemoveLiquidity
}
func (msg MsgExitSwapExternAmountOut) LiquidityChangeType() LiquidityChangeType {
	return RemoveLiquidity
}

func (msg MsgJoinPool) LiquidityChangeType() LiquidityChangeType {
	return AddLiquidity
}
func (msg MsgJoinSwapExternAmountIn) LiquidityChangeType() LiquidityChangeType {
	return AddLiquidity
}
func (msg MsgJoinSwapShareAmountOut) LiquidityChangeType() LiquidityChangeType {
	return AddLiquidity
}