package main

import (
	"C4M2H1/domain"
	"fmt"
	"log"
	"os"
)

func main() {

	content, err := os.ReadFile("script.txt")
	if err != nil {
		log.Fatalln(err)
	}

	adaptee := domain.NewSuperRelationshipAnalyzer()
	a := domain.NewSuperRelationshipAnalyzerAdapter(adaptee)
	relationshipGraph := a.Parse(string(content))
	//fmt.Println(a.GetMutualFriends("B", "C"))
	fmt.Println(relationshipGraph.HasConnection("B", "C"))
}
