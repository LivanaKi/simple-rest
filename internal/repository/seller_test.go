package repository

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"

	"github.com/Users/natza/simple-rest/internal/model"
)

func TestSave(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewSeller(db.DB)
	seller := &model.Seller{Name: "test", Phone: "+11111"}

	mock.ExpectBegin()
	mock.ExpectQuery("insert into sellers").
		WithArgs(seller.Name, seller.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectCommit()

	err = repo.Save(ctx, seller)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRead(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewSeller(db.DB)
	expectedSeller := []model.Seller{
		{ID: 1, Name: "test", Phone: "+11111"},
		{ID: 2, Name: "test1", Phone: "+22222"},
	}

	mock.ExpectBegin()
	mock.ExpectQuery("select id, name, phone from sellers").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone"}).
			AddRow(expectedSeller[0].ID, expectedSeller[0].Name, expectedSeller[0].Phone).
			AddRow(expectedSeller[1].ID, expectedSeller[1].Name, expectedSeller[1].Phone))
	mock.ExpectCommit()

	sellers, err := repo.Read(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedSeller, sellers)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewSeller(db.DB)
	seller := &model.Seller{ID: 1, Name: "test", Phone: "+11111"}

	mock.ExpectBegin()
	mock.ExpectExec("update sellers set name=").
		WithArgs(seller.Name, seller.Phone, seller.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = repo.Update(ctx, seller)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewSeller(db.DB)

	mock.ExpectBegin()
	mock.ExpectExec("delete from sellers where id=").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = repo.Delete(ctx, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByID(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := NewSeller(db.DB)
	expectedSeller := &model.Seller{ID: 1, Name: "test", Phone: "+11111"}

	mock.ExpectBegin()
	mock.ExpectQuery("select id, name, phone from sellers where id=").WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone"}).
			AddRow(expectedSeller.ID, expectedSeller.Name, expectedSeller.Phone))
	mock.ExpectCommit()

	sellers, err := repo.FindByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedSeller, sellers)
	assert.NoError(t, mock.ExpectationsWereMet())
}
