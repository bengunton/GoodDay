package twitter

import (
	"os"

	"github.com/bengunton/GoodDay/models"
)

type TweetFetcher struct {
	apiKey string
}

func CreateFetcher() models.Fetcher {
	key := os.Getenv("BearerToken")
	return TweetFetcher{apiKey: key}
}

func (t TweetFetcher) GetGoodDay() string {
	return t.apiKey
}
