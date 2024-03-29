package mtd

type lang struct {
}

func Lang() lang {
	s := SysHelper{}
	l := lang{}
	checkConf := s.CheckConf()
	if checkConf.Lang == "zh" {
	} else if checkConf.Lang == "en" {
	} else {
	}
	return l
}
