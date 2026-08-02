package main

import "fmt"

type Subscriber interface {
	update(string)
}

type EmailSubscriber struct {
	username string
}

func (e *EmailSubscriber) update(title string) {
	fmt.Printf("Updated email user %v with %v\n", e.username, title)
}

type MobileAppSubscriber struct {
	username string
}

func (m *MobileAppSubscriber) update(title string) {
	fmt.Printf("Updated mobile app user %v with %v\n", m.username, title)
}

type Channel interface {
	subscribe(Subscriber)
	unsubscribe(Subscriber)
	notifySubscriber(string)
	uploadVideo(string)
}

type YoutubeChannel struct {
	subscribers []Subscriber
	name        string
}

func (y *YoutubeChannel) subscribe(sub Subscriber) { // fix 2
	y.subscribers = append(y.subscribers, sub)
	fmt.Printf("New subscriber on %s\n", y.name) // fix 4: removed sub.username
}
func (y *YoutubeChannel) unsubscribe(sub Subscriber) { // fix 3
	for i, v := range y.subscribers {
		if v == sub {
			y.subscribers = append(y.subscribers[:i], y.subscribers[i+1:]...)
			break
		}
	}
	for _, s := range y.subscribers {
		fmt.Printf("  %+v\n", s)
	}
}

func (y *YoutubeChannel) notifySubscriber(msg string) {
	for _, val := range y.subscribers {
		val.update(msg)
	}
}

func (y *YoutubeChannel) uploadVideo(title string) {
	fmt.Printf("%s has been uploaded!!!\n", title)
	y.notifySubscriber("My new video is live!!")
}

func main() {
	emailSub := &EmailSubscriber{
		username: "teja",
	}
	mobileAppSub := &MobileAppSubscriber{
		username: "wini",
	}

	yt := &YoutubeChannel{
		name: "Teju YT Channel",
	}
	yt.subscribe(emailSub)
	yt.subscribe(mobileAppSub)
	yt.uploadVideo("What i eat in a day, day 1")
	yt.unsubscribe(mobileAppSub)
}
