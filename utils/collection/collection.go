package collection

import (
	"golang.org/x/exp/constraints"
	"sort"
	"strings"
)

// SupplementarySet 求差集(在切片1中删除切片2中的元素)
func SupplementarySet(slice1, slice2 []int) []int {
	if len(slice1) == 0 {
		return slice1
	}
	if len(slice2) == 0 {
		return DeduplicationIntList(slice1)
	}
	m := make(map[int]int)
	for _, v := range slice1 {
		m[v] = v
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			delete(m, v)
		}
	}
	var str []int
	for _, s2 := range m {
		str = append(str, s2)
	}
	return str
}

// DeduplicationIntList int数组去重
func DeduplicationIntList(req []int) (response []int) {
	intMap := make(map[int]bool)
	for _, data := range req {
		if !intMap[data] {
			intMap[data] = true
			response = append(response, data)
		}
	}
	return
}

// DeduplicationIntList64 int数组去重
func DeduplicationIntList64(req []int64) (response []int64) {
	intMap := make(map[int64]bool)
	for _, data := range req {
		if !intMap[data] {
			intMap[data] = true
			response = append(response, data)
		}
	}
	return
}

// DeduplicationStringList string数组去重
func DeduplicationStringList(req []string) (response []string) {
	intMap := make(map[string]bool)
	for _, data := range req {
		if !intMap[data] {
			intMap[data] = true
			response = append(response, data)
		}
	}
	return
}

// RemoveRepByIntMap 整数组去重 通过map主键唯一的特性过滤重复元素
func RemoveRepByIntMap(list []int) (result []int) {
	tempMap := map[int]int{}
	for _, e := range list {
		l := len(tempMap)
		tempMap[e] = e
		// 加入map后，map长度变化，则元素不重复
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// ContainInt 判断是否包含
func ContainInt(arr []int, i int) bool {
	if len(arr) == 0 {
		return false
	}
	for _, a := range arr {
		if a == i {
			return true
		}
	}
	return false
}

// 判断是否包含string
func ContainString(arr []string, i string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, a := range arr {
		if a == i {
			return true
		}
	}
	return false
}

// GetDuplicateInt 求重复的int值
func GetDuplicateInt(list1 []int, list2 []int) []int {
	var list []int
	for _, i1 := range list1 {
		if ContainInt(list2, i1) {
			list = append(list, i1)
		}
	}
	return list

}

// RemoveEmptyAndBlankString 去掉字符串切片中的空字符 []string{""," ","   "} ==>[]string{}
func RemoveEmptyAndBlankString(list *[]string) {
	var l []string
	if len(*list) == 0 {
		return
	}
	for _, str := range *list {
		if len(strings.TrimSpace(str)) == 0 {
			continue
		}
		l = append(l, str)
	}
	*list = l
}

func RemoveZeroInt(list *[]int) {
	var l []int
	if len(*list) == 0 {
		return
	}
	for _, i := range *list {
		if i == 0 {
			continue
		}
		l = append(l, i)
	}
	*list = l
}

// AppendToSliceIfAbsent 当item不在sl中存在时才append到sl尾
func AppendToSliceIfAbsent(sl []string, item string) []string {
	for _, ele := range sl {
		if ele == item {
			return sl
		}
	}
	sl = append(sl, item)
	return sl

}

// AppendToIntSliceIfAbsentAndSkipZero 当item不在sl中存在时才append到sl尾
func AppendToIntSliceIfAbsentAndSkipZero(sl []int, item int) []int {
	if item == 0 {
		return sl
	}
	for _, ele := range sl {
		if ele == item {
			return sl
		}
	}
	sl = append(sl, item)
	return sl

}

// DeduplicationStringListAndTrim 去重并Trim掉
func DeduplicationStringListAndTrim(req []string, itemToTrim ...string) (response []string) {
	intMap := make(map[string]bool)
	for _, data := range req {
		if ContainsStr(itemToTrim, data) {
			continue
		}
		if !intMap[data] {
			intMap[data] = true
			response = append(response, data)
		}
	}
	return
}

func DeduplicationIntListAndTrim(req []int, itemToTrim ...int) (response []int) {
	intMap := make(map[int]bool)
	for _, data := range req {
		if ContainsInt(itemToTrim, data) {
			continue
		}
		if !intMap[data] {
			intMap[data] = true
			response = append(response, data)
		}
	}
	return
}

func ContainsInt(items []int, ele int) bool {
	for _, item := range items {
		if item == ele {
			return true
		}
	}
	return false
}

func ContainsStr(items []string, ele string) bool {
	if len(items) == 0 {
		return false
	}

	for _, item := range items {
		if item == ele {
			return true
		}
	}
	return false
}

// Intersection 求交集
func Intersection[T constraints.Ordered](a, b []T) (ret []T) {
	for _, itemA := range a {
		for _, itemB := range b {
			if itemA == itemB {
				ret = append(ret, itemB)
			}
		}
	}
	return
}

func IntersectionInt(a, b []int) []int {
	m := make(map[int]int, 0)
	for _, v := range a {
		m[v] += 1
	}
	count := 0
	for _, v := range b {
		if m[v] > 0 {
			m[v] = 0
			a[count] = v
			count++
		}
	}
	return a[:count]
}

func ConvertInt64ToInt(a []int64) []int {
	res := make([]int, 0)
	if len(a) == 0 {
		return res
	}
	for _, v := range a {
		res = append(res, int(v))
	}
	return res
}

// CompareStrSliceIgnoreSort 比较两个string切片是否内容相同，忽略内部顺序
func CompareStrSliceIgnoreSort(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 1 {
		return s1[0] == s2[0]
	}
	sort.Strings(s1)
	sort.Strings(s2)
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// CompareIntSliceIgnoreSort 比较两个Int切片是否内容相同，忽略内部顺序
func CompareIntSliceIgnoreSort(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 1 {
		return s1[0] == s2[0]
	}
	sort.Ints(s1)
	sort.Ints(s2)
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func HaveSameElementAtLeastOne(arr1 []string, arr2 []string) bool {
	for _, a := range arr2 {
		for _, m := range arr1 {
			if a == m {
				return true
			}
		}
	}
	return false
}
