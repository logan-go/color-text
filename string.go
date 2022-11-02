package color_text

import (
	"fmt"
	"strings"
)

type String []*Phrase

func NewString(cap int) String {
	return make(String, 0, cap)
}

func (str String) Add(phrase *Phrase) String {
	s := phrase
	str = append(str, s)
	return str
}

func (str String) String() string {
	s := make([]string, len(str))
	for index, value := range str {
		s[index] = value.text
	}
	return strings.Join(s, "")
}

func (str String) ColorString() string {
	list := make([]string, 0, len(str))
	for _, p := range str {
		if p.text == "" {
			continue
		}
		s := make([]string, 0, 3)
		if p.frontColor != "" {
			s = append(s, string(p.frontColor))
		}
		if p.backgroundColor != "" {
			s = append(s, string(p.backgroundColor))
		}
		if p.style != "" {
			s = append(s, string(p.style))
		}

		list = append(list, fmt.Sprintf(stringTemplate, strings.Join(s, ";"), p.text))
	}
	return strings.Join(list, "")
}
