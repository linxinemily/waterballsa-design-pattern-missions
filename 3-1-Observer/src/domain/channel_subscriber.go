package domain

import "fmt"

type ChannelSubscriber struct {
	name string
}

func NewChannelSubscriber(name string) *ChannelSubscriber {
	return &ChannelSubscriber{
		name: name,
	}
}

func (subscriber *ChannelSubscriber) Like(video *Video) {
	fmt.Printf("%s 對影片 \"%s\" 按讚。\n", subscriber.name, video.Name)
}

func (subscriber *ChannelSubscriber) Name() string {
	return subscriber.name
}
