package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string, access_token string, meeting_id string) error {

	// Get the data
	// optional_url := "https://api.zoom.us/v2/meetings/" + meeting_id + "/recordings"

	req, err := http.NewRequest("GET", url, nil) //FIXME: if this url dosen't work use the optional_url and try
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
