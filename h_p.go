package pingyin

var HP = map[string]string{}

func init() {
	for py, hzs := range PH {
		for _, hz := range []rune(hzs) {
			HP[string(hz)] = py
		}
	}
}
