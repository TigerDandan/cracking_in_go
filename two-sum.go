package cracking_in_go

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for index, num := range nums {
		if v, found := hm[target-num]; found {
			return []int{v, index}
		}
		hm[num] = index
	}
	return nil
}
