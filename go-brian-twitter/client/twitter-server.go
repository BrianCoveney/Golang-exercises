package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"io/ioutil"
	"strings"
)


func main() {
	getStream()

}

func readKeys() [] string  {
	myKeysFile, err := ioutil.ReadFile("my-keys")
	if err != nil {
		fmt.Println("There was a problem with the twitter api keys")
	}
	return strings.Split(string(myKeysFile), "\n")
}

func auth() *anaconda.TwitterApi {
	myKeys := readKeys()
	consumerKey, consumerSecret := myKeys[0], myKeys[1]
	accessToken, accessSecret := myKeys[2], myKeys[3]

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	return api
}

func getStream() {
	api := auth()

	urlValues := url.Values{}
	urlValues.Set("track", "Brexit")
	twitterStream := api.PublicStreamFilter(urlValues)

	for t := range twitterStream.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			tweetText := v.Text
			screenName := v.User.ScreenName
			fmt.Printf("%-15s: %s\n", screenName, tweetText)

		}

	}

}



