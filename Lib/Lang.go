package lib

type lang struct {
}

func Lang() lang {
	l := lang{}
	checkConf := CheckConf()
	if checkConf.Lang == "zh" {
	} else if checkConf.Lang == "en" {
	} else {
	}
	return l
}
