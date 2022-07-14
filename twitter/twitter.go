package twitter

import (
	"github.com/bengunton/GoodDay/models"
)

type TweetFetcher struct {
	apiKey string
}

func CreateFetcher(key string) models.Fetcher {
	return TweetFetcher{apiKey: key}
}

func (t TweetFetcher) GetGoodDay() string {
	return t.apiKey
}
