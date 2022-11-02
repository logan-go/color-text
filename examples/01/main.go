package main

import (
	"fmt"
	text "github.com/logan-go/color-text"
)

func main() {
	//strs := text.NewString(3)
	//
	//str1 := text.NewPhrase("我是前面的闪耀小粉").WithFrontColor(text.FrontColorLightPurple).WithStyle(text.StyleFlash)
	//
	//str2 := text.NewPhrase("我是中间的蓝底小红").WithFrontColor(text.FrontColorRed).WithBackgroundColor(text.BackgroundColorBlue)
	//
	//str3 := text.NewPhrase("我是后面的正常")
	//
	//strs = strs.Add(str1).Add(str2).Add(str3)

	strs := text.NewString(3).
		Add(text.NewPhrase("我是前面的闪耀小粉").WithFrontColor(text.FrontColorLightRed).WithBackgroundColor(text.BackgroundColorLightRed).WithStyle(text.StyleFaint)).
		Add(text.NewPhrase("我是中间的蓝底小红").WithFrontColor(text.FrontColorRed).WithBackgroundColor(text.BackgroundColorBlue)).
		Add(text.NewPhrase("我是后面的正常"))

	fmt.Println(strs.ColorString())
}
