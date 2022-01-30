package display

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	// 完結にするために浮動小数点と複素数は省略
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " Ox" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		b := strings.Builder{}
		b.WriteString("{\n")
		for i := 0; i < v.NumField(); i++ {
			fieldPath := v.Type().Field(i).Name
			field := formatAtom(v.Field(i))
			b.WriteString(fmt.Sprintf("\t%s:\t%s\n", fieldPath, field))
		}
		b.WriteString("}")
		return b.String()
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
