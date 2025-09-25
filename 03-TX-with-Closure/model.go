package updatefn_closure

import "errors"

// Model

// Model

type UserModel struct {
	id    string     `json:"id"`
	email string     `json:"email"`
	point PointModel `json:"points"`
}

func (u *UserModel) ID() string {
	return u.id
}

func (u *UserModel) Email() string {
	return u.email
}

func (u *UserModel) AddUserPoint(points int) error {
	if points <= 0 {
		return errors.New("points must be greater than 0")
	}

	u.point.points += points

	return nil
}

type PointModel struct {
	id     string `json:"id"`
	points int    `json:"points"`
	userID string `json:"user_id"`
}

func (u *PointModel) Points() int {
	return u.points
}
