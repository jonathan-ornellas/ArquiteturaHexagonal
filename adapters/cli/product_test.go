package cli_test

import (
	"fmt"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/cli"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)




func TestRun(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productId := "abc"
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)	
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)


	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has the price %f and status %s", productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)



}