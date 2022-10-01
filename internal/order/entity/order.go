package entity

import "errors"

type Order struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
}

// Check if required fields have been filled
func (o Order) IsValid() error {
	if o.Id == "" {
		return errors.New("invalid id")
	}

	if o.Price == 0 {
		return errors.New("invalid price")
	}

	if o.Tax == 0 {
		return errors.New("invalid tax")
	}

	return nil
}

// Calculate the final price of the Order
func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.IsValid()

	if err != nil {
		return err
	}

	return nil
}

// Create a new instance for Order
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		Id:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()

	if err != nil {
		return nil, err
	}

	return order, nil
}
