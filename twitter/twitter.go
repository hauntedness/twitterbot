package twitter

import (
	"encoding/json"
	"net/http"

	"github.com/hauntedness/httputil"
	"github.com/hauntedness/twitterbot/config"
	"github.com/hauntedness/twitterbot/twitter/bean"
)

//SearchTweet
//	search twetter by key words
//	return a Tweet slice
func SearchTweet(keyWords string) bean.SearchTweetResult {
	url := "https://api.twitter.com/2/tweets/search/recent?query=" + keyWords
	bc := config.GetBotConfig()
	httputil.SetProxy("")
	headers := httputil.H{
		"Authorization": bc.Token,
	}
	res := httputil.Request(http.MethodGet, url, nil, headers)
	tweets := bean.SearchTweetResult{}
	json.Unmarshal(res, &tweets)
	return tweets
}
