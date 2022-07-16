package twitter

type twitterResponse struct {
	Data []tweet
	Meta responseMeta
}

type tweet struct {
	Created_at string
	Id         string
	Text       string
}

type responseMeta struct {
	Newest_id    string
	Oldest_id    string
	Result_count int
	Next_token   string
}
