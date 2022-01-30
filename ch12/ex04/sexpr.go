package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)
	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		indent++

		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				buf.Write(bytes.Repeat([]byte{' '}, indent))
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}

		buf.WriteByte(')')
		indent--
	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		indent++

		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				buf.Write(bytes.Repeat([]byte{' '}, indent))
			}

			fieldName := v.Type().Field(i).Name
			fmt.Fprintf(buf, "(%s ", fieldName)
			indent += 1 + len(fieldName) + 1 // 括弧とfieldNameの文字列とスペース分をインデントする

			if err := encode(buf, v.Field(i), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
			indent -= len(fieldName) + 1 + 1 // indentを戻す
		}

		buf.WriteByte(')')
		indent--
	case reflect.Map: // (key value) ...)
		buf.WriteByte('(')
		indent++
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')
				buf.Write(bytes.Repeat([]byte{' '}, indent))
			}
			buf.WriteByte('(')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
		indent--
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "t")
		} else {
			fmt.Fprintf(buf, "nil")
		}
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#(%g %g)", real(c), imag(c))
	case reflect.Interface:
		fmt.Fprintf(buf, "(%s %v)", v.Elem().Type().String(), v.Elem())
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
