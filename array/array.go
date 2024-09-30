package array

func MergeArray(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	curr := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[curr] = nums1[i]
			i--
		} else {
			nums1[curr] = nums2[j]
			j--
		}
		curr--
	}
	for j >= 0 {
		nums1[curr] = nums2[j]
		j--
		curr--
	}
}

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	tmpCount := 1
	curr := 1
	newTail := 0
	for curr < len(nums) {
		if nums[curr] != nums[newTail] {
			tmpCount = 1
			newTail++
			nums[newTail] = nums[curr]

		} else if tmpCount < 2 {
			tmpCount++
			newTail++
			nums[newTail] = nums[curr]
		}
		curr++
	}
	return newTail + 1
}

func ReverseArray(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	//fmt.Println(nums)
}

func RotateArray(nums []int, c int) {
	if c == 0 {
		return
	}
	t := c % len(nums)
	ReverseArray(nums[0 : len(nums)-t])
	ReverseArray(nums[len(nums)-t:])
	ReverseArray(nums)
}
