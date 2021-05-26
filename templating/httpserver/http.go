package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.zoom.us/v2/users/JRnAKNeqRvyGGZlUVuDqNw/meetings?page_size=30&type=live"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer eyJhbGciOiJIUzUxMiIsInYiOiIyLjAiLCJraWQiOiI2NmQ1NTZkZS0zYWFlLTQ0ZTctOTBkZS1hNjk1ZDRlZTI2ZjkifQ.eyJ2ZXIiOjcsImF1aWQiOiI2ZjQ1ZDAxYjk0ZDE5NTc5ZThiYmUzNmM2NGFlNzllYiIsImNvZGUiOiJqa2dkc2VEcVh3X0pSbkFLTmVxUnZ5R0dabFVWdURxTnciLCJpc3MiOiJ6bTpjaWQ6TFRXTHNYZktSdm1qdkFfSzFsR25FdyIsImdubyI6MCwidHlwZSI6MCwidGlkIjowLCJhdWQiOiJodHRwczovL29hdXRoLnpvb20udXMiLCJ1aWQiOiJKUm5BS05lcVJ2eUdHWmxVVnVEcU53IiwibmJmIjoxNjIxMjU5NzQzLCJleHAiOjE2MjEyNjMzNDMsImlhdCI6MTYyMTI1OTc0MywiYWlkIjoiMXgyN0FPYlRTWUdNb2x2SW1Kb19VQSIsImp0aSI6ImIwMWIwNTc4LTNlYjUtNDlmZi05YjY1LThjNTczZDRiMzFmOCJ9.OAscEg6FClPQFBW6WwOiduUP3MOMNPaZieinpaHXrzP-ItwhW2yvQfRxbZLyfOyI3eI-e85-86hDPAihQ6320Q")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
