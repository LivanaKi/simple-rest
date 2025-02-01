package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Users/natza/simple-rest/helper"
	"github.com/Users/natza/simple-rest/model"
)

type SellerImpl struct {
	DB *sql.DB
}

func NewSeller(db *sql.DB) SellerRepository {
	return &SellerImpl{DB: db}
}

// Delete
func (b *SellerImpl) Delete(ctx context.Context, sellerID int) error {
	tx, err := b.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	SQL := "delete from sellers where id=$1"
	_, errEx := tx.ExecContext(ctx, SQL, sellerID)
	if errEx != nil {
		return errEx
	}
	return nil
}

// Create
func (b *SellerImpl) Save(ctx context.Context, seller model.Seller) error {
	tx, err := b.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	SQL := "insert into sellers(name, phone) values($1, $2)"
	err = tx.QueryRowContext(ctx, SQL, seller.Name, seller.Phone).Scan(&seller.ID)
	if err != nil {
		return err
	}
	return nil
}

// Read
func (b *SellerImpl) Read(ctx context.Context) ([]model.Seller, error) {
	tx, err := b.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name, phone from sellers"
	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var sellers []model.Seller

	for result.Next() {
		seller := model.Seller{}
		err = result.Scan(&seller.ID, &seller.Name, &seller.Phone)
		if err != nil {
			return nil, err
		}

		sellers = append(sellers, seller)
	}
	return sellers, nil
}

// Update
func (b *SellerImpl) Update(ctx context.Context, seller model.Seller) error {
	tx, err := b.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	SQL := "update sellers set name=$1, phone=$2 where id=$3"
	_, err = tx.ExecContext(ctx, SQL, seller.Name, seller.Phone, seller.ID)
	if err != nil {
		return err
	}
	return nil
}

// Find by id
func (b *SellerImpl) FindByID(ctx context.Context, sellerID int) (model.Seller, error) {
	tx, err := b.DB.Begin()
	if err != nil {
		return model.Seller{}, err
	}
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name, phone from sellers where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, sellerID)
	if errQuery != nil {
		return model.Seller{}, errQuery
	}
	defer result.Close()

	seller := model.Seller{}

	if result.Next() {
		err := result.Scan(&seller.ID, &seller.Name, &seller.Phone)
		if err != nil {
			return model.Seller{}, err
		}
		return seller, nil
	}
	return seller, errors.New("sellers id not found")
}
