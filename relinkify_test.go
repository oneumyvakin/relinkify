package relinkify

import (
	"testing"
	"fmt"
)

func TestLinkify(t *testing.T) {
	data := map[string]string{
		Linkify("http://ya.ru"): `<a href="http://ya.ru">http://ya.ru</a>`,
		Linkify("http://ya.ru https://g.com http://ya.ru"): `<a href="http://ya.ru">http://ya.ru</a> <a href="https://g.com">https://g.com</a> <a href="http://ya.ru">http://ya.ru</a>`,
		Linkify("http://ya.ru\r\nhttps://g.com\nhttp://ya.ru"): "<a href=\"http://ya.ru\">http://ya.ru</a>\r\n<a href=\"https://g.com\">https://g.com</a>\n<a href=\"http://ya.ru\">http://ya.ru</a>",
	}

	for in, out := range data {
		if in != out {
			t.Error(fmt.Sprintf("Error: %s != %s", in, out))
		}
	}
}

func TestLinkificator(t *testing.T) {
	li := Linkificator{
		Pattern: Pattern,
		Re: re,
		Attrs: `target="__blank"`,
	}

	data := map[string]string{
		li.Linkify("http://ya.ru"): `<a href="http://ya.ru" target="__blank">http://ya.ru</a>`,
		li.Linkify("http://ya.ru https://g.com http://ya.ru"): `<a href="http://ya.ru" target="__blank">http://ya.ru</a> <a href="https://g.com" target="__blank">https://g.com</a> <a href="http://ya.ru" target="__blank">http://ya.ru</a>`,
		li.Linkify("http://ya.ru\r\nhttps://g.com\nhttp://ya.ru"): "<a href=\"http://ya.ru\" target=\"__blank\">http://ya.ru</a>\r\n<a href=\"https://g.com\" target=\"__blank\">https://g.com</a>\n<a href=\"http://ya.ru\" target=\"__blank\">http://ya.ru</a>",
	}

	for in, out := range data {
		if in != out {
			t.Error(fmt.Sprintf("Error: %s != %s", in, out))
		}
	}
}