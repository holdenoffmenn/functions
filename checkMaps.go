package functions

import (
	"reflect"
)

// func MapsCheck(map1, map2 map[string]string) bool {
// 	if len(map1) != len(map2) {
// 		return true
// 	}
// 	for key, val1 := range map1 {
// 		val2, ok := map2[key]
// 		if !ok || val1 != val2 {
// 			return true
// 		}
// 	}
// 	return false
// }

/*
CompareMaps receives two maps that must be of the same 
format, and checks if they are exactly the same, returning 
a bool value.
*/
func CompareMaps(map1, map2 interface{}) bool {
	if reflect.TypeOf(map1) != reflect.TypeOf(map2) {
		return false
	}

	map1Value := reflect.ValueOf(map1)
	map2Value := reflect.ValueOf(map2)

	if map1Value.Len() != map2Value.Len() {
		return false
	}

	for _, key := range map1Value.MapKeys() {
		if !map2Value.MapIndex(key).IsValid() {
			return false
		}

		map1Elem := map1Value.MapIndex(key).Interface()
		map2Elem := map2Value.MapIndex(key).Interface()

		if !reflect.DeepEqual(map1Elem, map2Elem) {
			return false
		}
	}

	return true
}