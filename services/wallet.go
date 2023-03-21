package services

import (
	"context"
	"strconv"
	"wallet-service/config"
	"wallet-service/prisma/db"
)

type Wallet struct {
	client *db.PrismaClient
	ctx    context.Context
	env    *config.Envs
}

func (w *Wallet) Create() (*db.WalletModel, error) {
	initialBalance, err := strconv.Atoi(w.env.WALLET_BALANCE_INIT)
	if err != nil {
		return nil, err
	}

	wallet, err := w.client.Wallet.CreateOne(
		db.Wallet.Balance.Set(initialBalance),
	).Exec(w.ctx)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
