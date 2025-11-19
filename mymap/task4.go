package mymap

import (
	"slices"
)

func MergeToMap(DataFrame map[string][]string, newValues []string, key string){
	if _, ok := DataFrame[key]; ok {
		for _, value := range newValues {
			if !slices.Contains(DataFrame[key], value) {
				DataFrame[key] = append(DataFrame[key], value)
			}
		}
		return
	}
	DataFrame[key]=newValues
}

