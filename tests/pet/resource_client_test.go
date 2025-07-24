package test_pet_client

import (
	fmt "fmt"
	sdk "pets_go/client"
	sdkcore "pets_go/core"
	nullable "pets_go/nullable"
	pet "pets_go/resources/pet"
	types "pets_go/types"
	testing "testing"
)

func TestDelete200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.Delete(pet.DeleteRequest{
		PetId: 123,
	})

	if err != nil {
		t.Fatalf("TestDelete200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestFindByStatus200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.FindByStatus(pet.FindByStatusRequest{
		Status: nullable.NewValue(types.PetFindByStatusStatusEnumAvailable),
	})

	if err != nil {
		t.Fatalf("TestFindByStatus200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestFindByStatus200SuccessRequiredOnly(t *testing.T) {
	// Success test using only required fields
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.FindByStatus(pet.FindByStatusRequest{})

	if err != nil {
		t.Fatalf("TestFindByStatus200SuccessRequiredOnly - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestGet200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.Get(pet.GetRequest{
		PetId: 123,
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
	res, err := client.Pet.Create(pet.CreateRequest{
		Category: nullable.NewValue(types.Category{
			Id:   nullable.NewValue(1),
			Name: nullable.NewValue("Dogs"),
		}),
		Id:   nullable.NewValue(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Status: nullable.NewValue(types.PetStatusEnumAvailable),
		Tags: nullable.NewValue([]types.Tag{
			types.Tag{
				Id:   nullable.NewValue(123),
				Name: nullable.NewValue("string"),
			},
		}),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestUploadImage200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.UploadImage(pet.UploadImageRequest{
		Data:               sdkcore.NewInMemoryFile("test.pdf", "123"),
		PetId:              123,
		AdditionalMetadata: nullable.NewValue("string"),
	})

	if err != nil {
		t.Fatalf("TestUploadImage200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestUploadImage200SuccessRequiredOnly(t *testing.T) {
	// Success test using only required fields
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.UploadImage(pet.UploadImageRequest{
		Data:  sdkcore.NewInMemoryFile("test.pdf", "123"),
		PetId: 123,
	})

	if err != nil {
		t.Fatalf("TestUploadImage200SuccessRequiredOnly - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestUpdate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithApiKey("API_KEY"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.Pet.Update(pet.UpdateRequest{
		Category: nullable.NewValue(types.Category{
			Id:   nullable.NewValue(1),
			Name: nullable.NewValue("Dogs"),
		}),
		Id:   nullable.NewValue(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Status: nullable.NewValue(types.PetStatusEnumAvailable),
		Tags: nullable.NewValue([]types.Tag{
			types.Tag{
				Id:   nullable.NewValue(123),
				Name: nullable.NewValue("string"),
			},
		}),
	})

	if err != nil {
		t.Fatalf("TestUpdate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
