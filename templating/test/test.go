package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2"
	// "html/template"
	// "context"
)

var (
	//state  = "state=random"
	client = &http.Client{}
	//access_token = ""
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	//http.HandleFunc("/user",handleUser)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Zoom oauth</a></body></html>`
	fmt.Fprintf(w, html)
}

// func handleLogin(w http.ResponseWriter, r *http.Request) {

// 	//url := "https://zoom.us/oauth/authorize?response_type=code&" + state + "&client_id=LTWLsXfKRvmjvA_K1lGnEw&redirect_uri=https%3A%2F%2F2b7161488e9e.ngrok.io%2Fcallback"
// 	url := "https://zoom.us/oauth/authorize?response_type=code&client_id=LTWLsXfKRvmjvA_K1lGnEw&redirect_uri=https%3A%2F%2F2b7161488e9e.ngrok.io%2Fcallback"

// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }
// func handleCallback(w http.ResponseWriter, r *http.Request) {

// 	// if r.FormValue("state") != state {
// 	// 	fmt.Println("state is not valid")
// 	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 	// 	return
// 	// }
// 	url := "https://zoom.us/oauth/token"
// 	content_type := "application/x-www-form-urlencoded"
// 	clientId := "LTWLsXfKRvmjvA_K1lGnEw"
// 	client_secret := "uEPN2thClfUrCe7FUzdH4KkpZXSZ8TgB"
// 	auth := clientId + ":" + client_secret
// 	URLencodindedAuth := b64.URLEncoding.EncodeToString([]byte(auth))
// 	req, _ := http.NewRequest("POST", url, nil)
// 	req.Header.Add("Authorization", URLencodindedAuth)
// 	req.Header.Add("Content-Type", content_type)
// 	query := req.URL.Query()
// 	query.Add("grant_type", "authorization_code")
// 	query.Add("code", r.FormValue("code"))
// 	query.Add("redirect_uri", "https://2b7161488e9e.ngrok.io/callback")
// 	//query.Add("state", state)
// 	req.URL.RawQuery = query.Encode()
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("could not parse content: %v", err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 	}
// 	fmt.Fprintf(w, "Response: %s", res.Body)

// }

// func handleUser(w http.ResponseWriter, r *http.Request)  {
// 	access_token := r.FormValue("access_token")
// 	//refresh_token := r.FormValue("refresh_token")
// 	url := "https://api.zoom.us/v2/users/karthik138105@gmail.com"
// 	req, _ := http.NewRequest("GET", url, nil)
// 	auth := "Bearer "+access_token
// 	req.Header.Add("Authorization",auth)
// 	res,err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("could not create a request %v",err.Error())
// 		http.Redirect(w,r,"/",http.StatusTemporaryRedirect)
// 	}
// 	defer res.Body.Close()
// 	content,err := ioutil.ReadAll(req.Body)
// 	if err!= nil {
// 		fmt.Printf("could not parse content: %v",err.Error())
// 		http.Redirect(w,r,"/",http.StatusTemporaryRedirect)
// 	}
// 	fmt.Fprintf(w,"Response: %s",content)
// }
var oauthStateString = "pseudo-random"

var endPoint = oauth2.Endpoint{
	AuthURL:   "https://zoom.us/oauth/authorize",
	AuthStyle: oauth2.AuthStyleAutoDetect,
	TokenURL:  "https://zoom.us/oauth/token",
}

var oauthConfig = &oauth2.Config{
	RedirectURL:  "https://2b7161488e9e.ngrok.io/callback",
	ClientID:     "LTWLsXfKRvmjvA_K1lGnEw",
	ClientSecret: "uEPN2thClfUrCe7FUzdH4KkpZXSZ8TgB",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     endPoint,
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// fmt.Fprintf(w, ")
}

func getUserInfo(state string, code string) error {
	ctx := context.Background()
	if state != oauthStateString {
		return fmt.Errorf("invalid oauth state")
	}
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return fmt.Errorf("code exchange failed: %s", err.Error())
	}
	spew.Dump(token)
	return nil
	// response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	// }
	// defer response.Body.Close()
	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	// }
	// return contents, nil
}
