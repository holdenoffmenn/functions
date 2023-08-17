package functions

import (
	"reflect"
)

//MakeMapCopy will received a map
func CopyMapStringUint32(m map[string]uint32) map[string]uint32 {
	copy := make(map[string]uint32)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func CopyMapStringBool(m map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

/*
CopyMap takes a map of any format and returns its copy. 
It is necessary to inform the format of the map in the function call. 
Ex: map2 := CopyMap(map1).(map[string]string)
*/
func CopyMap(original interface{}) interface{} {
	originalValue := reflect.ValueOf(original)
	if originalValue.Kind() != reflect.Map {
		return nil
	}
	copy := reflect.MakeMap(originalValue.Type())

	for _, key := range originalValue.MapKeys() {
		originalElem := originalValue.MapIndex(key)
		copyElem := reflect.New(originalElem.Type()).Elem()
		copyElem.Set(originalElem)
		copy.SetMapIndex(key, copyElem)
	}

	return copy.Interface()
}
