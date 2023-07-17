package main

import (
	"fmt"

	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func main() {
	mytext := "I love apples and coding so much"
	parsedtext := sentitext.Parse(mytext, lexicon.DefaultLexicon)
	sentiment := sentitext.PolarityScore(parsedtext)
	fmt.Println("Pos:", sentiment.Positive)
	fmt.Println("Neg:", sentiment.Negative)
	fmt.Println("Neu:", sentiment.Neutral)
	fmt.Println("Compound/Final Sentiment:", sentiment.Compound)
}
