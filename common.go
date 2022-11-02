package color_text

type FrontColor string

const (
	FrontColorDefault     FrontColor = ""
	FrontColorBlack       FrontColor = "30"
	FrontColorRed         FrontColor = "31"
	FrontColorGreen       FrontColor = "32"
	FrontColorBrown       FrontColor = "33"
	FrontColorBlue        FrontColor = "34"
	FrontColorPurple      FrontColor = "35"
	FrontColorCyan        FrontColor = "36"
	FrontColorWhite       FrontColor = "37"
	FrontColorLightBlack  FrontColor = "90"
	FrontColorLightRed    FrontColor = "91"
	FrontColorLightGreen  FrontColor = "92"
	FrontColorLightBrown  FrontColor = "93"
	FrontColorLightBlue   FrontColor = "94"
	FrontColorLightPurple FrontColor = "95"
	FrontColorLightCyan   FrontColor = "96"
	FrontColorLightWhite  FrontColor = "97"
)

type BackgroundColor string

const (
	BackgroundColorDefault     BackgroundColor = ""
	BackgroundColorBlack       BackgroundColor = "40"
	BackgroundColorRed         BackgroundColor = "41"
	BackgroundColorGreen       BackgroundColor = "42"
	BackgroundColorBrown       BackgroundColor = "43"
	BackgroundColorBlue        BackgroundColor = "44"
	BackgroundColorPurple      BackgroundColor = "45"
	BackgroundColorCyan        BackgroundColor = "46"
	BackgroundColorWhite       BackgroundColor = "47"
	BackgroundColorLightBlack  BackgroundColor = "100"
	BackgroundColorLightRed    BackgroundColor = "101"
	BackgroundColorLightGreen  BackgroundColor = "102"
	BackgroundColorLightBrown  BackgroundColor = "103"
	BackgroundColorLightBlue   BackgroundColor = "104"
	BackgroundColorLightPurple BackgroundColor = "105"
	BackgroundColorLightCyan   BackgroundColor = "106"
	BackgroundColorLightWhite  BackgroundColor = "107"
)

type Style string

const (
	StyleDefault   Style = ""
	StyleBold      Style = "1"
	StyleFaint     Style = "2"
	StyleUnderline Style = "4"
	StyleFlash     Style = "5"
	StyleReversed  Style = "7"
	StyleHidden    Style = "8"
)

const stringTemplate = "\033[%sm%s\033[m"
