package bean

type SearchTweetResult struct {
	Data []Datum `json:"data"`
	Meta Meta    `json:"meta"`
}

type Datum struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Meta struct {
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	ResultCount int64  `json:"result_count"`
	NextToken   string `json:"next_token"`
}
