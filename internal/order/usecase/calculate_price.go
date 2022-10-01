package usecase

import "github.com/Sup3r-Us3r/pfa-go/internal/order/entity"

type OrderInputDTO struct {
	Id    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

// Create a new instance for CalculateFinalPriceUseCase
func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

// Calculate the final order price and save that order
func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		Id:         order.Id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
