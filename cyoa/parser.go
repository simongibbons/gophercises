package cyoa

import (
	"encoding/json"
	"errors"
)

type StoryArc struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type CYOA map[string]StoryArc

const (
	initialArc = "intro"
)

var (
	noInitialArcError  = errors.New("no `intro` arc present")
	invalidOptionError = errors.New("invalid Option")
)

func (c CYOA) GetArc(arc_id string) *StoryArc {
	arc, ok := c[arc_id]
	if !ok {
		return nil
	}
	return &arc
}

// ParseJSON parses the JSON representation of a CYOA data file.
func ParseJSON(cyoaJson []byte) (cyoa CYOA, err error) {
	err = json.Unmarshal(cyoaJson, &cyoa)
	if err != nil {
		return nil, err
	}

	err = validateCYOA(cyoa)
	if err != nil {
		return nil, err
	}

	return cyoa, nil
}

func validateCYOA(cyoa CYOA) error {
	if arc := cyoa.GetArc(initialArc); arc == nil {
		return noInitialArcError
	}

	// ensure that all story arcs referenced in options are present.
	for _, arc := range cyoa {
		for _, option := range arc.Options {
			option_arc := cyoa.GetArc(option.Arc)
			if option_arc == nil {
				return invalidOptionError
			}
		}
	}

	return nil
}
