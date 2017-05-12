package xkcdnews

import (
	"github.com/dghubble/go-twitter/twitter"
	"time"
)

// Starts the timer to continuously tweet
func (b *Bot) startTweeting() {
	for {
		tryAgain := time.After(b.tryPostNewTweet())
		select {
		case <-tryAgain:
			continue
		case <-b.done:
			tryAgain = nil
			return
		}
	}
}

// Makes sure that there hasn't been a new tweet
// in the past hour before tweeting.
//
// Return: The duration until this should be called again
func (b *Bot) tryPostNewTweet() time.Duration {

	// Get our most recent tweet
	utlParams := &twitter.UserTimelineParams{
		UserID: b.myId,
		Count:  1,
	}
	lastTweet, _, err := b.client.Timelines.UserTimeline(utlParams)
	if err != nil {
		b.log.Println("Error fetching last tweet:", err)
		return 5 * time.Minute
	}

	// Get the time the tweet was posted
	lastTime, err := time.Parse(time.RubyDate, lastTweet[0].CreatedAt)
	if err != nil {
		// A time format error means an error on my part most likely,
		// so make the duration longer
		b.log.Println("Error formatting time:", err)
		return 5 * time.Hour
	}

	b.log.Println("Last tweet", lastTime.Format(time.RubyDate))

	// if it hasn't been an hour since our last tweet, retry when the hour is up
	if lastTime.Add(time.Hour).After(time.Now()) {
		return time.Now().Sub(lastTime) + time.Minute
	}

	return b.postNewTweet()
}

// Posts a new tweet.
//
// Return: The duration until tryPostNewTweet() should be called
func (b *Bot) postNewTweet() time.Duration {
	// Get the last 20 posts
	tlParams := &twitter.HomeTimelineParams{
		Count:           20,
		TrimUser:        twitter.Bool(true),
		ExcludeReplies:  twitter.Bool(true),
		IncludeEntities: twitter.Bool(false),
	}
	tweets, _, err := b.client.Timelines.HomeTimeline(tlParams)
	if err != nil {
		b.log.Println("Error fetching timeline:", err)
		return 5 * time.Minute
	}

	// Look through the tweets, replace words,
	// and pick the tweet with the most replaced words
	var bestTweet string
	bestTweetRating := 0
	for _, tweet := range tweets {
		newTweet, replaced := Substitute(tweet.Text)
		if len(newTweet) > 140 {
			continue
		}
		if replaced > bestTweetRating {
			bestTweet = newTweet
			bestTweetRating = replaced
		}
	}

	// Post the tweet!
	if bestTweetRating > 0 {
		params := &twitter.StatusUpdateParams{
			TrimUser: twitter.Bool(true),
		}
		_, _, err := b.client.Statuses.Update(bestTweet, params)
		if err != nil {
			b.log.Println("Error posting tweet:", err)
			return 5 * time.Minute
		}
		return 1 * time.Hour
	}
	return 5 * time.Minute
}
