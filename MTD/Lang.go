package mtd

type Language struct {
	IncorrectAccount  string
	IncorrectPassword string
}

func SysLang() Language {
	f := FileHelper{}
	checkConf := f.CheckConf()
	language := Language{}
	if checkConf.Lang == "zh" {
		language.IncorrectAccount = "账号错误"
		language.IncorrectPassword = "密码错误"
	} else if checkConf.Lang == "en" {
		language.IncorrectAccount = "Incorrect account"
		language.IncorrectPassword = "Incorrect password"
	} else {
		language.IncorrectAccount = ""
		language.IncorrectPassword = ""
	}
	return language
}
