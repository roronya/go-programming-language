package search

import (
	"fmt"
	"net/http"
	mail2 "net/mail"
	"reflect"
	"strconv"
	"strings"
)

type Param struct {
	value reflect.Value
	mail  bool
}

// Unpakcはreq内のHTTPリクエストパラメータから
// ptrが指す構造体のフィールドに値を移し替えます
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// 実行的な名前をキーとするフィールドのマップを構築する
	fields := make(map[string]Param)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		mail, err := strconv.ParseBool(tag.Get("mail"))
		if err != nil {
			return err
		}
		fields[name] = Param{v.Field(i), mail}
	}

	// リクエスト内の個々のパラメータに対する構造体のフィールドを更新
	for name, values := range req.Form {
		f := fields[name].value
		if !f.IsValid() {
			continue // 認識されなかったHTTPパラメータを無視
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if fields[name].mail {
					if err := validateMail(value); err != nil {
						return err
					}
				}
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if fields[name].mail {
					if err := validateMail(value); err != nil {
						return err
					}
				}
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

func validateMail(mail string) error {
	_, err := mail2.ParseAddress(mail)
	return err
}
