package link

import (
	"testing"
	"strings"
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
