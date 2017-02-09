# relinkify

[![Go Report Card](https://goreportcard.com/badge/github.com/oneumyvakin/relinkify)](https://goreportcard.com/report/github.com/oneumyvakin/relinkify)

Simple linkify plan-text with HTML `<a>` tags

```go
package main
 
import (
    "fmt"
    
    "github.com/oneumyvakin/relinkify"
)

func main() {
    text := "This is a link http://google.com"
    
    fmt.Println(relinkify.Linkify(text))
}
```

Prints

```
This is a link <a href="http://google.com">http://google.com</a>
```