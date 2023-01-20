package twitter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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
		log.Println(err)
		return "failed"
	}

	tweets := parseTweets(&resp.Body)

	for i := 0; i < len(tweets); i++ {
		goodDay, success := tweets[i].getGoodDay()
		if success {
			return goodDay
		}
	}

	return "fail to find any tweets :("
}

func (t *TweetFetcher) makeRequest() (*http.Response, error) {
	client := &http.Client{}

	url := `https://api.twitter.com/2/tweets/search/recent?query="it's%20a%20good%20day%20to"-is:retweet&tweet.fields=created_at`
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Failed to create request")
		return nil, err
	}

	req.Header.Add("Authorization", t.bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send request")
		return nil, err
	}

	return resp, nil
}

func parseTweets(body *io.ReadCloser) []tweet {
	decoder := json.NewDecoder(*body)
	// decoder.DisallowUnknownFields()

	var response twitterResponse

	err := decoder.Decode(&response)
	if err != nil {
		log.Println(err)
		log.Println(response)
		return nil
	}

	tweets := make([]tweet, len(response.Data))
	copy(tweets, response.Data)
	return tweets
}

func (t tweet) getGoodDay() (string, bool) {
	content := t.Text

	_, after, found := strings.Cut(content, "It’s a good day to ")
	if !found {
		_, after, found = strings.Cut(content, "It's a good day to ")
	}

	return after, found
}
