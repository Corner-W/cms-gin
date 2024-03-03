package collection

import "golang.org/x/exp/constraints"

// SliceInter 求交集
func SliceInter[T constraints.Ordered](slice1, slice2 []T) []T {
	m := make(map[T]struct{})
	ret := make([]T, 0)

	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			ret = append(ret, v)
		}
	}
	return ret
}

// SliceUnion 求并集
func SliceUnion[T constraints.Ordered](slice1, slice2 []T) []T {
	m := make(map[T]struct{})

	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		if _, ok := m[v]; !ok {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// SliceDiff 求差集(slice1-并集)
func SliceDiff[T constraints.Ordered](slice1, slice2 []T) []T {
	m := make(map[T]struct{})
	ret := make([]T, 0)

	inter := SliceInter(slice1, slice2)
	for _, v := range inter {
		m[v] = struct{}{}
	}
	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			ret = append(ret, v)
		}
	}
	return ret
}

// SliceComple 求补集
func SliceComple[T constraints.Ordered](slice1, slice2 []T) []T {
	m := make(map[T]int)
	ret := make([]T, 0)

	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		m[v]++
	}
	for value, num := range m {
		if num == 1 {
			ret = append(ret, value)
		}
	}
	return ret
}

// SlicesEqual 判断是否相等
func SlicesEqual[T constraints.Ordered](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// SliceUnique 去重
func SliceUnique[T constraints.Ordered](slice []T) []T {
	m := map[T]struct{}{}
	ret := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			ret = append(ret, v)
		}
	}
	return ret
}

// SliceContain 判断是否包含
func SliceContain[T constraints.Ordered](slice []T, val T) bool {
	if len(slice) == 0 {
		return false
	}
	for _, a := range slice {
		if a == val {
			return true
		}
	}
	return false
}

// SliceConvert 转换类型
func SliceConvert[T2, T constraints.Integer](s []T) (res []T2) {
	for _, v := range s {
		res = append(res, T2(v))
	}
	return
}
