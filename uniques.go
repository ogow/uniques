package uniques

import (
	"sort"
)

func StringSlice(ss []string) uniqueStrings {
	keys := make(map[string]bool)
	var uniques uniqueStrings

	for _, item := range ss {
		if _, value := keys[item]; !value {
			keys[item] = true
			uniques = append(uniques, item)
		}
	}
	return uniques
}

func IntSlice(ss []int) uniqueInts {
	keys := make(map[int]bool)
	var uniques uniqueInts

	for _, item := range ss {
		if _, value := keys[item]; !value {
			keys[item] = true
			uniques = append(uniques, item)
		}
	}
	return uniques
}

func NewItems(oldSlice, newSlice []string) []string {
	oldMap := make(map[string]bool)
	var newi []string

	for _, ele := range oldSlice {
		oldMap[ele] = true
	}

	for _, el := range newSlice {
		if _, ok := oldMap[el]; !ok {
			newi = append(newi, el)
		}
	}

	return newi
}

type uniqueStrings []string

func (us uniqueStrings) Sort() []string {
	sort.Strings(us)
	return us
}

type uniqueInts []int

func (ui uniqueInts) Sort() []int {
	sort.Ints(ui)
	return ui
}
