package util

func MaxPage(sum, pageNum int) int {
	page := sum / pageNum
	if sum%pageNum != 0 {
		page = page + 1
	}
	return page
}
