package main

import (
	//"fmt"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	mathEth "github.com/ethereum/go-ethereum/common/math"
	flip "github.com/griffindavis02/package"
)

type Transaction struct {
	from        string
	to          string
	value       *big.Int
	receiptHash string
	nonce       uint
}

var (
	mjsonOut         flip.Output
	mNonce           uint = 0
	mGasPrice        uint = 21000
	mReceipt         string
	mTransactionTrie []Transaction
)

func main() {
	arrRates := []float64{0.1, 0.01, 0.001}

	// Initialize testing environment with error rates and output
	flip.Initalize("iteration", 10, arrRates, &mjsonOut)

	// Set up transaction
	wei, _ := toWei(0.05, "ether")
	receipt := sendTransaction("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8",
		"0xac03bb73b6a9e108530aff4df5077c2b3d481e5a",
		wei,
		22000)

	fmt.Println(mTransactionTrie)
	fmt.Println(receipt)

	// mjsonOut.PostAPI("http://localhost:5000/express")
}

func sendTransaction(from string, to string, value *big.Int, gasCap uint) string {
	fAddress, _ := new(big.Int).SetString(from[2:], 16)
	tAddress, _ := new(big.Int).SetString(to[2:], 16)
	gasCap = uint(mathEth.U256(big.NewInt(int64(gasCap))).Uint64())

	fAddress = (&mjsonOut).BitFlip(fAddress)
	tAddress = (&mjsonOut).BitFlip(tAddress)
	value = (&mjsonOut).BitFlip(value)

	hash := big.NewInt(0).
		Rand(rand.New(rand.NewSource(time.Now().Unix())),
			mathEth.U256(big.NewInt(mathEth.MaxBig256.Int64())))

	if gasCap > mGasPrice {
		mTransactionTrie = append(mTransactionTrie, Transaction{
			from:        "0x" + hex.EncodeToString(fAddress.Bytes()),
			to:          "0x" + hex.EncodeToString(tAddress.Bytes()),
			value:       value,
			receiptHash: "0x" + hex.EncodeToString(hash.Bytes()),
			nonce:       mNonce,
		})
		mNonce++
	}

	hash = (&mjsonOut).BitFlip(hash)
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
		returnValue = (&mjsonOut).BitFlip(returnValue)
		return returnValue, nil
	case "gwei":
		returnValue := big.NewInt(0).Mul(bigValue, big.NewInt(int64(math.Pow(10, 18-multiplicand))))
		returnValue = (&mjsonOut).BitFlip(returnValue)
		return big.NewInt(0).Mul(bigValue, big.NewInt(int64(math.Pow(10, 9-multiplicand)))), nil
	}
	return nil, errors.New("invalid_factor")
}
