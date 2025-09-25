package typesafetransactionproviderwithgeneric

import "context"

// repoTx interface defines operations that must be done within a transaction.
type repoTx interface {
	CreateUser(ctx context.Context, name string) error
	CreateOrder(ctx context.Context, items []string) error
}

// [T] is any type that satisfies the required interface (in this case, repoTx).
type transactor[T any] interface {
	InTx(ctx context.Context, fn func(T) error) error
}

type User struct {
	ID   int64
	Name string
}
