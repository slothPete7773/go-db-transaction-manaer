package main

import "errors"

// Model

type UserModel struct {
	id    int
	email string
	point PointModel
}

func (u *UserModel) ID() int {
	return u.id
}

func (u *UserModel) Email() string {
	return u.email
}

func (u *UserModel) UsePointsAsDiscount(points int) error {
	if points <= 0 {
		return errors.New("points must be greater than 0")
	}

	if u.point.points < points {
		return errors.New("not enough points")
	}

	u.point.points -= points

	return nil
}

type PointModel struct {
	points int
}

func (u *PointModel) Points() int {
	return u.points
}
