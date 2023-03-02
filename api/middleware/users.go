package middleware

import (
	"encoding/json"
	"fanchann/Go-APIWithArchitecture2/internals/models/web"
	"net/http"
)

type UsersMiddleware struct {
	Handler http.Handler
}

func UserAuthMiddleware(handler http.Handler) *UsersMiddleware {
	return &UsersMiddleware{Handler: handler}
}

func (middl *UsersMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "supers3cret" == request.Header.Get("API-KEY") {
		middl.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")

		failedResponse := web.WebResponse{
			Status:  http.StatusUnauthorized,
			Message: "Who?",
			Data:    nil,
		}

		encoder := json.NewEncoder(writer)
		encoder.Encode(failedResponse)

	}
}
