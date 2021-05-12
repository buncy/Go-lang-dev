package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func Convert(localFilePath string) (string, error) {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	absPath, _ := filepath.Abs(localFilePath)
	body, err := ioutil.ReadFile(absPath)
	if err != nil {

		spew.Println("error")
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	// text := string(content)
	// fmt.Println(text)

	//If there's an error, print the error
	if err != nil {
		fmt.Println(err)
	}

	// initialize our variable to hold the json
	var awstranscript Awstranscript

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'awstranscript' which we defined above
	json.Unmarshal(body, &awstranscript)
	fmt.Println("starting srt ")
	var transcription []Item
	transcription = awstranscript.Results.Items
	//transcriptionLength := len(transcription)
	var index, number, sequence int = 0, 1, 1 //FIXME:changed from 0,0
	var starttime = ""
	var subdetail, subtitle, endtime string

	for index = 0; index < len(transcription); {
		//Variable initiation for length of subtitle text, sequence number if its the first line and the subtitle text

		// sequence++

		// subtitle = ""
		//fmt.Println("started srt loop")

		if transcription[index].Classification == "pronunciation" {
			if starttime == "" {
				//Grab the start time, convert it to a number, then convert the number an SRT valid time string
				starttime = transcription[index].Starttime
				fstarttime, err := strconv.ParseFloat(starttime, 64)
				if err != nil {
					fmt.Println(err, "================", starttime)
				}
				starttime = Getsrttime(fstarttime)
			}

			//set endtime

			endtime = transcription[index].Endtime
			fendtime, err := strconv.ParseFloat(endtime, 64)
			if err != nil {
				fmt.Println(err, "================", endtime)
			}
			endtime = Getsrttime(fendtime)

			subtitle += transcription[index].Alternatives[0].Content + " "
			sequence++
		}
		if transcription[index].Classification == "punctuation" && transcription[index].Alternatives[0].Content == "." {
			//subtitle += transcription[index].Alternatives[0].Content
			subdetail += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", sequence, starttime, endtime, subtitle)
			subtitle = ""
			starttime = ""
			number++
			sequence = 1
		}

		index++
	}

	log.Printf(subdetail)

	return subdetail, nil
}

// Getsrttime - Generates an SRT format time string
func Getsrttime(numerator float64) (timestring string) {

	var h = 3600
	var m = 60
	var s = 1

	integer, frac := math.Modf(numerator)
	integerpart := int(integer)

	hours := integerpart / h
	remainder := integerpart % h

	minutes := remainder / m
	remainder = remainder % m

	seconds := remainder / s
	stringfrac := strconv.FormatFloat(frac, 'f', 3, 64)
	runes := []rune(stringfrac)
	safeSubstring := string(runes[1:len(stringfrac)])

	timestring = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	timestring += safeSubstring
	return
}

//Awstranscript - Top level struc for an AWS transcript job output
type Awstranscript struct {
	JobName   string `json:"jobName"`
	Accountid string `json:"accountId"`
	Results   Result `json:"results"`
	Status    string `json:"status"`
}

//Result - Result structure
type Result struct {
	Transcripts []Transcript `json:"transcripts"`
	Items       []Item       `json:"items"`
}

//Transcript - Transcription
type Transcript struct {
	Transcript string `json:"transcript"`
}

//Item - Individual translation word/punctuation record
type Item struct {
	Starttime      string        `json:"start_time"`
	Endtime        string        `json:"end_time"`
	Alternatives   []Alternative `json:"alternatives"`
	Classification string        `json:"type"`
}

// Alternative - Actual translated word and confidence of accuracy
type Alternative struct {
	Confidence string `json:"confidence"`
	Content    string `json:"content"`
}
