package test_order_client

import (
	fmt "fmt"
	sdk "pets_go/client"
	nullable "pets_go/nullable"
	order "pets_go/resources/store/order"
	types "pets_go/types"
	testing "testing"
)

func TestDelete200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Store.Order.Delete(order.DeleteRequest{
		OrderId: 123,
	})

	if err != nil {
		t.Fatalf("TestDelete200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestGet200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Store.Order.Get(order.GetRequest{
		OrderId: 123,
	})

	if err != nil {
		t.Fatalf("TestGet200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Store.Order.Create(order.CreateRequest{
		Complete: nullable.NewValue(true),
		Id:       nullable.NewValue(10),
		PetId:    nullable.NewValue(198772),
		Quantity: nullable.NewValue(7),
		ShipDate: nullable.NewValue("1970-01-01T00:00:00"),
		Status:   nullable.NewValue(types.OrderStatusEnumApproved),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
