// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO Orders(VisitId, DishId, Rating, ReviewText)
VALUES (?, ?, ?, ?)
`

type CreateOrderParams struct {
	Visitid    sql.NullInt64
	Dishid     sql.NullInt64
	Rating     sql.NullFloat64
	Reviewtext sql.NullString
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder,
		arg.Visitid,
		arg.Dishid,
		arg.Rating,
		arg.Reviewtext,
	)
	return err
}

const createVisit = `-- name: CreateVisit :exec
INSERT INTO Visit(Date, Time, UserId, RestaurantId)
VALUES (?, ?, ?, ?)
`

type CreateVisitParams struct {
	Date         time.Time
	Time         interface{}
	Userid       sql.NullInt64
	Restaurantid sql.NullInt64
}

func (q *Queries) CreateVisit(ctx context.Context, arg CreateVisitParams) error {
	_, err := q.db.ExecContext(ctx, createVisit,
		arg.Date,
		arg.Time,
		arg.Userid,
		arg.Restaurantid,
	)
	return err
}

const getOrdersForVisit = `-- name: GetOrdersForVisit :many
SELECT d.Name, o.Rating, o.ReviewText from
    Orders o join Dish d on o.DishId = d.id
    where o.VisitId = ?
`

type GetOrdersForVisitRow struct {
	Name       string
	Rating     sql.NullFloat64
	Reviewtext sql.NullString
}

func (q *Queries) GetOrdersForVisit(ctx context.Context, visitid sql.NullInt64) ([]GetOrdersForVisitRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersForVisit, visitid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrdersForVisitRow
	for rows.Next() {
		var i GetOrdersForVisitRow
		if err := rows.Scan(&i.Name, &i.Rating, &i.Reviewtext); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRestaurantHistory = `-- name: GetRestaurantHistory :many
SELECT id, Date, Time from Visit
    where UserId = ? and RestaurantId = ?
`

type GetRestaurantHistoryParams struct {
	Userid       sql.NullInt64
	Restaurantid sql.NullInt64
}

type GetRestaurantHistoryRow struct {
	ID   int64
	Date time.Time
	Time interface{}
}

func (q *Queries) GetRestaurantHistory(ctx context.Context, arg GetRestaurantHistoryParams) ([]GetRestaurantHistoryRow, error) {
	rows, err := q.db.QueryContext(ctx, getRestaurantHistory, arg.Userid, arg.Restaurantid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRestaurantHistoryRow
	for rows.Next() {
		var i GetRestaurantHistoryRow
		if err := rows.Scan(&i.ID, &i.Date, &i.Time); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRestaurantsLike = `-- name: GetRestaurantsLike :many
SELECT id, name, area, address, mapslink, mapsratingoutof5 FROM Restaurant where Name like ?
`

func (q *Queries) GetRestaurantsLike(ctx context.Context, name string) ([]Restaurant, error) {
	rows, err := q.db.QueryContext(ctx, getRestaurantsLike, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Restaurant
	for rows.Next() {
		var i Restaurant
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Area,
			&i.Address,
			&i.Mapslink,
			&i.Mapsratingoutof5,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
