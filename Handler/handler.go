package handler

import (
	"encoding/json"
	"fmt"
	"gobasics/utils"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOAuthConfig = &oauth2.Config{
		RedirectURL:  "",
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	state = strconv.FormatInt(rand.Int63(), 10)
)

func HomeHandler(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./template/login.html")
	if err != nil {
		log.Fatal("Error loading file", err)
	}
	t.Execute(writer, false)
}

func AnotherHandler(writer http.ResponseWriter, req *http.Request) {
	url, err := mux.CurrentRoute(req).Subrouter().Get("home").URL()
	if err != nil {
		panic(err)
	}
	http.Redirect(writer, req, url.String(), 302)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	get := utils.GetEnvWithKey

	googleOAuthConfig.RedirectURL = get("RedirectURL")
	googleOAuthConfig.ClientID = get("ClientID")
	googleOAuthConfig.ClientSecret = get("ClientSecret")

	url := googleOAuthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

type GoogleOAuthResponse struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Language      string `json:"locale"`
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	urlState := params.Get("state")

	if urlState != state {
		fmt.Println("state not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOAuthConfig.Exchange(oauth2.NoContext, params.Get("code"))
	if err != nil {
		fmt.Printf("could not get token : %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("Access Token 		----->	 	", token.AccessToken)
	fmt.Println("Expiration Time 	----->	 	", token.Expiry.String())
	fmt.Println("Refresh Token 		-----> 		", token.RefreshToken)

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could get url : %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer res.Body.Close()

	var googleRes GoogleOAuthResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&googleRes)

	if err != nil {
		fmt.Printf("Failed to decode response : %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	t, err := template.ParseFiles("./template/success.html")
	if err != nil {
		log.Fatal("Error loading file", err)
	}
	t.Execute(w, googleRes)
}
