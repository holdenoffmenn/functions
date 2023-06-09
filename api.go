package functions

type Client interface {

	//MapsCheck(map1, map2 map[string]string) bool

	CompareMapsIFaceIFace(map1, map2 interface{}) bool

	// CopyMapStringString(m map[string]string) map[string]string

	// CopyMapStringInt(m map[string]int64) map[string]int64

	CopyMap(original interface{}) interface{}

	CompareMapsStrIFace(map1, map2 map[string]interface{}) bool
}
