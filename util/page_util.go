package util

// MaxPage 通过总数和页面最大数，计算出最大页数
func MaxPage(sum, pageNum int) int {
	if sum <= 0 {
		return 1
	}
	page := sum / pageNum
	if sum%pageNum != 0 {
		page = page + 1
	}
	return page
}
