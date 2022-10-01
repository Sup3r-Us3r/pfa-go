package usecase

import "github.com/Sup3r-Us3r/pfa-go/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int `json:"total"`
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

// Create a new instance for GetTotalUseCase
func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

// Check the total number of registered orders
func (o *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	totalOrders, err := o.OrderRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{Total: totalOrders}, nil
}
