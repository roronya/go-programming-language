package equal

import (
	"reflect"
	"unsafe"
)

func isCyclic(v reflect.Value, seen map[comparison]bool) bool {
	if v.CanAddr() {
		p := unsafe.Pointer(v.UnsafeAddr())
		c := comparison{p, v.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		return isCyclic(v.Elem(), seen)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isCyclic(v.Field(i), seen) {
				return true
			}
		}
		return false
	// Map, Slice, Arrayも判定すべき
	default:
		return false
	}
}

func IsCyclic(x interface{}) bool {
	seen := make(map[comparison]bool)
	return isCyclic(reflect.ValueOf(x), seen)
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}
