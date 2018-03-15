package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/BrianCoveney/stringutil/credentials"
	"github.com/cdipaolo/sentiment"
)

var model sentiment.Models

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", credentials.ConsumerKey, "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", credentials.ConsumerSecret, "Twitter Consumer Secret")
	accessToken := flags.String("access-token", credentials.AccessToken, "Twitter Access Token")
	accessSecret := flags.String("access-secret", credentials.AccessSecret, "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}



	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)

	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	fmt.Printf("User's HOME TIMELINE:\n%+v\n", tweets)

	// Mention Timeline
	mentionTimelineParams := &twitter.MentionTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.MentionTimeline(mentionTimelineParams)
	fmt.Printf("User's MENTION TIMELINE:\n%+v\n", tweets)

	// Retweets of Me Timeline
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, _ = client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)
	fmt.Printf("User's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)

	// Update (POST!) Tweet (uncomment to run)
	// tweet, _, _ := client.Statuses.Update("just setting up go-twitter", nil)
	// fmt.Printf("Posted Tweet\n%v\n", tweet)

	TestPositiveWordSentimentShouldPass1()
	TestNegativeWordSentimentShouldPass1()
}



/*----------------------------------------------------------------------------------------------
      Test run of Cdipaolo's Simple Sentiment Analysis in Golang
      @link - https://github.com/cdipaolo/sentiment
*/


func init() {
	var err error

	//model, err = Train()

	model, err = sentiment.Restore()
	if err != nil {
		panic(err.Error())
	}
}



func TestPositiveWordSentimentShouldPass1() {
	w := []string{"happy", "love", "happiness", "humanity", "awesome", "great", "fun", "super", "trust", "fearless", "creative", "dream", "good", "compassion", "joy", "independent", "success"}
	for _, word := range w {
		s := model.SentimentAnalysis(word, sentiment.English)
		if s.Score != uint8(1) {
			fmt.Print("Sentiment of < %v > (returned %v) should be greater than 0.5!\n", word, s)
		} else {
			fmt.Print("Sentiment of < %v > valid\n\tReturned %v\n", word, s)
		}
	}
}


func TestNegativeWordSentimentShouldPass1() {
	w := []string{"not", "resent", "deplorable", "bad", "terrible", "hate", "scary", "terrible", "concerned", "wrong", "rude!!!", "sad", "horrible", "unimpressed", "useless", "offended", "disrespectful"}
	for _, word := range w {
		s := model.SentimentAnalysis(word, sentiment.English)
		if s.Score != uint8(0) {
			fmt.Print("Sentiment of < %v > (returned %v) should be less than 0.5!\n", word, s)
		} else {
			fmt.Print("Sentiment of < %v > valid\n\tReturned %v\n", word, s)
		}
	}
}


