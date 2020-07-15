package dingbot

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Message dingTalk all message type
type Message struct {
	MsgType    string      `json:"msgtype"`
	Text       *text       `json:"text,omitempty"`
	Link       *link       `json:"link,omitempty"`
	Markdown   *markdown   `json:"markdown,omitempty"`
	ActionCard *actionCard `json:"actionCard,omitempty"`
	FeedCard   *feedCard   `json:"feedCard,omitempty"`
	AtMobiles  *at         `json:"at,omitempty"`
}

// String message to string
func (th Message) String() string {
	b, _ := json.Marshal(th)
	return string(b)
}

// At someone
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

// buttons actionCard buttons
type buttons struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// actionCard message
// todo
type actionCard struct {
	Title          string    `json:"title"`          // 必填
	Text           string    `json:"text"`           // 必填
	BtnOrientation string    `json:"btnOrientation"` // 0-按钮竖直排列，1-按钮横向排列
	SingleTitle    string    `json:"singleTitle"`
	SingleURL      string    `json:"singleURL"`
	Buttons        []buttons `json:"btns"`
}

// Links ...
type Links struct {
	Title      string `json:"title"`      // 必填
	MessageURL string `json:"messageURL"` // 必填
	PicURL     string `json:"picURL"`     // 必填
}

// feedCard message
// todo
type feedCard struct {
	Links []Links `json:"links"`
}

// markdown message
type markdown struct {
	Title string `json:"title"` // 必填
	Text  string `json:"text"`  // 必填
}

// link message
type link struct {
	Text       string `json:"text"`       // 必填
	Title      string `json:"title"`      // 必填
	MessageURL string `json:"messageUrl"` // 必填
	PicURL     string `json:"picUrl"`
}

// text message
type text struct {
	Content string `json:"content"` // 必填
}

// at user
type at struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// TextMessage dingTalk text message
func TextMessage(content string) (msg Message) {
	msg.MsgType = "text"
	msg.Text = &text{Content: content}
	return
}

// MarkdownMessage dingTalk markdown message
func MarkdownMessage(title, content string) (msg Message) {
	msg.MsgType = "markdown"
	msg.Markdown = &markdown{Text: content, Title: title}
	return
}

// LinkMessage dingTalk link message
func LinkMessage(title, text, messageURL, picURL string) (msg Message) {
	msg.MsgType = "link"
	msg.Link = &link{
		Text:       title,
		Title:      text,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	return
}

// FeedCardLink create feedCard link
func FeedCardLink(title, messageURL, picURL string) (l Links) {
	l.Title = title
	l.MessageURL = messageURL
	l.PicURL = picURL
	return
}

// FeedCardMessage dingTalk feedCard message
func FeedCardMessage(links ...Links) (msg Message) {
	msg.MsgType = "feedCard"
	msg.FeedCard = &feedCard{Links: links}
	return
}
