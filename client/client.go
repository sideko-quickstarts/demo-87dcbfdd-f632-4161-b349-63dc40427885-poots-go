package client

import (
	sdkcore "pets_go/core"
	pet "pets_go/resources/pet"
	store "pets_go/resources/store"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	Pet        *pet.Client
	Store      *store.Client
}

// Instantiate a new API client
func NewClient(builders ...func(*sdkcore.CoreClient)) *Client {
	defaultEnv := sdkcore.DefaultBaseURL(Environment.String())
	coreClient := sdkcore.NewCoreClient(defaultEnv)
	for _, b := range builders {
		b(coreClient)
	}

	client := Client{
		coreClient: coreClient,
		Pet:        pet.NewClient(coreClient),
		Store:      store.NewClient(coreClient),
	}

	return &client
}
