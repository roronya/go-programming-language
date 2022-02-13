package search

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Unpakcはreq内のHTTPリクエストパラメータから
// ptrが指す構造体のフィールドに値を移し替えます
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// 実行的な名前をキーとするフィールドのマップを構築する
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// リクエスト内の個々のパラメータに対する構造体のフィールドを更新
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // 認識されなかったHTTPパラメータを無視
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

func Pack(ptr interface{}) string {
	v := reflect.ValueOf(ptr).Elem()
	builder := strings.Builder{}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag
		name := tag.Get("http")
		if name == "" {
			continue
		}
		var value string
		switch v.Field(i).Kind() {
		case reflect.String:
			value = v.Field(i).String()
		case reflect.Int:
			value = strconv.Itoa(int(v.Field(i).Int()))
		case reflect.Bool:
			value = strconv.FormatBool(v.Field(i).Bool())
		}

		// 日本語に対応したいならBase64 encodeする
		if i == 0 {
			builder.WriteByte('?')
		} else {
			builder.WriteByte('&')
		}
		builder.WriteString(fmt.Sprintf("%s=%s", name, value))
	}
	return builder.String()
}
