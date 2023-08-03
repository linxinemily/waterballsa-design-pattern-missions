package domain

import (
	"fmt"
)

type Channel struct {
	name      string
	observers []ChannelVideoUploadedObserver
}

func NewChannel(name string) *Channel {
	return &Channel{name: name}
}

func (channel *Channel) Upload(video *Video) {
	video.Channel = channel
	fmt.Printf("頻道 %s 上架了一則新影片 \"%s\"。\n", channel.name, video.Name)
	channel.notify(video)
}

func (channel *Channel) notify(video *Video) {
	for _, observer := range channel.observers {
		observer.react(video)
	}
}

func (channel *Channel) Subscribe(observer ChannelVideoUploadedObserver) {
	fmt.Printf("%s 訂閱了 %s。\n", observer.Name(), channel.name)

	channel.observers = append(channel.observers, observer)
}

func (channel *Channel) Unsubscribe(observer ChannelVideoUploadedObserver) {
	fmt.Printf("%s 解除訂閱了 %s。\n", observer.Name(), channel.name)
	for i := 0; i < len(channel.observers); i++ {
		if channel.observers[i] == observer {
			channel.observers = append(channel.observers[:i], channel.observers[i+1:]...)
		}
	}
}
