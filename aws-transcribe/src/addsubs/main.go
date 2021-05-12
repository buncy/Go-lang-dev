package main

import (
	"bytes"
	"flag"
	"fmt"
	helper "golangdev/aws-transcribe/src/addsubs/helpers"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	inputFile := flag.String("inputFilePath", "./media/hindi.mp4", "*required field, The input mp4 file without subtitles")
	jsonFile := flag.String("jsonFile", "./media/hindi.json", "*required field, The aws-transcribe json file")
	srtFile := flag.String("srtFile", "./media/hindiscript.srt", "*required field, The name of the srt file")
	outputFile := flag.String("ouputFile", "./media/hindisrt.mp4", "*required field, The name of the output mp4 file")
	flag.Parse()

	srtpath, _ := filepath.Abs(*srtFile)
	inputpath, _ := filepath.Abs(*inputFile)
	outputpath, _ := filepath.Abs(*outputFile)
	jsonpath, _ := filepath.Abs(*jsonFile)
	srtContent, cerr := helper.Convert(jsonpath)

	if cerr != nil {
		fmt.Println(cerr)
	}

	f, err := os.Create(*srtFile)
	if err != nil {
		fmt.Printf("while opening file %v", err)
	}
	defer f.Close()

	_, werr := f.WriteString(srtContent)
	if err != nil {
		fmt.Printf("this is while writing %v", werr)
	}

	f.Sync()

	//command := fmt.Sprintf(`ffmpeg -i %s -vf "subtitles=%s:force_style='Borderstyle=4,Fontsize=16,BackColour=&H80000000'" %s `, inputpath, srtpath, outputpath)
	// args := strings.Split(command, " ")
	// cmd := exec.Command(args[0], args[1:]...)

	// subtitle := "subtitles=" + (srtpath) + ":force_style='Borderstyle=4,Fontsize=16,BackColour=&H80000000'"
	// cmd := exec.Command("ffmpeg", "-i", inputpath, "-vf", subtitle, outputpath)
	// cmd := exec.Command("ffmpeg", "-i", inputpath, "-f", "srt", "-i", srtpath, "-c:v", "copy", "-c:a", "copy", outputpath)
	// cmd.CombinedOutput()
	subtitle := "subtitles=" + srtpath + ":force_style='Fontsize=16,PrimaryColour=&H00FFFFFF,SecondaryColour=&H0300FFFF'"

	cmd := exec.Command("ffmpeg", "-i", inputpath, "-vf", subtitle, outputpath)
	var out bytes.Buffer
	if cmd != nil {
		//cmd.Stderr = &out
		err := cmd.Run()

		if err != nil {
			fmt.Println(out.String(), " <<<------ error is here while creating final video: ", cmd.String())
			return
		} else {
			fmt.Println("Here we go.. success!")
		}
	}
}
