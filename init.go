package pingyin

import (
	"sort"
	"strings"
)

func init() {
	// 拆开PH
	for py, hzl := range PH { // 每行
		mask := map[string]bool{}   // 用于去重
		for _, hzstr := range hzl { // 每个value
			hzs := strings.Split(hzstr, "") // 拆开成单个
			for _, hz := range hzs {        // 每个
				if mask[hz] { // 去重
					continue
				}
				mask[hz] = true
			}
		}
		var hzs []string // 去重排序
		for key := range mask {
			hzs = append(hzs, key)
		}
		sort.Strings(hzs)
		PH[py] = hzs
	}

	for py, hzs := range PH {
		for _, hz := range hzs {
			HP[string(hz)] = append(HP[string(hz)], py)
		}
	}
	// 排序
	for hz, pys := range HP {
		sort.Strings(pys)
		HP[hz] = pys
	}
}
