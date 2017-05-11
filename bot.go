package xkcdnews

import (
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    "io/ioutil"
    "log"
    "os"
    "path"
)

type Bot struct {
    done   chan bool
    myId   int64
    client *twitter.Client
    log    *log.Logger
}

func New(log *log.Logger) *Bot {
    return &Bot{make(chan bool), 0, nil, log}
}

func (b *Bot) Start() {
    b.log.Println("Starting XKCDNews!")
    consumer, err := ioutil.ReadFile(path.Join(os.Getenv("HOME"), "XKCDNEWS", "CONSUMER_SECRET"))
    if err != nil {
        b.log.Println("Error loading consumer secret: ", err)
        return
    }
    access, err := ioutil.ReadFile(path.Join(os.Getenv("HOME"), "XKCDNEWS", "ACCESS_SECRET"))
    if err != nil {
        b.log.Println("Error loading access secret: ", err)
        return
    }
    b.log.Println("API Keys loaded")

    config := oauth1.NewConfig("iscYDnN0k5thQJUlGi0BLkQDl", string(consumer))
    token := oauth1.NewToken("862666402515255297-rFZpHgQX21sdMUXHKUhnVsN43qUNsXx", string(access))
    httpClient := config.Client(oauth1.NoContext, token)

    b.client = twitter.NewClient(httpClient)

    verify := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }
    user, _, err := b.client.Accounts.VerifyCredentials(verify)
    if err != nil {
        b.log.Println("Error verifying credentials:", err)
        return
    }
    b.log.Println("Logged into ", user.ScreenName)
    b.myId = user.ID

    go b.startTweeting()
}

func (b *Bot) Stop() {
    b.done <- true
}
