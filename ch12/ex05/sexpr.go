package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

/**
ref RFC 8259 https://www.rfc-editor.org/rfc/pdfrfc/rfc8259.txt.pdf
*/
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Array, reflect.Slice: // [ value ... ]
		buf.WriteString("[ ")
		for i := 0; i < v.Len(); i++ {
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}

			if i == v.Len()-1 {
				buf.WriteByte(' ')
			} else {
				buf.WriteString(", ")
			}
		}
		buf.WriteByte(']')
	case reflect.Struct: // { "key": "value", ...}
		buf.WriteString("{ ")
		for i := 0; i < v.NumField(); i++ {
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}

			if i == v.NumField()-1 {
				buf.WriteByte(' ')
			} else {
				buf.WriteString(", ")
			}
		}
		buf.WriteByte('}')
	case reflect.Map: // { "key": "value", ...}
		buf.WriteString("{ ")
		for i, key := range v.MapKeys() {
			if key.Kind() != reflect.String {
				return fmt.Errorf("non string map key is unsupported")
			}
			fmt.Fprintf(buf, "%q: ", key.String())
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}

			if i == v.Len()-1 {
				buf.WriteByte(' ')
			} else {
				buf.WriteString(", ")
			}
		}
		buf.WriteByte('}')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
