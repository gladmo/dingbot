# Dingbot
DingTalk robot golang library, and command line

# Usage
## As library
```golang
package main

import (
	"github.com/gladmo/dingbot"
)

func main() {
	token := "you ding talk access token"
	secret := "you ding talk secret, skip if old robot"

	text := "hello world"
	msg := dingbot.TextMessage(text)
	ding := &dingbot.DingTalk{
		AccessToken: token,
		Secret:      secret,
	}

	err := ding.Send(msg)
	if err != nil {
		panic(err)
	}
}
```

## Command line
### preview
```shell script
send dingtalk message

Usage:
  dingbot [command]

Available Commands:
  feedcard    send dingtalk feedcard message
  help        Help about any command
  init        init config
  link        send dingtalk link message
  markdown    send dingtalk markdown message
  text        send dingtalk text message
  version     Print the version number of Dingbot

Flags:
      --config string   config file (default is $HOME/dingbot.toml)
  -h, --help            help for dingbot
      --secret string   dingtalk robot secret
      --token string    dingtalk robot token (require)

Use "dingbot [command] --help" for more information about a command.
```

### use Docker
```shell script
$ docker pull gladmo/dingbot
$ docker run --rm gladmo/dingbot dingbot --help
```

### Docker command alias
```shell script
$ alias dingbot="docker run --rm gladmo/dingbot dingbot --token 'you token' --secret 'you secret'"
# then
$ dingbot text "hello world"
```

### build
```shell script
$ git clone https://github.com/gladmo/dingbot.git
$ cd dingbot
$ go mod vendor
$ go build -o dingbot cmd/main.go
$ ./dingbot version
```

### send text message
```shell script
$ ./dingbot --token "you token" --secret "you secret" text "hello world"
# or
$ echo "hello world" | ./dingbot --token "you token" --secret "you secret" text
```

### send markdown message
```shell script
$ ./dingbot markdown --title hello '## hello world'
# or
$ echo "hello world" | ./dingbot --token "you token" --secret "you secret" markdown --title hello
```

### send link message
```shell script
$ ./dingbot --token "you token" --secret "you secret" link --title hello --message-url 'https://6cm.co' 'hello world'
# or
$ echo "hello world" | ./dingbot --token "you token" --secret "you secret" link --title hello --message-url 'https://6cm.co'
```

### send feed card message
```shell script
$ ./dingbot --token "you token" --secret "you secret" feedcard --json-link '{"title":"时代的火车向前开","messageURL":"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI","picURL":"https://gw.alicdn.com/tfs/TB1ayl9mpYqK1RjSZLeXXbXppXa-170-62.png"}' --json-link '{"title":"时代的火车向前开","messageURL":"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI","picURL":"https://gw.alicdn.com/tfs/TB1ayl9mpYqK1RjSZLeXXbXppXa-170-62.png"}'
```

### command alias
```shell script
# replace path/to to you dingbot path
$ alias dingbot="path/to/dingbot --token 'you token' --secret 'you secret'"
# then
$ dingbot text "hello world"
```

### use config 
```shell script
$ ./dingbot --token "you token" --secret "you secret" init dintbot.toml
$ ./dingbot --config dintbot.toml text "hello world"
```