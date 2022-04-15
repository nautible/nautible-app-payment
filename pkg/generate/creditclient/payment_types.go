// Package creditclient provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package creditclient

// RestCreateCreditPayment defines model for RestCreateCreditPayment.
type RestCreateCreditPayment struct {
	CustomerId int32  `json:"customerId"`
	OrderDate  string `json:"orderDate"`
	OrderNo    string `json:"orderNo"`
	TotalPrice int32  `json:"totalPrice"`
}

// RestCreditPayment defines model for RestCreditPayment.
type RestCreditPayment struct {
	AcceptDate *string `json:"acceptDate,omitempty"`
	AcceptNo   *string `json:"acceptNo,omitempty"`
	CustomerId *int32  `json:"customerId,omitempty"`
	OrderDate  *string `json:"orderDate,omitempty"`
	OrderNo    *string `json:"orderNo,omitempty"`
	TotalPrice *int32  `json:"totalPrice,omitempty"`
}

// RestUpdateCreditPayment defines model for RestUpdateCreditPayment.
type RestUpdateCreditPayment struct {
	AcceptDate *string `json:"acceptDate,omitempty"`
	AcceptNo   string  `json:"acceptNo"`
	CustomerId *int32  `json:"customerId,omitempty"`
	OrderDate  *string `json:"orderDate,omitempty"`
	OrderNo    *string `json:"orderNo,omitempty"`
	TotalPrice *int32  `json:"totalPrice,omitempty"`
}

// CreateJSONBody defines parameters for Create.
type CreateJSONBody RestCreateCreditPayment

// UpdateJSONBody defines parameters for Update.
type UpdateJSONBody RestUpdateCreditPayment

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody CreateJSONBody

// UpdateJSONRequestBody defines body for Update for application/json ContentType.
type UpdateJSONRequestBody UpdateJSONBody