package storage

import "github.com/SamuelTJackson/best-of-comments/youtube"

type storage interface {
	SaveVideos(v ...*youtube.Item) error
	GetVideo(id string) (*youtube.Video, error)
	SaveComments(c ...*youtube.Item) error
}
