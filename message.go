package dingbot

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Message struct {
	MsgType    string      `json:"msgtype"`
	Text       *text       `json:"text,omitempty"`
	Link       *link       `json:"link,omitempty"`
	Markdown   *markdown   `json:"markdown,omitempty"`
	ActionCard *actionCard `json:"actionCard,omitempty"`
	FeedCard   *feedCard   `json:"feedCard,omitempty"`
	AtMobiles  *at         `json:"at,omitempty"`
}

func (th Message) String() string {
	b, _ := json.Marshal(th)
	return string(b)
}

func (th *Message) At(atAll bool, mobiles ...string) (err error) {
	switch th.MsgType {
	case "":
		err = errors.New("msgtype is empty")
		return
	}

	// todo check mobile

	th.AtMobiles = &at{
		IsAtAll:   atAll,
		AtMobiles: mobiles,
	}
	return
}

type buttons struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type actionCard struct {
	Title          string    `json:"title"`          // 必填
	Text           string    `json:"text"`           // 必填
	BtnOrientation string    `json:"btnOrientation"` // 0-按钮竖直排列，1-按钮横向排列
	SingleTitle    string    `json:"singleTitle"`
	SingleURL      string    `json:"singleURL"`
	Buttons        []buttons `json:"btns"`
}

type links struct {
	Title      string `json:"title"`      // 必填
	MessageURL string `json:"messageURL"` // 必填
	PicURL     string `json:"picURL"`     // 必填
}

type feedCard struct {
	Links []links `json:"links"`
}

type markdown struct {
	Title string `json:"title"` // 必填
	Text  string `json:"text"`  // 必填
}

type link struct {
	Text       string `json:"text"`       // 必填
	Title      string `json:"title"`      // 必填
	MessageURL string `json:"messageUrl"` // 必填
	PicURL     string `json:"picUrl"`
}

type text struct {
	Content string `json:"content"` // 必填
}

type at struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

func TextMessage(content string) (msg Message) {
	msg.MsgType = "text"
	msg.Text = &text{Content: content}
	return
}

func MarkdownMessage(title, content string) (msg Message) {
	msg.MsgType = "markdown"
	msg.Markdown = &markdown{Text: content, Title: title}
	return
}
