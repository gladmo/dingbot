package dingbot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/gladmo/dingbot"
)

var link = &cobra.Command{
	Use:   "link",
	Short: "send dingtalk link message",
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.PersistentFlags().GetString("title")
		if err != nil {
			fmt.Println(err.Error())
		}
		messageURL, err := cmd.PersistentFlags().GetString("message-url")
		if err != nil {
			fmt.Println(err.Error())
		}
		picURL, err := cmd.PersistentFlags().GetString("pic-url")
		if err != nil {
			fmt.Println(err.Error())
		}

		var text string
		if len(args) == 0 {
			b, _ := ioutil.ReadAll(os.Stdin)
			text = string(b)
		} else {
			text = args[0]
		}

		msg := dingbot.LinkMessage(title, text, messageURL, picURL)

		err = dingTalk.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// 从stdin中获取数据
			f, err := os.Stdin.Stat()
			if err != nil {
				return err
			}
			if f.Size() == 0 {
				return errors.New("you need some message to send")
			}
		}
		return nil
	},
}

func init() {
	link.PersistentFlags().String("title", "", "link title")
	link.PersistentFlags().String("message-url", "", "link message url")
	link.PersistentFlags().String("pic-url", "", "link pic url")
}
