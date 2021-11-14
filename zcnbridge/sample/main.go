package main

import (
	"context"

	"github.com/0chain/gosdk/zcnbridge"
	"github.com/0chain/gosdk/zcnbridge/config"
	"github.com/0chain/gosdk/zcnbridge/log"
	"go.uber.org/zap"
)

const (
	ConvertAmount = 10000000
)

// How should we manage nonce? - when user starts again on another server - how should we restore the value?

// 1. Init config
// 2. Init logs
// 2. Init SDK
// 3. Register wallet
// 4. Init bridge and make transactions

func main() {
	config.ParseClientConfig()
	config.Setup()
	zcnbridge.InitBridge()

	fromERCtoZCN()
	fromZCNtoERC()
}

func fromZCNtoERC() {
	trx, err := zcnbridge.BurnZCN(context.TODO(), config.Bridge.Value)
	if err != nil {
		log.Logger.Fatal("failed to burn", zap.Error(err), zap.String("hash", trx.Hash))
	}

	// ASK authorizers for burn tickets

	// Send burn tickets to Ethereum bridge

	// Confirm that transaction
}

func fromERCtoZCN() {
	transaction, err := zcnbridge.IncreaseBurnerAllowance(ConvertAmount)
	if err != nil {
		log.Logger.Fatal("failed to execute IncreaseBurnerAllowance", zap.Error(err))
	}

	res := zcnbridge.ConfirmTransactionStatus(transaction.Hash().Hex(), 60, 2)
	if res == 0 {
		log.Logger.Fatal("failed to confirm transaction", zap.String("hash", transaction.Hash().Hex()))
	}

	burnTrx, err := zcnbridge.BurnWZCN(ConvertAmount)
	burnTrxHash := burnTrx.Hash().Hex()
	if err != nil {
		log.Logger.Fatal("failed to execute BurnWZCN", zap.Error(err), zap.String("hash", burnTrxHash))
	}

	res = zcnbridge.ConfirmTransactionStatus(burnTrxHash, 60, 2)
	if res == 0 {
		log.Logger.Fatal("failed to confirm transaction ConfirmTransactionStatus", zap.String("hash", burnTrxHash))
	}

	mintPayload, err := zcnbridge.CreateMintPayload(burnTrxHash)
	if err != nil {
		log.Logger.Fatal("failed to CreateMintPayload", zap.Error(err), zap.String("hash", burnTrxHash))
	}

	trx, err := zcnbridge.MintZCN(context.TODO(), mintPayload)
	if err != nil {
		log.Logger.Fatal("failed to MintZCN", zap.Error(err), zap.String("hash", trx.Hash))
	}
}
