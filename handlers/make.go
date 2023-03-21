package handlers

import (
	"wallet-service/config"
	"wallet-service/services"
)

type Handlers struct {
	Shopkeeper  ShopkeeperHandler
	User        UserHandler
	Transaction TransactionHandler
}

func MakeHandlers(env *config.Envs) *Handlers {
	services := services.MakeServices(env)

	return &Handlers{
		Shopkeeper: ShopkeeperHandler{
			service: services,
			env:     env,
		},
		User: UserHandler{
			service: services,
			env:     env,
		},
		Transaction: TransactionHandler{
			service: services,
			env:     env,
		},
	}
}
