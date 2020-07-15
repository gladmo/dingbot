package dingbot

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gladmo/dingbot"
)

var feedcard = &cobra.Command{
	Use:   "feedcard",
	Short: "send dingtalk feedcard message",
	Run: func(cmd *cobra.Command, args []string) {
		link, err := cmd.PersistentFlags().GetStringArray("json-link")
		if err != nil {
			fmt.Println(err.Error())
		}

		var links []dingbot.Links
		for _, v := range link {
			var oneLink dingbot.Links
			err := json.Unmarshal([]byte(v), &oneLink)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			links = append(links, oneLink)
		}

		msg := dingbot.FeedCardMessage(links...)

		err = dingTalk.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	feedcard.PersistentFlags().StringArray("json-link", []string{}, "feedcard link string, use json")
}
