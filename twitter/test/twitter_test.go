package test

import (
	"fmt"
	"testing"

	"github.com/hauntedness/twitterbot/twitter"
)

func TestTwitter_SearchTweet(t *testing.T) {
	tweets := twitter.SearchTweet("airdrop")
	for i, d := range tweets.Data {
		fmt.Println(i)
		fmt.Println(d)
	}
}
