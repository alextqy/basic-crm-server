package mtd

type lang struct {
}

func Lang() lang {
	f := FileHelper{}
	l := lang{}
	checkConf := f.CheckConf()
	if checkConf.Lang == "zh" {
	} else if checkConf.Lang == "en" {
	} else {
	}
	return l
}
