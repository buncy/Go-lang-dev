package main

import (
	"context"
	"fmt"

	"github.com/hironobu-s/go-corenlp" // exposes "corenlp"
	"github.com/hironobu-s/go-corenlp/connector"
)

func main() {
	// sample text from https://stanfordnlp.github.io/CoreNLP/
	text := `President Xi Jinping of Chaina, on his first state visit to the United States, showed off his familiarity with American history and pop culture on Tuesday night.`

	// LocalExec connector is responsible to run Stanford CoreNLP process.
	c := connector.NewLocalExec(context.TODO())
	c.ClassPath = "./corenlp/*" // set Java class path
	c.Annotators = []string{"tokenize", "ssplit", "pos"}

	// Annotate text
	doc, err := corenlp.Annotate(c, text)
	if err != nil {
		panic(err)
	}

	// Output words and pos
	for _, sentence := range doc.Sentences {
		for _, token := range sentence.Tokens {
			fmt.Printf("%s(%s)%s\n\n", token.Word, token.Pos, token.After)
		}
	}
}
