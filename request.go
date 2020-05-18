package dingbot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const dingTalkHost = `https://oapi.dingtalk.com`

type DingTalk struct {
	AccessToken string
	Secret      string
}

// getURL 构造请求地址
func (th *DingTalk) getURL() (URL string) {
	URL = fmt.Sprintf("%s/robot/send?access_token=%s", dingTalkHost, th.AccessToken)

	if th.Secret != "" {
		ts := time.Now().UnixNano() / 1e6
		stringToSign := fmt.Sprintf("%d\n%s", ts, th.Secret)
		mac := hmac.New(sha256.New, []byte(th.Secret))
		mac.Write([]byte(stringToSign))
		signData := mac.Sum(nil)
		base64sign := base64.StdEncoding.EncodeToString(signData)

		URL += fmt.Sprintf("&timestamp=%d&sign=%s", ts, url.QueryEscape(base64sign))
	}

	return
}

// Send 发送钉钉消息
func (th DingTalk) Send(msg Message) (err error) {
	req, err := http.NewRequest("POST", th.getURL(), strings.NewReader(msg.String()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("status not ok, code: %d", res.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var r dingRes
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}

	if r.ErrMsg != "ok" {
		err = errors.New(fmt.Sprintf("Ding ding send error. res: %s", string(body)))
		return
	}

	return
}

type dingRes struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
