package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	handler "golangdev/zoomAPI/handlers"
	zoomJWT "golangdev/zoomAPI/jwt"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2"
)

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

type Recordings struct {
	Meetings []Meeting `json:"meetings"`
}

type Meeting struct {
	UUID           string      `json:"uuid"`
	Topic          string      `json:"topic"`
	RecordingFiles []Recording `json:"recording_files"`
}

type Recording struct {
	FileType    string `json:"file_type"`
	FileSize    int    `json:"file_size"`
	DownloadUrl string `json:"download_url"`
}

var (
	oauthStateString = "pseudo-random"
	endPoint         = oauth2.Endpoint{
		AuthURL:   "https://zoom.us/oauth/authorize",
		AuthStyle: oauth2.AuthStyleAutoDetect,
		TokenURL:  "https://zoom.us/oauth/token",
	}
	oauthConfig = &oauth2.Config{
		//RedirectURL: "https://abd7feb4151d.ngrok.io/callback",
		RedirectURL:  "https://aec573c036d2.ngrok.io/callback",
		ClientID:     "RagRtzP8TLyIbjzYVGIsFQ",
		ClientSecret: "F65tiTOVh1HWT38Tc8dZZunvKa2PGsK2",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     endPoint,
	}
	client = &http.Client{}

	user_recordings Recordings
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
	// userEmail := "gabdo@intecorp.net" //TODO: change this user email with the one which has the recordings
	getUserURL := "https://api.zoom.us/v2/users/me"

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
	spew.Dump(user1)

	getRecordings(token.AccessToken, user1.ID, user1.CreatedAt)

	fmt.Fprintf(w, " this is the userID %s ", user1.ID)

}

func getRecordings(acess_token string, userID string, userCreatedAt string) {

	const layoutISO = "2006-01-02" //example date

	dateTimeFormat, _ := GetTimeObjFromString(userCreatedAt)
	startDate := dateTimeFormat.Format(layoutISO)
	endDateTime := dateTimeFormat.AddDate(0, 1, -1)
	endDate := endDateTime.Format(layoutISO)

	for dateTimeFormat.Before(time.Now()) {

		url := "https://api.zoom.us/v2/users/" + userID + "/recordings?mc=false&page_size=30&from=" + startDate + "&to=" + endDate
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			fmt.Errorf("get user recordings failed: %s", err.Error())
		}
		authValue := "Bearer " + acess_token
		req.Header.Add("Authorization", authValue)
		res, _ := client.Do(req)
		defer res.Body.Close()
		content, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(content, &user_recordings)
		fmt.Printf("recordings response %s", string(content))

		//get the JWT token
		jwtToken, err := zoomJWT.CreateJWT()
		if err != nil {
			fmt.Errorf("getting JWT failed: %s", err.Error())
		}
		if len(user_recordings.Meetings) > 0 {
			for _, v := range user_recordings.Meetings {
				file_name := v.Topic
				meetingID := v.UUID
				for _, rec := range v.RecordingFiles {
					file_type := rec.FileType
					file_name_with_ext := file_name + "." + strings.ToLower(file_type)
					file_path, err := filepath.Abs(filepath.Join("./Downloads/", file_name_with_ext))
					if err != nil {
						fmt.Errorf("error creating download file path: %s", err.Error())
					}
					go func(recording Recording) {
						fmt.Println("Starting download......", file_path)
						err := handler.DownloadFile(file_path, recording.DownloadUrl, jwtToken, meetingID)
						if err != nil {
							fmt.Errorf("error downloading file: %s", err.Error())
						}
						fmt.Println("DOWNLOADING COMPLETED.........", file_path)
					}(rec)

				}
			}
		}

		//add a month to the current date
		nextDate := dateTimeFormat.AddDate(0, 1, 0) //oneMonthLater
		dateTimeFormat = &nextDate
		startDate = dateTimeFormat.Format(layoutISO)
		endDateTime = dateTimeFormat.AddDate(0, 1, -1)
		endDate = endDateTime.Format(layoutISO)
		// return string(content)

	}

}

func GetTimeObjFromString(timeString string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05Z", timeString)

	if err != nil {
		return nil, err
	}
	return &t, nil
}
