package twitter

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bengunton/GoodDay/models"
)

type TweetFetcher struct {
	bearerToken string
}

func CreateFetcher() models.Fetcher {
	key := os.Getenv("BearerToken")
	return TweetFetcher{bearerToken: "Bearer " + key}
}

func (t TweetFetcher) GetGoodDay() string {
	client := &http.Client{}

	url := `https://api.twitter.com/2/tweets/search/recent?query="it's%20a%20good%20day%20to"&tweet.fields=created_at&max_results=10`
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Print(err)
		return nil
	}

	req.Header.Add("Authorization", t.bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil
	}

	return string(body)
}
