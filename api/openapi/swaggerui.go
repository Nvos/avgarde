package openapi

import (
	"avgarde/api/openapi/account"
	"avgarde/api/openapi/item"
	"net/http"
)

func AccountHandler() http.HandlerFunc {
	swagger, err := account.GetSwagger()
	if err != nil {
		return nil
	}

	json, err := swagger.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return byteHandler(json)
}

func ItemHandler() http.HandlerFunc {
	swagger, err := item.GetSwagger()
	if err != nil {
		return nil
	}

	json, err := swagger.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return byteHandler(json)
}


func byteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write(b)
	}
}