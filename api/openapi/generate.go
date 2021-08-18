package openapi

//go:generate oapi-codegen -generate types -package account -o ./account/types.gen.go account.yml
//go:generate oapi-codegen -generate client -package account -o ./account/client.gen.go account.yml
//go:generate oapi-codegen -generate chi-server -package account -o ./account/server.gen.go account.yml
//go:generate oapi-codegen -generate spec -package account -o ./account/spec.go account.yml

//go:generate oapi-codegen -generate types -package item -o ./item/types.gen.go item.yml
//go:generate oapi-codegen -generate client -package item -o ./item/client.gen.go item.yml
//go:generate oapi-codegen -generate chi-server -package item -o ./item/server.gen.go item.yml
//go:generate oapi-codegen -generate spec -package item -o ./item/spec.go item.yml
