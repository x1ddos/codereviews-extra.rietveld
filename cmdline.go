package main

import (
	"fmt"
	"log"
	"./rietveld"
)

func main() {
	list, err := rietveld.Search()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(list)
	for _, issue := range list.Issues {
		fmt.Println(issue)
	}
}
