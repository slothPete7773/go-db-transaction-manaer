package tx_injection

import (
	"context"
	"example/database"
	"log"
	"testing"
)

func TestAddPoint(t *testing.T) {

	db := database.NewDatabase()

	userService := NewUserService(db.DbConn)

	err := userService.AddPoint(context.TODO(), "ID-123", 10)
	if err != nil {
		log.Fatal(err)
	}

}
