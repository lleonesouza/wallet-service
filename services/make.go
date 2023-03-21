package services

import (
	"context"
	"wallet-service/config"
	"wallet-service/prisma/db"
	"wallet-service/services/validation"

	"gopkg.in/go-playground/validator.v9"
)

type Services struct {
	User        *User
	Shopkeeper  *Shopkeeper
	Transaction *Transaction
	Wallet      *Wallet
	Validation  *validator.Validate
}

func MakeServices(env *config.Envs) *Services {
	client := db.NewClient()
	err := client.Prisma.Connect()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	wallet := &Wallet{client, ctx, env}

	return &Services{
		Validation: validation.MakeValidator(),
		User: &User{
			client: client,
			wallet: wallet,
			ctx:    ctx,
			env:    env,
		},
		Shopkeeper: &Shopkeeper{
			client: client,
			wallet: wallet,
			ctx:    ctx,
			env:    env,
		},
		Transaction: &Transaction{
			client: client,
			ctx:    ctx,
			env:    env,
		},
		Wallet: wallet,
	}
}
