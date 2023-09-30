package study

import (
	"fmt"
)

func TestSlice() {
	i1, i2, i3, i4 := 1, 2, 3, 4
	arr := [4]*int{&i1, &i2, &i3, &i4}
	s1 := arr[:]
	s2 := make([]*int, 4, 6)
	copy(s2, s1)
	i5, i6, i7 := 5, 6, 7
	s1 = append(s1, &i5, &i6)
	s2 = append(s2, &i7) //append会对切片动态扩容
	s3 := make([]*int, len(s2), cap(s2)*2)
	copy(s3, s2)
	fmt.Printf("arr type=%T, mem=%v, value=%v, len=%d, cap=%d\n", arr, &arr, arr, len(arr), cap(arr))
	fmt.Printf("s1  type=%T, mem=%v, value=%v, len=%d, cap=%d\n", s1, &s1, s1, len(s1), cap(s1))
	fmt.Printf("s2  type=%T, mem=%v, value=%v, len=%d, cap=%d\n", s2, &s2, s2, len(s2), cap(s2))
	fmt.Printf("s3  type=%T, mem=%v, value=%v, len=%d, cap=%d\n", s3, &s3, s3, len(s3), cap(s3))
}
