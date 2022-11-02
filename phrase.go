package color_text

type Phrase struct {
	text            string
	frontColor      FrontColor
	backgroundColor BackgroundColor
	style           Style
}

func NewPhrase(text string) *Phrase {
	return &Phrase{
		text:            text,
		frontColor:      FrontColorDefault,
		backgroundColor: BackgroundColorDefault,
		style:           StyleDefault,
	}
}

func (p *Phrase) WithFrontColor(color FrontColor) *Phrase {
	p.frontColor = color
	return p
}

func (p *Phrase) WithBackgroundColor(color BackgroundColor) *Phrase {
	p.backgroundColor = color
	return p
}

func (p *Phrase) WithStyle(style Style) *Phrase {
	p.style = style
	return p
}
