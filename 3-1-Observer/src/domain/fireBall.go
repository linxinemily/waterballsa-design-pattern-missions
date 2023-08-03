package domain

import "time"

type FireBall struct {
	*ChannelSubscriber
}

func NewFireBall() *FireBall {
	return &FireBall{
		NewChannelSubscriber("火球"),
	}
}

func (fireball *FireBall) react(video *Video) {
	if video.Duration <= 1*time.Minute {
		video.Channel.Unsubscribe(fireball)
	}
}
