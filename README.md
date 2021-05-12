# lm-sdk-go
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/vkumbhar94/lm-sdk-go)

Getting Started
---------------

```go
package main

import (
    "fmt"
    "github.com/vkumbhar94/lm-sdk-go/client"
    "github.com/vkumbhar94/lm-sdk-go/client/lm"
)

func NewLMClient() *client.LMSdkGo {
    domain := "YOUR_COMPANY.logicmonitor.com"
    accessID := "YOUR_ACCESS_ID"
    accessKey := "YOUR_ACCESS_KEY"

    config := client.NewConfig()
    config.SetAccountDomain(&domain)
    config.SetAccessID(&accessID)
    config.SetAccessKey(&accessKey)

    return client.New(config)
}

func main() {
    client := NewLMClient()
}
```

### License
[![license](https://img.shields.io/github/license/logicmonitor/lm-sdk-go.svg?style=flat-square)](https://github.com/vkumbhar94/lm-sdk-go/blob/master/LICENSE)
