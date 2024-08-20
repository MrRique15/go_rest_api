package repositories

import (
	"errors"

	"github.com/MrRique15/go_rest_api/main_api/pkg/shared/models"

	"database/sql"
)

type PostgresDBHandlerShipping struct {
	shippingDatabase *sql.DB
}

func (dbh *PostgresDBHandlerShipping) InitiateCollection() {
	dbh.shippingDatabase = ShippingPostgresDB
}

func (dbh PostgresDBHandlerShipping) getShippingById(id int) (models.Shipping, error) {
	sql_query := `SELECT _id, shipping_type, order_id, shipping_price FROM shipping WHERE _id = $1`

	var shipping models.Shipping

	error_sql := dbh.shippingDatabase.QueryRow(
		sql_query,
		id).Scan(
		&shipping.ID,
		&shipping.ShippingType,
		&shipping.OrderID,
		&shipping.ShippingPrice)

	if error_sql != nil {
		return models.Shipping{}, errors.New("error during shipping get by id")
	}

	return shipping, nil
}

func (dbh PostgresDBHandlerShipping) registerShipping(newShipping models.Shipping) (models.Shipping, error) {
	sql_query := `INSERT INTO shipping (shipping_type, order_id, shipping_price) VALUES ($1, $2, $3) RETURNING _id, order_id, shipping_type, shipping_price`

	var shipping_register models.Shipping

	error_sql := dbh.shippingDatabase.QueryRow(
		sql_query,
		newShipping.ShippingType,
		newShipping.OrderID,
		newShipping.ShippingPrice).Scan(
		&shipping_register.ID,
		&shipping_register.OrderID,
		&shipping_register.ShippingType,
		&shipping_register.ShippingPrice)

	if error_sql != nil {
		return models.Shipping{}, errors.New("error during shipping registration")
	}

	return shipping_register, nil
}

func (dbh PostgresDBHandlerShipping) getShippingByOrderId(orderId string) (models.Shipping, error) {
	sql_query := `SELECT _id, shipping_type, order_id, shipping_price FROM shipping WHERE order_id = $1`

	var shipping models.Shipping

	error_sql := dbh.shippingDatabase.QueryRow(
		sql_query,
		orderId).Scan(
		&shipping.ID,
		&shipping.ShippingType,
		&shipping.OrderID,
		&shipping.ShippingPrice)

	if error_sql != nil {
		return models.Shipping{}, error_sql
	}

	return shipping, nil
}

func (dbh PostgresDBHandlerShipping) listShipping() ([]models.Shipping, error) {
	sql_query := `SELECT _id, shipping_type, order_id, shipping_price FROM shipping`

	rows, error_sql := dbh.shippingDatabase.Query(sql_query)

	if error_sql != nil {
		return []models.Shipping{}, errors.New("error during shipping listing")
	}

	var shippings []models.Shipping

	for rows.Next() {
		var shipping models.Shipping

		rows.Scan(
			&shipping.ID,
			&shipping.ShippingType,
			&shipping.OrderID,
			&shipping.ShippingPrice)

		shippings = append(shippings, shipping)
	}

	return shippings, nil
}