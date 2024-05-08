package mtd

type Language struct {
	IncorrectAccount             string
	IncorrectPassword            string
	TheAccountDoesNotExist       string
	AccountDisabled              string
	IncorrectToken               string
	The16bitKeyIsNotSet          string
	IncorrectName                string
	TheAccountIsTooShort         string
	ThePasswordIsTooShort        string
	TheAccountAlreadyExists      string
	NoData                       string
	PermissionDenied             string
	DataWithTheSameNameExists    string
	IncorrectBirthday            string
	IncorrectGender              string
	IncorrectPhoneNumber         string
	IncorrectPriority            string
	CompanyDataDoesNotExist      string
	CustomerDataDoesNotExist     string
	IncorrectGroup               string
	GroupDataDoesNotExist        string
	IncorrectExpirationDate      string
	IncorrectCustomer            string
	SalesTargetDataDoesNotExist  string
	IncorrectSalesTarget         string
	TheSalesPlanDataDoesNotExist string
	TypeError                    string
	AfterSalesPersonnelDoNot     string
	TheProductDataDoesNotExist   string
	IncorrectOrderNo             string
	TheSalesManagerDoesNotExist  string
	IncorrectOrderPrice          string
	TheOrderDataDoesNotExist     string
	TheOrderNumberIsDuplicated   string
	IncorrectTitle               string
	IncorrectContent             string
	AnnouncementDataDoesNotExist string
	DistributorDataDoesNotExist  string
	SupplierDataDoesNotExist     string
	QADataDoesNotExist           string
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
		language.IncorrectExpirationDate = "截至日期错误"
		language.IncorrectCustomer = "归属客户错误"
		language.SalesTargetDataDoesNotExist = "销售目标数据不存在"
		language.IncorrectSalesTarget = "销售目标数据错误"
		language.TheSalesPlanDataDoesNotExist = "销售计划数据不存在"
		language.TypeError = "类型错误"
		language.AfterSalesPersonnelDoNot = "售后人员不存在"
		language.TheProductDataDoesNotExist = "产品数据不存在"
		language.IncorrectOrderNo = "订单编号错误"
		language.TheSalesManagerDoesNotExist = "销售经理不存在"
		language.IncorrectOrderPrice = "订单价格错误"
		language.TheOrderDataDoesNotExist = "订单数据不存在"
		language.TheOrderNumberIsDuplicated = "订单编号重复"
		language.IncorrectTitle = "标题错误"
		language.IncorrectContent = "内容错误"
		language.AnnouncementDataDoesNotExist = "公告数据不存在"
		language.DistributorDataDoesNotExist = "渠道商数据不存在"
		language.SupplierDataDoesNotExist = "供应商数据不存在"
		language.QADataDoesNotExist = "Q&A数据不存在"
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
		language.IncorrectExpirationDate = "Incorrect expiration date"
		language.IncorrectCustomer = "Incorrect customer"
		language.SalesTargetDataDoesNotExist = "Sales target data does not exist"
		language.IncorrectSalesTarget = "Incorrect sales target"
		language.TheSalesPlanDataDoesNotExist = "The sales plan data does not exist"
		language.TypeError = "Type error"
		language.AfterSalesPersonnelDoNot = "After-sales personnel do not exist"
		language.TheProductDataDoesNotExist = "The product data does not exist"
		language.IncorrectOrderNo = "Incorrect order no"
		language.TheSalesManagerDoesNotExist = "The sales manager does not exist"
		language.IncorrectOrderPrice = "Incorrect order price"
		language.TheOrderDataDoesNotExist = "The order data does not exist"
		language.TheOrderNumberIsDuplicated = "The order number is duplicated"
		language.IncorrectTitle = "Incorrect title"
		language.IncorrectContent = "Incorrect content"
		language.AnnouncementDataDoesNotExist = "Announcement data does not exist"
		language.DistributorDataDoesNotExist = "Distributor data does not exist"
		language.SupplierDataDoesNotExist = "Supplier data does not exist"
		language.QADataDoesNotExist = "Q&A data does not exist"
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
		language.IncorrectExpirationDate = ""
		language.IncorrectCustomer = ""
		language.SalesTargetDataDoesNotExist = ""
		language.IncorrectSalesTarget = ""
		language.TheSalesPlanDataDoesNotExist = ""
		language.TypeError = ""
		language.AfterSalesPersonnelDoNot = ""
		language.TheProductDataDoesNotExist = ""
		language.IncorrectOrderNo = ""
		language.TheSalesManagerDoesNotExist = ""
		language.IncorrectOrderPrice = ""
		language.TheOrderDataDoesNotExist = ""
		language.TheOrderNumberIsDuplicated = ""
		language.IncorrectTitle = ""
		language.IncorrectContent = ""
		language.AnnouncementDataDoesNotExist = ""
		language.DistributorDataDoesNotExist = ""
		language.SupplierDataDoesNotExist = ""
		language.QADataDoesNotExist = ""
	}
	return language
}
