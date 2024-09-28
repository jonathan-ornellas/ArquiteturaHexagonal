package application_test

import (
	"testing"
	 uuid "github.com/satori/go.uuid"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater than zero", err.Error())

	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)
}

func TestProduct_GetId(t *testing.T) {
	product := application.Product{}
	id := uuid.NewV4().String()
	product.Id = id

	require.Equal(t, id, product.GetId())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"

	require.Equal(t, "Hello", product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.DISABLED

	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	require.Equal(t, 10.0, product.GetPrice())
}

