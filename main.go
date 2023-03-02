package main

import (
	"fanchann/Go-APIWithArchitecture2/api"
	"fanchann/Go-APIWithArchitecture2/api/middleware"
	"net/http"
)

func main() {
	// sha256 := sha256.New()
	// sha256.Write([]byte("fardaayu@gmail.com"))

	// hashByte := sha256.Sum(nil)

	// hashResult := hex.EncodeToString(hashByte)

	// fmt.Println(hashResult)

	// fmt.Println(time.Now())

	router := api.NewRouter()
	server := http.Server{
		Addr:    ":3000",
		Handler: &middleware.UsersMiddleware{&router},
	}

	server.ListenAndServe()
}
