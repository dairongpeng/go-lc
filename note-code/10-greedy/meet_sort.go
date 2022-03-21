package main

import "sort"

type Program struct {
	start int
	end   int
}

type Programs []*Program

func bestArrange(programs Programs) int {
	sort.Sort(programs)
	// timeline表示来到的时间点
	timeLine := 0
	// result表示安排了多少个会议
	result := 0
	// 由于刚才按照结束时间排序，当前是按照谁结束时间早的顺序遍历
	for i := 0; i < len(programs); i++ {
		if timeLine <= programs[i].start {
			result++
			timeLine = programs[i].end
		}
	}
	return result
}

func (p Programs) Len() int {
	return len(p)
}

// Less 根据谁的结束时间早排序
func (p Programs) Less(i, j int) bool {
	return p[i].end-p[j].end > 0
}

func (p Programs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
