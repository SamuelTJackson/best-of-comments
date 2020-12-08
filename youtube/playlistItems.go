package youtube

import (
	"fmt"
	"time"
)
const (
	MaxResults = 1000
)
func (c *Client) GetUploadsForID(id string, nextPageToken string, ch chan *Video) {
	call := c.service.PlaylistItems.List([]string{"contentDetails"})
	call.MaxResults(MaxResults)
	call.PlaylistId(id)
	if len(nextPageToken) != 0 {
		call.PageToken(nextPageToken)
	}
	resp, err := call.Do()
	if err != nil {
		close(ch)
		return
	}
	for _, item := range resp.Items {
		t, err := time.Parse(time.RFC3339Nano, item.ContentDetails.VideoPublishedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ch <- &Video{
			ID:        item.ContentDetails.VideoId,
			Published: t,
		}
	}
	fmt.Printf("Got %d new video informations\n", len(resp.Items))
	if resp.NextPageToken != "" {
		c.GetUploadsForID(id, resp.NextPageToken, ch)
	} else {
		close(ch)
	}

}
