package main

import (
	"flag"
	"fmt"

	scago "github.com/myriolang/scago"
)

func main() {
	inputFile := flag.String("i", "", "file containing a list of input words to be changed")
	outputFile := flag.String("o", "", "filename for the output of the sound changes")
	rulesetFile := flag.String("f", "", "file containing a list of rules to be applied to all words")
	ruleLiteral := flag.String("r", "", "a single rule to apply to the word(s)")
	flag.Parse()
	inputLiteral := flag.Arg(0)

	s := scago.New()

	if *rulesetFile != "" {
		// TODO: open file, add rules
		fmt.Println("[DEBUG] This feature hasn't been added yet. Sorry!")
		return
	} else if *ruleLiteral != "" {
		s.AddRule(*ruleLiteral)
	} else {
		fmt.Println("No rule(s) specified.")
		return
	}

	if *inputFile != "" {

	} else if inputLiteral != "" {

	}
}
