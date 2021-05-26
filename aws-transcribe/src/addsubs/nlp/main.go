package main

import (
	"fmt"
	"os"

	"github.com/jdkato/prose/v2"
)

func main() {

	doc, _ := prose.NewDocument("We are down to the last day of our stay at Lake Tekapo and as we had hoped the dip in the hot tub on the deck at night has made our stay here very relaxing. We are pinching ourselves because we had found such a restful accommodation.There had been a bit more rain overnight but the prospects for a reasonable day without any further rain was on the cards which will ensure we get around the three close locations to the township that we want to explore and get a feeling for the history of this unique part of New Zealand.Heading out for some sightseeing our first location to visit were the twin lakes of Lake McGregor and Lake Alexandrina to the northwest of Tekapo township.From Tekapo you can’t see these lakes as they are hidden behind the imposing hill that the Mt John Observatory sits on. We hadn’t included a visit to the observatory as the nights had not been clear enough to take in what is reputed to be the best night sky area in New Zealand.")
	f, err := os.Create("/home/karthik/go/src/golangdev/aws-transcribe/src/addsubs/nlp/result.txt")
	if err != nil {
		fmt.Printf("error creating file %v", err)
		return
	}
	number := 1
	for _, ent := range doc.Tokens() {
		entry := fmt.Sprintf("%d) %v %v %v\n\n", number, ent.Text, ent.Tag, ent.Label)
		f.WriteString(entry)
		//fmt.Println(ent.Text, ent.Tag)
		number++
		// Lebron James PERSON
		// Los Angeles GPE
	}
	f.Close()
}
