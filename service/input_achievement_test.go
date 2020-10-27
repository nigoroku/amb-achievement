package service

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"local.packages/models"
)

func TestAddInput(t *testing.T) {
	// DB接続
	db, err := sql.Open("mysql", "moizumi:base0210@tcp(localhost:3306)/ambitious_test?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)

	t.Parallel()

	ctx := context.Background()
	tx, _ := boil.BeginTx(ctx, nil)
	defer func() { _ = tx.Rollback() }()

	user := &models.User{Email: "tedt@fwdse.com"}
	user.Insert(ctx, tx, boil.Infer())

	oa := &models.InputAchievement{
		UserID:       user.UserID,
		Summary:      null.NewString("Summary1", true),
		ReferenceURL: null.NewString("http://test1.com", true),
		InputTime:    null.NewString("11:00", true),
	}

	inputService := &InputService{ctx, tx}

	c1 := &models.MCategory{CategoryID: 1, Name: "test1"}
	c2 := &models.MCategory{CategoryID: 2, Name: "test2"}
	c3 := &models.MCategory{CategoryID: 3, Name: "test3"}
	c1.Insert(ctx, tx, boil.Infer())
	c2.Insert(ctx, tx, boil.Infer())
	c3.Insert(ctx, tx, boil.Infer())

	err2 := inputService.AddInput(oa, []string{"1", "2", "3"})

	if err2 != nil {
		t.Error(err2)
	}

	got, err3 := inputService.FindByUser(user.UserID)

	if err3 != nil {
		t.Error(err3)
	}

	assert.Equal(t, got.Summary, oa.Summary)
	assert.Equal(t, got.ReferenceURL, oa.ReferenceURL)
}
