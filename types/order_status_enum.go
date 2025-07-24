package types

// Order Status
type OrderStatusEnum string

const (
	OrderStatusEnumApproved  OrderStatusEnum = "approved"
	OrderStatusEnumDelivered OrderStatusEnum = "delivered"
	OrderStatusEnumPlaced    OrderStatusEnum = "placed"
)
