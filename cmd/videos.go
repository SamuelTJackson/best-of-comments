/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/SamuelTJackson/best-of-comments/storage/file"
	"github.com/SamuelTJackson/best-of-comments/youtube"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		channel := youtube.Channel{
			ID:   channelID,
			Name: channelName,
		}
		fmt.Printf("Output file: %s, Channel: %s\n", Output, channel)
		if _, err := os.Create(Output); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		storage := file.NewStorage(Output, Input, file.Mode(Mode))
		c, err := youtube.GetClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		id, err := c.GetUploadIDFromUsername(channel)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var ch = make(chan *youtube.Video)
		go c.GetUploadsForID(id, "",ch)
		var videos []*youtube.Video
		for video := range ch {
			videos = append(videos, video)
		}
		storage.Save(youtube.Comment{
			VideoID:            "",
			Text:               "",
			Author:             "",
			AuthorProfileImage: "",
			ID:                 "",
		})
		if err := storage.SaveVideos(videos...); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}
var channelName string
var channelID string
func init() {
	getCmd.AddCommand(videosCmd)
	videosCmd.Flags().StringVarP(&channelName, "channel-name","c",viper.GetString("CHANNEL_NAME"),"Set the channel name")
	videosCmd.Flags().StringVarP(&channelID, "channel-id", "d", viper.GetString("CHANNEL_ID"), "set the channel id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
