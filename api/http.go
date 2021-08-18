package api

import (
	"avgarde/api/openapi/account"
	"avgarde/common/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type Mock struct {
	
}

func StrPtr(value string) *string {
	return &value
}

func (*Mock) CurrentAccount(w http.ResponseWriter, r *http.Request)  {
	render.Respond(w, r, account.Account{
		Email: StrPtr("mock@gmail.com"),
		Id:    StrPtr("1"),
	})
}

func Server() {
	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return account.HandlerFromMux(
			&Mock{},
			router,
		)
	})
}
