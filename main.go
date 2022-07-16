package main

import (
	"fmt"

	"github.com/bengunton/GoodDay/twitter"
)

func main() {
	t := twitter.CreateFetcher()
	fmt.Println("It's a good day to...")
	fmt.Println(t.GetGoodDay())
}
