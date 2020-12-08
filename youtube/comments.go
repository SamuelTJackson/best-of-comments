package youtube

import "fmt"


func (c *Client) GetCommentsByVideoID(id string, nextPageToken string, ch chan *Comment) {
	call := c.service.CommentThreads.List([]string{"snippet"})
	call.VideoId(id)
	if len(nextPageToken) != 0 {
		call.PageToken(nextPageToken)
	}
	call.MaxResults(MaxResults)
	resp, err := call.Do()
	if err != nil {
		fmt.Println(err)
		close(ch)
	}

	var comments []*Comment
	for _, comment := range resp.Items {
		comments = append(comments, &Comment{
			VideoID:            id,
			Text:               comment.Snippet.TopLevelComment.Snippet.TextDisplay,
			Author:             comment.Snippet.TopLevelComment.Snippet.AuthorDisplayName,
			AuthorProfileImage: comment.Snippet.TopLevelComment.Snippet.AuthorProfileImageUrl,
			ID:                 comment.Id,
		})

	}
	fmt.Printf("Got %call new Comments for video %s\n", len(resp.Items), id)
	if resp.NextPageToken != "" {
		c.GetCommentsByVideoID(id, resp.NextPageToken, ch)
	} else {
		close(ch)
	}
}

