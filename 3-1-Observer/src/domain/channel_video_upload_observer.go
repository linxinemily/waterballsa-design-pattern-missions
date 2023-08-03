package domain

type ChannelVideoUploadedObserver interface {
	react(video *Video)
	Name() string
}
