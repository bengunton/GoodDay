package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bengunton/GoodDay/models"
	"github.com/bengunton/GoodDay/twitter"
)

type MyEvent struct {
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	return GetContents(), nil
}

func main() {
	lambda.Start(HandleRequest)
}

func GetContents() string {
	t := twitter.CreateFetcher()
	log.Print("Test")

	resp := models.Response{"It's a good day to...\n" + t.GetGoodDay()}
	mapResp, _ := json.Marshal(resp)
	return string(mapResp)
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
