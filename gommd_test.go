package gommd

import (
	"strings"
	"testing"
)

var testPage = `
This is a test
==

Hello, I am testing *MultiMarkdown*.
`

var expect = `<h1 id="thisisatest">This is a test</h1>

<p>Hello, I am testing <em>MultiMarkdown</em>.</p>
`

func TestMarkdown(t *testing.T) {
	res := MarkdownToString(testPage, 0, FORMAT_HTML)
	if strings.TrimSpace(res) != strings.TrimSpace(expect) {
		t.Fatalf("simple translation failed, expected:\n%s\ngot:\n%s\n",
			expect, res)
	}
}
