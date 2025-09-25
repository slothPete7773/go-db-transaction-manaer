package typesafetransactionproviderwithgeneric

import "context"

type Adapter struct {
	repoUser  *RepoUser
	repoOrder *RepoOrder
}

func NewAdapter(repoUser *RepoUser, repoOrder *RepoOrder) *Adapter {
	return &Adapter{
		repoUser:  repoUser,
		repoOrder: repoOrder,
	}
}

func (a *Adapter) WithTx(tx Transaction) *Adapter {
	return &Adapter{
		repoUser:  a.repoUser.WithTx(tx),
		repoOrder: a.repoOrder.WithTx(tx),
	}
}

func (a *Adapter) CreateUser(ctx context.Context, name string) error {
	return a.repoUser.CreateUser(ctx, name)
}

func (a *Adapter) CreateOrder(ctx context.Context, items []string) error {
	return a.repoOrder.CreateOrder(ctx, items)
}
