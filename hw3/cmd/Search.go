package cmd

//бинарный поиск
func BinSearch(arr []int, need int) int {
	lowKey := 0
	highKey := len(arr) - 1
	var index int
	if arr[lowKey] > need || arr[highKey] < need {
		index = -1
	}
	for lowKey <= highKey {
		mid := (lowKey + highKey) / 2
		if arr[mid] == need {
			return mid
		}
		if arr[mid] < need {
			lowKey = mid + 1
			continue
		}
		if arr[mid] > need {
			highKey = mid - 1
		}
	}
	return index
}
