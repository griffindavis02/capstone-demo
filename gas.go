package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"time"

	mathEth "github.com/ethereum/go-ethereum/common/math"
	"github.com/griffindavis02/eth-bit-flip/config"
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
	// intTest          int    = 1
	// intTest32        int32  = 2
	// intTest64        int64  = 3
	// uintTest         uint   = 4
	// uintTest32       uint32 = 5
	// uintTest64       uint64 = 6
)

func main() {

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	bytes, err := json.MarshalIndent(cfg, "", "    ")
	if err == nil {
		fmt.Println(string(bytes))
	}

	// Set up transaction
	for range make([]int, 20) {
		// intTest = injection.BitFlip(intTest, &cfg).(int)
		// // time.Sleep(time.Millisecond)
		// intTest32 = injection.BitFlip(intTest32, &cfg).(int32)
		// // time.Sleep(time.Millisecond)
		// intTest64 = injection.BitFlip(intTest64, &cfg).(int64)
		// // time.Sleep(time.Millisecond)
		// uintTest = injection.BitFlip(uintTest, &cfg).(uint)
		// // time.Sleep(time.Millisecond)
		// uintTest32 = injection.BitFlip(uintTest32, &cfg).(uint32)
		// // time.Sleep(time.Millisecond)
		// uintTest64 = injection.BitFlip(uintTest64, &cfg).(uint64)
		// time.Sleep(time.Millisecond)
		wei, _ := toWei(0.05, "ether", &cfg)
		sendTransaction("0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8",
			"0xac03bb73b6a9e108530aff4df5077c2b3d481e5a",
			wei,
			22000,
			&cfg)
	}
	// mjsonOut.PostAPI("http://localhost:5000/express")
}

func sendTransaction(from string, to string, value *big.Int, gasCap uint, cfg *config.Config) string {

	gasCap = uint(mathEth.U256(big.NewInt(int64(gasCap))).Uint64())

	fAddress := injection.BitFlip(from[2:], cfg).(string)
	tAddress := injection.BitFlip(to[2:], cfg).(string)
	value = injection.BitFlip(value, cfg).(*big.Int)
	// fmt.Printf("%s\n%s\n%v\n\n", fAddress, tAddress, value)

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

	hash = injection.BitFlip(hash, cfg).(*big.Int)
	return "0x" + hex.EncodeToString((hash.Bytes()))
}

func toWei(value float64, factor string, cfg *config.Config) (*big.Int, error) {
	var multiplicand float64 = 0
	for value < 1 {
		multiplicand++
		value *= 10
	}
	bigValue := mathEth.U256(big.NewInt(int64(value)))

	switch factor {
	case "ether":
		returnValue := big.NewInt(0).Mul(bigValue, big.NewInt(int64(math.Pow(10, 18-multiplicand))))
		returnValue = injection.BitFlip(returnValue, cfg).(*big.Int)
		return returnValue, nil
	case "gwei":
		returnValue := big.NewInt(0).Mul(bigValue, big.NewInt(int64(math.Pow(10, 18-multiplicand))))
		returnValue = injection.BitFlip(returnValue, cfg).(*big.Int)
		return returnValue, nil
	}
	return nil, errors.New("invalid_factor")
}
