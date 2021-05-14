package main

import (
	"github.com/gorilla/mux"
	"net/http"
	goth "github.com/leechongyan/Studtor_backend/authentication_service/internal"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// init gothic 
	goth.InitializeViper()
	goth.InitGothic()

	router.HandleFunc("/auth/{provider}/callback", goth.GothicCallbackHandler).Methods(http.MethodGet)
	router.HandleFunc("/auth/{provider}", goth.GothicLoginHandler).Methods(http.MethodGet)
	router.HandleFunc("/logout/{provider}", goth.GothicLogoutHandler).Methods(http.MethodGet)

	http.ListenAndServe(":3000", router)
}