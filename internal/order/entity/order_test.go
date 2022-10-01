package entity_test

import (
	"testing"

	"github.com/Sup3r-Us3r/pfa-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}

	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{Id: "1"}

	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{Id: "1", Price: 1.0}

	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10.0, 3.0)

	assert.NoError(t, err)
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 3.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10.0, 3.0)

	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 13.0, order.FinalPrice)
}
