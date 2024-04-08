package mtd

type Language struct {
	IncorrectAccount       string
	IncorrectPassword      string
	TheAccountDoesNotExist string
	IncorrectToken         string
	The16bitKeyIsNotSet    string
	IncorrectName          string
}

func SysLang() Language {
	f := FileHelper{}
	checkConf := f.CheckConf()
	language := Language{}
	if checkConf.Lang == "zh" {
		language.IncorrectAccount = "账号错误"
		language.IncorrectPassword = "密码错误"
		language.TheAccountDoesNotExist = "账号不存在"
		language.IncorrectToken = "Token信息异常"
		language.The16bitKeyIsNotSet = "未设置16位密钥"
		language.IncorrectName = "名称不正确"
	} else if checkConf.Lang == "en" {
		language.IncorrectAccount = "Incorrect account"
		language.IncorrectPassword = "Incorrect password"
		language.TheAccountDoesNotExist = "The account does not exist"
		language.IncorrectToken = "Incorrect token"
		language.The16bitKeyIsNotSet = "The 16-bit key is not set"
		language.IncorrectName = "Incorrect name"
	} else {
		language.IncorrectAccount = ""
		language.IncorrectPassword = ""
		language.TheAccountDoesNotExist = ""
		language.IncorrectToken = ""
		language.The16bitKeyIsNotSet = ""
		language.IncorrectName = ""
	}
	return language
}
