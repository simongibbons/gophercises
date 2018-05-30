package cyoa

import "testing"

func TestParseCYOA(t *testing.T) {
	CYOAJson := `{
  "intro": {
    "title": "The Little Blue Gopher",
    "story": [
		"Some text",
		"Some more text"
    ],
    "options": [
      {
        "text": "Interesting Option",
        "arc": "home"
      }
    ]
  },
  "home": {
    "title": "Home Sweet Home",
    "story": [
      "Lovely at Home"
    ],
    "options": []
  }
}`
	cyoa, err := ParseCYOA([]byte(CYOAJson))

	if err != nil {
		t.Errorf("Error shouldn't have been thrown while parsing")
	}

	if len(cyoa) != 2 {
		t.Errorf("Should have 2 arcs")
	}
}

func TestValidateCYOANoIntro(t *testing.T) {
	noIntroJSON := `{
  "home": {
    "title": "Home Sweet Home",
    "story": [
      "Lovely at Home"
    ],
    "options": []
  }
}`
	_, err := ParseCYOA([]byte(noIntroJSON))

	if err != noInitialArcError {
		t.Errorf("Should have detected `intro` arc isn't present")
	}
}

func TestValidateCYOAInvalidOption(t *testing.T) {
	invalidOptionJSON := `{
  "intro": {
    "title": "Intro",
    "story": [],
    "options": [{"arc": "not_an_arc"}]
  }
}`
	_, err := ParseCYOA([]byte(invalidOptionJSON))

	if err != invalidOptionError {
		t.Errorf("Should have detected invalid options is present")
	}
}
