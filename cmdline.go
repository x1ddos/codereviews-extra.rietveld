package main

import (
	"fmt"
	"log"
	"code.google.com/p/codereviews-extra.rietveld/rietveld"
)

func main() {
	list, err := rietveld.Search(nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(list)
	for _, issue := range list.Issues {
		fmt.Println(issue)
	}
}
