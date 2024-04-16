package mtd

type Language struct {
	IncorrectAccount          string
	IncorrectPassword         string
	TheAccountDoesNotExist    string
	AccountDisabled           string
	IncorrectToken            string
	The16bitKeyIsNotSet       string
	IncorrectName             string
	TheAccountIsTooShort      string
	ThePasswordIsTooShort     string
	TheAccountAlreadyExists   string
	NoData                    string
	PermissionDenied          string
	DataWithTheSameNameExists string
	IncorrectBirthday         string
	IncorrectGender           string
	IncorrectPhoneNumber      string
	IncorrectPriority         string
	CompanyDataDoesNotExist   string
	CustomerDataDoesNotExist  string
	IncorrectGroup            string
	GroupDataDoesNotExist     string
}

func SysLang() Language {
	f := FileHelper{}
	checkConf := f.CheckConf()
	language := Language{}
	if checkConf.Lang == "zh" {
		language.IncorrectAccount = "账号错误"
		language.IncorrectPassword = "密码错误"
		language.TheAccountDoesNotExist = "账号不存在"
		language.AccountDisabled = "账户已禁用"
		language.IncorrectToken = "Token信息异常"
		language.The16bitKeyIsNotSet = "未设置16位密钥"
		language.IncorrectName = "名称错误"
		language.TheAccountIsTooShort = "账号长度不够"
		language.ThePasswordIsTooShort = "密码长度不够"
		language.TheAccountAlreadyExists = "账号已存在"
		language.NoData = "数据不存在"
		language.PermissionDenied = "无权限"
		language.DataWithTheSameNameExists = "存在同名数据"
		language.IncorrectBirthday = "出生日期错误"
		language.IncorrectGender = "性别错误"
		language.IncorrectPhoneNumber = "电话号码错误"
		language.IncorrectPriority = "优先级错误"
		language.CompanyDataDoesNotExist = "公司数据不存在"
		language.CustomerDataDoesNotExist = "客户数据不存在"
		language.IncorrectGroup = "小组数据错误"
		language.GroupDataDoesNotExist = "小组数据不存在"
	} else if checkConf.Lang == "en" {
		language.IncorrectAccount = "Incorrect account"
		language.IncorrectPassword = "Incorrect password"
		language.TheAccountDoesNotExist = "The account does not exist"
		language.AccountDisabled = "Account disabled"
		language.IncorrectToken = "Incorrect token"
		language.The16bitKeyIsNotSet = "The 16-bit key is not set"
		language.IncorrectName = "Incorrect name"
		language.TheAccountIsTooShort = "The account is too short"
		language.ThePasswordIsTooShort = "The password is too short"
		language.TheAccountAlreadyExists = "The account already exists"
		language.NoData = "No data"
		language.PermissionDenied = "Permission denied"
		language.DataWithTheSameNameExists = "Data with the same name exists"
		language.IncorrectBirthday = "Incorrect birthday"
		language.IncorrectGender = "Incorrect gender"
		language.IncorrectPhoneNumber = "Incorrect phone number"
		language.IncorrectPriority = "Incorrect priority"
		language.CompanyDataDoesNotExist = "Company data does not exist"
		language.CustomerDataDoesNotExist = "Customer data does not exist"
		language.IncorrectGroup = "Incorrect group"
		language.GroupDataDoesNotExist = "Group data does not exist"
	} else {
		language.IncorrectAccount = ""
		language.IncorrectPassword = ""
		language.TheAccountDoesNotExist = ""
		language.AccountDisabled = ""
		language.IncorrectToken = ""
		language.The16bitKeyIsNotSet = ""
		language.IncorrectName = ""
		language.TheAccountIsTooShort = ""
		language.ThePasswordIsTooShort = ""
		language.TheAccountAlreadyExists = ""
		language.NoData = ""
		language.PermissionDenied = ""
		language.DataWithTheSameNameExists = ""
		language.IncorrectBirthday = ""
		language.IncorrectGender = ""
		language.IncorrectPhoneNumber = ""
		language.IncorrectPriority = ""
		language.CompanyDataDoesNotExist = ""
		language.CustomerDataDoesNotExist = ""
		language.IncorrectGroup = ""
		language.GroupDataDoesNotExist = ""
	}
	return language
}
