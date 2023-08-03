package domain

import "time"

type Video struct {
	Name        string
	Description string
	Duration    time.Duration
	Channel     *Channel
}

func NewVideo(name string, description string, duration time.Duration) *Video {
	return &Video{
		Name:        name,
		Description: description,
		Duration:    duration,
	}
}
