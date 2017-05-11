package xkcdnews

import (
    "github.com/dghubble/go-twitter/twitter"
    "time"
)

func (b *Bot) startTweeting() {
    ticker := time.NewTicker(30 * time.Minute)

    for {
        select {
        case <-ticker.C:
            b.postNewTweet()
        case <-b.done:
            ticker.Stop()
            return
        }
    }
}

func (b *Bot) postNewTweet() {
    utlParams := &twitter.UserTimelineParams{
        UserID: b.myId,
        Count:  1,
    }
    lastTweet, _, err := b.client.Timelines.UserTimeline(utlParams)
    if err != nil {
        b.log.Println("Error fetching last tweet:", err)
        return
    }
    lastTime, err := time.Parse(time.UnixDate, lastTweet[0].CreatedAt)
    if err != nil {
        b.log.Println("Error formatting time:", err)
        return
    }

    if lastTime.Add(time.Hour).After(time.Now()) {
        // if it hasn't been an hour, retry later
        return
    }

    tlParams := &twitter.HomeTimelineParams{
        Count:           20,
        TrimUser:        twitter.Bool(true),
        ExcludeReplies:  twitter.Bool(true),
        IncludeEntities: twitter.Bool(false),
    }
    tweets, _, err := b.client.Timelines.HomeTimeline(tlParams)
    if err != nil {
        b.log.Println("Error fetching timeline:", err)
        return
    }

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

    if bestTweetRating > 0 {
        params := &twitter.StatusUpdateParams{
            TrimUser: twitter.Bool(true),
        }
        _,_, err := b.client.Statuses.Update(bestTweet, params)
        if err != nil {
            b.log.Println("Error posting tweet:", err)
        }
    }
}
