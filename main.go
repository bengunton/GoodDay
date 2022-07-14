package main

import (
	"fmt"

	"github.com/bengunton/GoodDay/twitter"
)

func main() {
	t := twitter.CreateFetcher("blah")
	fmt.Println(t.GetGoodDay())
}
