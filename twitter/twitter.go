package twitter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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

	resp, err := t.makeRequest()
	if err != nil {
		log.Print(err)
		return "failed"
	}

	tweets := parseTweets(&resp.Body)

	reply := "got " + strconv.Itoa(len(tweets)) + " tweets"
	return reply
}

func (t *TweetFetcher) makeRequest() (*http.Response, error) {
	client := &http.Client{}

	url := `https://api.twitter.com/2/tweets/search/recent?query="it's%20a%20good%20day%20to"&tweet.fields=created_at&max_results=10`
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Print("Failed to create request")
		return nil, err
	}

	req.Header.Add("Authorization", t.bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Failed to send request")
		return nil, err
	}

	return resp, nil
}

func parseTweets(body *io.ReadCloser) []tweet {
	decoder := json.NewDecoder(*body)
	decoder.DisallowUnknownFields()

	var response twitterResponse

	err := decoder.Decode(&response)
	if err != nil {
		log.Print(err)
		return nil
	}

	tweets := make([]tweet, len(response.Data))
	copy(tweets, response.Data)
	return tweets
}
