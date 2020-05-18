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