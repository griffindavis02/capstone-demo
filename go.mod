module github.com/griffindavis02/go-chain

go 1.17

replace github.com/griffindavis02/eth-bit-flip => ..\eth-bit-flip

require (
	github.com/ethereum/go-ethereum v1.10.9
	github.com/griffindavis02/eth-bit-flip v0.1.0
)

require gopkg.in/urfave/cli.v1 v1.20.0 // indirect
