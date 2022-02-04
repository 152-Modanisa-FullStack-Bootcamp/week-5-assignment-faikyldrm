package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	var sumResult uint64
	sumResult = uint64(x) + uint64(y)
	if sumResult > math.MaxUint32 {
		result := (sumResult - math.MaxUint32) - 1
		return uint32(result), true
	} else {
		return uint32(sumResult), false
	}

}

func CeilNumber(f float64) float64 {
	ceilNumber := 0.25
	_, fraction := math.Modf(f)
	modOfFraction := math.Mod(fraction, ceilNumber)
	if modOfFraction == 0 {
		return f
	}
	if f < 0 {
		modOfFraction += ceilNumber
	}
	diff := ceilNumber - (modOfFraction)

	result := f + diff
	return result
}

func AlphabetSoup(s string) string {
	spliced := strings.Split(s, "")
	sort.Strings(spliced)
	return strings.Join(spliced, "")
}

func StringMask(s string, n uint) string {
	spliced := strings.Split(s, "")
	if cap(spliced) == 0 {
		spliced = make([]string, 1)
		spliced[0] = "*"
	}
	if n >= uint(len(spliced)) || len(spliced) == 0 {
		n = 0
	}
	for i := n; i < uint(len(spliced)); i++ {
		spliced[i] = "*"
	}
	return strings.Join(spliced, "")
}

func WordSplit(arr [2]string) string {
	spliced := strings.Split(arr[1], ",")
	founded := ""
	i:=0
	for _, s := range spliced {
		i++
		if strings.Contains(arr[0], s) {
			sFirst := s + founded
			fFirst := founded + s
			if sFirst == arr[0] ||fFirst == arr[0]{

				return strings.Join([]string{s, founded}, ",")
			}
			founded = s
		}
	}
	fmt.Println(i)
	return "not possible"
}

func VariadicSet(i ...interface{}) []interface{} {
	allKeys := make(map[interface{}]bool)
	var list []interface{}
	for _, item := range i {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
