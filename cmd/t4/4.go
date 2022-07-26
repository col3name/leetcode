package t4

func findMedianSortedArrays(arr1 []int, arr2 []int) float64 {
	m := len(arr1)
	n := len(arr2)
	arr3 := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if arr1[i] < arr2[j] || arr1[i] <= arr2[j] {
			arr3[k] = arr1[i]
			i++
			k++
		} else {
			arr3[k] = arr2[j]
			j++
			k++
		}
	}
	for i < m {
		arr3[k] = arr1[i]
		i++
		k++
	}
	for j < n {
		arr3[k] = arr2[j]
		j++
		k++
	}

	len3 := len(arr3)
	if len3%2 == 0 {
		i2 := arr3[len3/2]
		i3 := arr3[(len3/2)-1]
		return float64(i2+i3) / 2.0
	}
	return float64(arr3[(len3 / 2)])
}

//func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
//	len1 := len(nums1)
//	len2 := len(nums2)
//	nums := make([]int, len1+len2)
//	lhs, rhs := 0, 0
//	for lhs < len1 || rhs < len2 {
//		i := math.MaxInt64
//		j := math.MaxInt64
//		if lhs < len1 {
//			i = nums1[lhs]
//		}
//		if rhs < len2 {
//			j = nums2[rhs]
//		}
//		if i < j {
//			i++
//		} else {
//			j++
//		}
//	}
//
//	if len(nums)%2 == 0 {
//		return float64((nums[len(nums)/2-1] + nums[len(nums)/2]) / 2)
//	}
//
//	return float64(nums[len(nums)/2-1])
//}
