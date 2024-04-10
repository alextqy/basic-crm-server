package mtd

type Language struct {
	IncorrectAccount        string
	IncorrectPassword       string
	TheAccountDoesNotExist  string
	IncorrectToken          string
	The16bitKeyIsNotSet     string
	IncorrectName           string
	TheAccountIsTooShort    string
	ThePasswordIsTooShort   string
	TheAccountAlreadyExists string
	NoData                  string
	PermissionDenied        string
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
		language.TheAccountIsTooShort = "账号长度不够"
		language.ThePasswordIsTooShort = "密码长度不够"
		language.TheAccountAlreadyExists = "账号已存在"
		language.NoData = "数据不存在"
		language.PermissionDenied = "无权限"
	} else if checkConf.Lang == "en" {
		language.IncorrectAccount = "Incorrect account"
		language.IncorrectPassword = "Incorrect password"
		language.TheAccountDoesNotExist = "The account does not exist"
		language.IncorrectToken = "Incorrect token"
		language.The16bitKeyIsNotSet = "The 16-bit key is not set"
		language.IncorrectName = "Incorrect name"
		language.TheAccountIsTooShort = "The account is too short"
		language.ThePasswordIsTooShort = "The password is too short"
		language.TheAccountAlreadyExists = "The account already exists"
		language.NoData = "No data"
		language.PermissionDenied = "Permission denied"
	} else {
		language.IncorrectAccount = ""
		language.IncorrectPassword = ""
		language.TheAccountDoesNotExist = ""
		language.IncorrectToken = ""
		language.The16bitKeyIsNotSet = ""
		language.IncorrectName = ""
		language.TheAccountIsTooShort = ""
		language.ThePasswordIsTooShort = ""
		language.TheAccountAlreadyExists = ""
		language.NoData = ""
		language.PermissionDenied = ""
	}
	return language
}
