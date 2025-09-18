package main

type UsePointsAsDiscountHandler struct {
}

type txProvider interface {
	Transact(txFunc func(adapters Adapters) error) error
}

type Adapters struct {
	UserRepository  UserRepository
	PointRepository pointRepository
}
