package domain

import "time"

type WaterBall struct {
	*ChannelSubscriber
}

func NewWaterBall() *WaterBall {
	return &WaterBall{
		NewChannelSubscriber("水球"),
	}
}

func (fireball *WaterBall) react(video *Video) {
	if video.Duration >= 3*time.Minute {
		fireball.Like(video)
	}
}
