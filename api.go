package functions

type Client interface {

	CompareMapsStrIFace(map1, map2 map[string]interface{})

	CompareMapsStrBool(map1, map2 map[string]bool) bool

	CompareMapsStrUint32(map1, map2 map[string]uint32) bool

	CompareMapsIFaceIFace(map1, map2 interface{}) bool

	CopyMapStringUint32(m map[string]uint32) map[string]uint32

	CopyMapStringBool(m map[string]bool) map[string]bool

	CopyMap(original interface{}) interface{}
}
