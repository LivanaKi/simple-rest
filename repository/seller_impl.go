package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Users/natza/simple-rest/helper"
	"github.com/Users/natza/simple-rest/model"
)

type SellerImpl struct {
	Db *sql.DB
}

func NewSeller(Db *sql.DB) SellerRepository {
	return &SellerImpl{Db: Db}
}

// Delete
func (b *SellerImpl) Delete(ctx context.Context, sellerId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from sellers where id=$1"
	_, errEx := tx.ExecContext(ctx, SQL, sellerId)
	helper.PanicIfError(errEx)
}

// Create
func (b *SellerImpl) Create(ctx context.Context, seller model.Seller) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into sellers(name, phone) values($1, $2)"
	_, errEx := tx.ExecContext(ctx, SQL, seller.Name, seller.Phone)
	helper.PanicIfError(errEx)
}

// Read
func (b *SellerImpl) Read(ctx context.Context) []model.Seller {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name, phone from sellers"
	result, errReq := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errReq)
	defer result.Close()

	var sellers []model.Seller

	for result.Next() {
		seller := model.Seller{}
		errModel := result.Scan(&seller.Id, &seller.Name, &seller.Phone)
		helper.PanicIfError(errModel)

		sellers = append(sellers, seller)
	}
	return sellers
}

// Update
func (b *SellerImpl) Update(ctx context.Context, seller model.Seller) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update sellers set name=$1, phone=$2 where id=$3"
	_, err = tx.ExecContext(ctx, SQL, seller.Name, seller.Phone, seller.Id)
	helper.PanicIfError(err)
}

// Find by id
func (b *SellerImpl) FindById(ctx context.Context, sellerId int) (model.Seller, error) {
	tx, err := b.Db.Begin()

	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name, phone from sellers where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, sellerId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	seller := model.Seller{}

	if result.Next() {
		err := result.Scan(&seller.Id, &seller.Name, &seller.Phone)
		helper.PanicIfError(err)
		return seller, nil
	} else {
		return seller, errors.New("sellers id not found")
	}
}
