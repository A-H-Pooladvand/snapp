package str

import (
	"fmt"
	"strings"
)

func Wrap(word, char string) string {
	return fmt.Sprintf("%s%s%s", char, word, char)
}

func Between(target, start, end string) string {
	target = strings.TrimLeft(target, end)
	return strings.TrimRight(target, start)
}

func ToString(value any) string {
	return fmt.Sprintf("%v", value)
}
