package youtube

import (
	"fmt"
	"time"
)

type Item interface {
	GetID() string
}

type Video struct {
	ID string `json:"id"`
	Published time.Time `json:"published"`
}

func (v Video) GetID() string {
	return v.ID
}

func (v Video) String() string {
	return fmt.Sprintf("id: %s, published at: %v",v.ID, v.Published)
}
type Comment struct {
	VideoID string `json:"video_id"`
	Text string `json:"text"`
	Author string `json:"author"`
	AuthorProfileImage string `json:"author_profile_image"`
	ID string `json:"id"`
}
func (c Comment) GetID() string {
	return c.ID
}
type Channel struct {
	ID string
	Name string
}

func (c Channel) GetID() string {
	return c.ID
}

func (c Channel) String() string {
	if len(c.ID) > 0 && len(c.Name) > 0 {
		return fmt.Sprintf("[ID: %s Name: %s]", c.ID, c.Name)
	}
	if len(c.ID) > 0 {
		return fmt.Sprintf("[ID: %s]", c.ID)
	}
	return fmt.Sprintf("[Name: %s]", c.Name)

}
