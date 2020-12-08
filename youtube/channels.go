package youtube

import (
	"errors"
	"fmt"
)

func (c *Client) GetUploadIDFromUsername(channel Channel) (string, error) {
	if len(channel.ID) == 0 && len(channel.Name) == 0{
		return "", errors.New("you have to specify at least channel id or channel name")
	}
	call :=	c.service.Channels.List([]string{"contentDetails"})
	if len(channel.ID) > 0 {
		call.Id(channel.ID)
	} else {
		call.ForUsername(channel.Name)
	}
	resp, err := call.Do()
	if err != nil {
		return "", err
	}

	if len(resp.Items) != 1 {
		return "", errors.New(fmt.Sprintf("could not find a channel with the name %s", channel.Name))
	}
	return resp.Items[0].ContentDetails.RelatedPlaylists.Uploads, nil

}
