package main

import "fmt"

// removeDuplicateDataByKey removes duplicates from a slice of map[string]interface{} by the given key
func removeDuplicateDataByKey(arr []map[string]interface{}, key string) []map[string]interface{} {
	seen := make(map[interface{}]bool)
	result := []map[string]interface{}{}

	for _, item := range arr {
		if val, ok := item[key]; ok {
			if !seen[val] {
				seen[val] = true
				result = append(result, item)
			}
		} else {
			// If the key does not exist, include the item
			result = append(result, item)
		}
	}

	return result
}

func main() {
	input1 := []map[string]interface{}{
		{"id": 1, "name": "A"},
		{"id": 2, "name": "B"},
		{"id": 1, "name": "A"},
		{"id": 3, "name": "C"},
	}

	input3 := []map[string]interface{}{
		{"id": 1}, {"id": 2}, {"id": 3},
	}

	input4 := []map[string]interface{}{
		{"id": 1, "name": "A"},
		{"id": 2, "name": "A"},
		{"id": 3, "name": "B"},
	}

	fmt.Println(removeDuplicateDataByKey(input1, "id"))
	fmt.Println(removeDuplicateDataByKey([]map[string]interface{}{}, "id"))
	fmt.Println(removeDuplicateDataByKey(input3, "id"))
	fmt.Println(removeDuplicateDataByKey(input4, "name"))
}
