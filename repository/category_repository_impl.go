package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into customer(name) values (?)"

	result, err := tx.ExecContext(
		ctx,
		sql,
		category.Name,
	)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(
		ctx,
		sql,
		category.Name,
		category.Id,
	)
	helper.PanicError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"

	_, err := tx.ExecContext(
		ctx,
		sql,
		category.Id,
	)
	helper.PanicError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	sql := "select id, name from category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicError(err)

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "select id, name from category"

	result, err := tx.QueryContext(ctx, sql)
	helper.PanicError(err)

	var categories []domain.Category

	for result.Next() {
		category := domain.Category{}

		err := result.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		categories = append(categories, category)
	}

	return categories
}
