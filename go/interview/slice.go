package interview

import "fmt"

// 这里前两次打印结果都是slice的底层数组里面的内容，s2追加两次后触发了扩容导致第二次扩容底层数组重新分配，所有200只会出现在s2，100会在s1
func SliceTestPointer() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	fmt.Printf("%p,%v,%v\r\n", s1, len(s1), cap(s1))
	s2 = append(s2, 100)
	fmt.Printf("%p,%v,%v\r\n", s2, len(s2), cap(s2))
	s2 = append(s2, 200)
	fmt.Printf("%p,%v,%v\r\n", s2, len(s2), cap(s2))

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

func SliceTestLenAndCap() {
	s := []int{1, 2, 3}
	fmt.Println(len(s), cap(s))
}

func SliceTestAppend() {
	var nums1 []interface{}
	nums2 := []int{1, 3, 4}
	num3 := append(nums1, nums2)
	fmt.Println(num3)
}
