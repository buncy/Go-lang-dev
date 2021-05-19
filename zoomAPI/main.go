package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"golang.org/x/oauth2"
)

type User struct {
	ID string `json:"id"`
}

var (
	oauthStateString = "pseudo-random"
	endPoint         = oauth2.Endpoint{
		AuthURL:   "https://zoom.us/oauth/authorize",
		AuthStyle: oauth2.AuthStyleAutoDetect,
		TokenURL:  "https://zoom.us/oauth/token",
	}
	oauthConfig = &oauth2.Config{
		RedirectURL:  "https://abd7feb4151d.ngrok.io/callback",
		ClientID:     "2jIPrcuUS3iKLtzm3TQRpA",
		ClientSecret: "5LSKtiJrvFW90paOAX6QAdlg60VkPuM3",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     endPoint,
	}
	client = &http.Client{}
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

func handleLogin(w http.ResponseWriter, r *http.Request) {

	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	ctx := context.Background()

	if r.FormValue("state") != oauthStateString {
		fmt.Errorf("invalid oauth state")
	}
	token, err := oauthConfig.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		fmt.Errorf("code exchange failed: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//spew.Dump(token)
	userEmail := "gabdo@intecorp.net" //TODO: change this user email with the one which has the recordings
	getUserURL := "https://api.zoom.us/v2/users/" + userEmail

	//get the user info

	//create a custom request with auth header
	req, _ := http.NewRequest("GET", getUserURL, nil)

	//set auth acess_token in header
	authValue := "Bearer " + token.AccessToken
	req.Header.Add("Authorization", authValue)
	res, err := client.Do(req)

	if err != nil {
		fmt.Errorf("get user failed: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("get user content failed: %s", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	var user1 User
	//parse user info
	json.Unmarshal(content, &user1)

	userRecordings := getRecordings(token.AccessToken, user1.ID)

	fmt.Fprintf(w, " this is the userID %s \n these are the user recordings %s", user1.ID, string(userRecordings))

}

func getRecordings(acess_token string, userID string) string {
	url := "https://api.zoom.us/v2/users/" + userID + "/recordings?trash_type=meeting_recordings&mc=false&page_size=30"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Errorf("get user recordings failed: %s", err.Error())
	}
	authValue := "Bearer " + acess_token
	req.Header.Add("Authorization", authValue)
	res, _ := client.Do(req)
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("recordings response %s", string(content))
	return string(content)

}

// func handleUser(w http.ResponseWriter, r *http.Request)  {

// }
