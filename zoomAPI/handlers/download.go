package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string, access_token string, meeting_id string) error {

	// Get the data
	// optional_url := "https://api.zoom.us/v2/meetings/" + meeting_id + "/recordings"

	urlWithToken := url + "?access_token=" + access_token
	fmt.Println(urlWithToken)
	req, err := http.NewRequest("GET", urlWithToken, nil) //FIXME: if this url dosen't work use the optional_url and try
	if err != nil {
		return err
	}
	// authValue := "Bearer " + access_token
	// req.Header.Add("Authorization", authValue)

	resp, err := http.DefaultClient.Do(req)
	fmt.Println("Status code for "+filepath+" is ", resp.StatusCode, resp.Status)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check if file exists
	var _, statErr = os.Stat(filepath)

	// create file if not exists
	if os.IsNotExist(statErr) {
		CreateDir(filepath)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func CreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println(" -- error creating " + dir)
			return err
		}
	}
	return nil
}
