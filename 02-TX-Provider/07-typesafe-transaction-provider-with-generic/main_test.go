package typesafetransactionproviderwithgeneric

import (
	"context"
	"example/database"
	"fmt"
	"testing"
)

func TestAddPoint(t *testing.T) {

	dbConn := database.NewDatabase()
	db := dbConn.DbConn

	// var db *sql.DB

	repoUser := NewRepoUser(db)
	repoOrder := NewRepoOrder(db)

	adapter := NewAdapter(repoUser, repoOrder)

	tr := NewTrm(db, adapter)

	service := NewService(tr)
	_ = service.Create(context.Background(), "John Doe", []string{"item1", "item2"})

	fmt.Printf("service = %+v\n", service)

}
