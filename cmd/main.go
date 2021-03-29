package main

import (
	"flag"
	"fmt"

	scago "github.com/myriolang/scago"
)

func main() {
	//inputFile := flag.String("i", "", "file containing a list of input words to be changed")
	//outputFile := flag.String("o", "", "filename for the output of the sound changes")
	//rulesetFile := flag.String("f", "", "file containing a list of rules to be applied to all words")
	ruleLiteral := flag.String("r", "", "a single rule to apply to the word(s)")
	flag.Parse()
	inputLiteral := flag.Arg(0)

	s := scago.New()

	// TODO: allow reading of ruleset file for adding rules to ScagoInstance
	/*if *rulesetFile != "" {
		fmt.Println("[DEBUG] This feature hasn't been added yet. Sorry!")
		return
	} else*/
	if *ruleLiteral != "" {
		s.AddRule(*ruleLiteral)
	} else {
		fmt.Println("No rule(s) specified.")
		return
	}

	// TODO: allow reading of input file for processing words to Apply
	/*if *inputFile != "" {
		fmt.Println("[DEBUG] This feature hasn't been added yet. Sorry!")
		return
	} else if inputLiteral != "" {*/
	output, err := s.Apply(inputLiteral)
	//}

	if err != nil {
		fmt.Println(output)
	} else {
		fmt.Println("Something went wrong!", err)
	}
}
