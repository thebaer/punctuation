// punctuation takes text and gives back only punctuation.
package punctuation

import (
	"bufio"
	"os"
	"strings"
)

const validPunc = ";:'\",!?.-()"

func Extract(f *os.File) string {
	var out string
	in := bufio.NewScanner(f)
	for in.Scan() {
		out += strings.Map(func(r rune) rune {
			if strings.IndexRune(validPunc, r) >= 0 {
				return r
			}
			return -1
		}, in.Text())
	}

	return out
}
