package mymap

import "slices"

func MergeToMap(data map[string][]string, newValues []string, key string){
	if _, ok := data[key]; ok {
		for _, value := range newValues {
			if !slices.Contains(data[key], value) {
				data[key] = append(data[key], value)
			}
		}
		return
	}
	data[key]=newValues
}
