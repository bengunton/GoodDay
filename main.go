package main

import (
	"fmt"

	"github.com/bengunton/GoodDay/twitter"
)

func main() {
	t := twitter.CreateFetcher()
	fmt.Println(t.GetGoodDay())
}
