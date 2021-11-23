package main

import (
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"math/rand"
	"time"

	mathEth "github.com/ethereum/go-ethereum/common/math"
	"github.com/griffindavis02/eth-bit-flip/injection"
)

type Transaction struct {
	from        string
	to          string
	value       *big.Int
	receiptHash string
	nonce       uint
}

var (
	mNonce           uint = 0
	mGasPrice        uint = 21000
	mTransactionTrie []Transaction
)

func main() {

	// Initialize testing environment with error rates and output

	// Set up transaction
	for range make([]int, 10000) {
		wei, _ := toWei(0.05, "ether")
		sendTransaction("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8",
			"0xac03bb73b6a9e108530aff4df5077c2b3d481e5a",
			wei,
			22000)
	}
	// mjsonOut.PostAPI("http://localhost:5000/express")
}

func sendTransaction(from string, to string, value *big.Int, gasCap uint) string {

	gasCap = uint(mathEth.U256(big.NewInt(int64(gasCap))).Uint64())

	fAddress := injection.BitFlip(from[2:], `E:\Libraries\Documents\Projects\Code\Capstone\go-ethereum\cmd\flipconfig\flipconfig.json`).(string)
	tAddress := injection.BitFlip(to[2:], `E:\Libraries\Documents\Projects\Code\Capstone\go-ethereum\cmd\flipconfig\flipconfig.json`).(string)
	value = injection.BitFlip(value, `E:\Libraries\Documents\Projects\Code\Capstone\go-ethereum\cmd\flipconfig\flipconfig.json`).(*big.Int)

	hash := big.NewInt(0).
		Rand(rand.New(rand.NewSource(time.Now().UnixNano())),
			mathEth.U256(big.NewInt(mathEth.MaxBig256.Int64())))

	if gasCap > mGasPrice {
		mTransactionTrie = append(mTransactionTrie, Transaction{
			from:        "0x" + hex.EncodeToString([]byte(fAddress)),
			to:          "0x" + hex.EncodeToString([]byte(tAddress)),
			value:       value,
			receiptHash: "0x" + hex.EncodeToString(hash.Bytes()),
			nonce:       mNonce,
		})
		mNonce++
	}

	hash = injection.BitFlip(hash, `E:\Libraries\Documents\Projects\Code\Capstone\go-ethereum\cmd\flipconfig\flipconfig.json`).(*big.Int)
	return "0x" + hex.EncodeToString((hash.Bytes()))
}

func toWei(value float64, factor string) (*big.Int, error) {
	var multiplicand float64 = 0
	for value < 1 {
		multiplicand++
		value *= 10
	}
	bigValue := mathEth.U256(big.NewInt(int64(value)))

	switch factor {
	case "ether":
		returnValue := big.NewInt(0).Mul(bigValue, big.NewInt(int64(math.Pow(10, 18-multiplicand))))
		returnValue = injection.BitFlip(returnValue, `E:\Libraries\Documents\Projects\Code\Capstone\go-ethereum\cmd\flipconfig\flipconfig.json`).(*big.Int)
		return returnValue, nil
	case "gwei":
		returnValue := big.NewInt(0).Mul(bigValue, big.NewInt(int64(math.Pow(10, 18-multiplicand))))
		returnValue = injection.BitFlip(returnValue, `E:\Libraries\Documents\Projects\Code\Capstone\go-ethereum\cmd\flipconfig\flipconfig.json`).(*big.Int)
		return returnValue, nil
	}
	return nil, errors.New("invalid_factor")
}
