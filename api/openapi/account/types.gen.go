// Package account provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package account

// Account defines model for Account.
type Account struct {
	Email *string `json:"email,omitempty"`
	Id    *string `json:"id,omitempty"`
}
