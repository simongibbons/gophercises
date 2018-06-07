package link

import (
	"strings"
	"testing"
)

func TestParseWithComment(t *testing.T) {
	html := `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`
	links, err := Parse(strings.NewReader(html))
	if err != nil {
		t.Fatalf("There should be no error in parsing")
	}

	if len(links) != 1 {
		t.Fatalf("There should be 1 link")
	}

	link := links[0]
	if link.Href != "/dog-cat" {
		t.Fatalf("Link href is wrong")
	}

	if link.Text != "dog cat " {
		t.Fatalf("Link text is wrong")
	}
}

func TestParseNoHref(t *testing.T) {
	html := "<a>no href</a>"
	links, err := Parse(strings.NewReader(html))

	if err != nil {
		t.Fatalf("There should be no error in parsing")
	}

	if len(links) != 1 {
		t.Fatalf("There should be 1 link")
	}

	link := links[0]
	if link.Href != "" {
		t.Fatalf("Link href is wrong")
	}

	if link.Text != "no href" {
		t.Fatalf("Link text is wrong")
	}
}

func TestParseMultipleStrings(t *testing.T) {
	html := `<a href="/dog">
<span>Something in a span</span>
Text not in a span
<b>Bold text!</b>
</a>
`
	links, err := Parse(strings.NewReader(html))

	if err != nil {
		t.Fatalf("There should be no error in parsing")
	}

	if len(links) != 1 {
		t.Fatalf("There should be 1 link")
	}

	link := links[0]
	if link.Href != "/dog" {
		t.Fatalf("Link href is wrong")
	}

	if link.Text != "Something in a spanText not in a spanBold text!" {
		t.Fatalf("Link text is wrong")
	}
}
