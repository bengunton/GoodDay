package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bengunton/GoodDay/twitter"
)

func main() {
	t := twitter.CreateFetcher()
	fmt.Println("It's a good day to...")

	goodDay := t.GetGoodDay()
	fmt.Println(goodDay)

	WriteToFile(goodDay)
}

func WriteToFile(contents string) {
	f, err := os.Create("./output/output.txt")
	if err != nil {
		log.Print(err)
	}

	_, err = f.WriteString(contents + "\n")
	if err != nil {
		log.Print(err)
	}
}
