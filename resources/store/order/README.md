
### Delete purchase order by identifier. <a name="delete"></a>

For valid response try integer IDs with value < 1000. Anything above 1000 or non-integers will generate API errors.

**API Endpoint**: `DELETE /store/order/{orderId}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `orderId` | ✓ | ID of the order that needs to be deleted | `123` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	order "pets_go/resources/store/order"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Store.Order.Delete(order.DeleteRequest{
		OrderId: 123,
	})
}

```

### Find purchase order by ID. <a name="get"></a>

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

**API Endpoint**: `GET /store/order/{orderId}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `orderId` | ✓ | ID of order that needs to be fetched | `123` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	order "pets_go/resources/store/order"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Store.Order.Get(order.GetRequest{
		OrderId: 123,
	})
}

```

### Place an order for a pet. <a name="create"></a>

Place a new order in the store.

**API Endpoint**: `POST /store/order`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `complete` | ✗ |  | `true` |
| `id` | ✗ |  | `10` |
| `petId` | ✗ |  | `198772` |
| `quantity` | ✗ |  | `7` |
| `shipDate` | ✗ |  | `"1970-01-01T00:00:00"` |
| `status` | ✗ | Order Status | `OrderStatusEnumApproved` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	nullable "pets_go/nullable"
	order "pets_go/resources/store/order"
	types "pets_go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Store.Order.Create(order.CreateRequest{
		Id:       nullable.NewValue(10),
		PetId:    nullable.NewValue(198772),
		Quantity: nullable.NewValue(7),
		Status:   nullable.NewValue(types.OrderStatusEnumApproved),
	})
}

```

#### Response

##### Type
[Order](/types/order.go)

##### Example
`Order {
Id: nullable.NewValue(10),
PetId: nullable.NewValue(198772),
Quantity: nullable.NewValue(7),
Status: nullable.NewValue(OrderStatusEnumApproved),
}`
