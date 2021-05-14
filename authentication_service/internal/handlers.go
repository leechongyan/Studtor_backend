package internal

import (
	"fmt"
	"strconv"
	"github.com/spf13/viper"
	
	"net/http"
	"github.com/gorilla/sessions"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func InitGothic(){
	key := viper.GetString("oauthStateString")  // Replace with your SESSION_SECRET or similar
	maxAge, _ := strconv.Atoi(viper.GetString("gothic.maxAge")) 
	maxAge = maxAge * 86400
	isProd, _ := strconv.ParseBool(viper.GetString("gothic.isProd")) // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly, _ = strconv.ParseBool(viper.GetString("gothic.httpOnly"))   // HttpOnly should always be enabled
	store.Options.Secure = isProd
	gothic.Store = store
	goth.UseProviders(
		google.New(viper.GetString("google.clientID"), viper.GetString("google.clientSecret"), viper.GetString("google.callbackURL"), "email", "profile"),
	)
}


func GothicLoginHandler(res http.ResponseWriter, req *http.Request){
	// try to get the user without re-authenticating
	if user, err := gothic.CompleteUserAuth(res, req); err == nil {
		fmt.Fprintf(res, "%s", user)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func GothicCallbackHandler(res http.ResponseWriter, req *http.Request){
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	
	fmt.Fprintf(res, "%s", user)
}

func GothicLogoutHandler(res http.ResponseWriter, req *http.Request){
	gothic.Logout(res, req)
	fmt.Fprintf(res, "Logout")
}