package relinkify

import (
	"fmt"
	"regexp"
)

// Pattern for URLs
var Pattern = `https?:\/\/\S+`
var re = regexp.MustCompile(Pattern)

// Linkificator holds stuff to customize linkification
type Linkificator struct {
	Pattern string
	Re      *regexp.Regexp
	Attrs   string
}

// Linkify plain-text
func (li *Linkificator) Linkify(text string) string {
	if li.Pattern == "" {
		li.Pattern = Pattern
	}
	if li.Re == nil {
		li.Re = re
	}
	return re.ReplaceAllStringFunc(text, li.replacer)
}

func (li *Linkificator) replacer(in string) string {
	attrs := li.Attrs
	if attrs != "" {
		attrs = " " + li.Attrs
	}
	return fmt.Sprintf(`<a href="%[1]s"%[2]s>%[1]s</a>`, in, attrs)
}

// Linkify is shortcut for Linkificator.Linkify
func Linkify(text string) string {
	li := Linkificator{}
	return li.Linkify(text)
}
