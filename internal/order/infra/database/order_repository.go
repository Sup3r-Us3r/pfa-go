package database

import (
	"database/sql"

	"github.com/Sup3r-Us3r/pfa-go/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepository struct {
	Database *sql.DB
}

func NewOrderRepository(database *sql.DB) *OrderRepository {
	return &OrderRepository{Database: database}
}

func (o *OrderRepository) Save(order *entity.Order) error {
	stmt, err := o.Database.Prepare(`
		INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.Id, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) GetTotal() (int, error) {
	var totalOrders int

	err := o.Database.QueryRow(
		"SELECT COUNT(*) AS totalOrders FROM orders",
	).Scan(&totalOrders)

	if err != nil {
		return 0, err
	}

	return totalOrders, nil
}
