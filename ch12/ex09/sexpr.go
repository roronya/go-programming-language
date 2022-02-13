package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
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
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map: // (key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default: // float, complex, bool, chan, func, interface
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

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // 最初のトークンを取得する
	defer func() {
		// note: This is not an ideal error handling example.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

type Decoder struct {
	lex   *lexer
	stack []Token
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &Decoder{lex, []Token{}}
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	dec.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()
	read(dec.lex, reflect.ValueOf(v).Elem())
	return nil
}

func (dec *Decoder) Token() (token Token, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()

	dec.lex.next()
	err = dec.token()
	return dec.pop(), err
}

func (dec *Decoder) push(token Token) {
	dec.stack = append(dec.stack, token)
}

func (dec *Decoder) pop() Token {
	top := dec.stack[0]
	if len(dec.stack) > 2 {
		for i := 0; i < len(dec.stack)-1; i++ {
			dec.stack[i] = dec.stack[i+1]
		}
	}
	dec.stack = dec.stack[:len(dec.stack)-1]
	return top
}

func (dec *Decoder) token() error {
	switch dec.lex.token {
	case scanner.EOF:
		return fmt.Errorf("EOF")
	case scanner.Ident:
		if dec.lex.text() == "nil" {
			dec.push(Symbol("nil"))
		} else { // Key of Struct
			dec.push(Symbol(dec.lex.text()))
		}
		dec.lex.next()
	case scanner.String:
		s, _ := strconv.Unquote(dec.lex.text()) // ignore a error
		dec.push(String(s))
		dec.lex.next()
	case scanner.Int:
		i, _ := strconv.Atoi(dec.lex.text()) // ignore a error
		dec.push(Int(i))
		dec.lex.next()
	case '(':
		dec.push(StartList('('))
		dec.lex.next() // consume '('
		for !endList(dec.lex) {
			dec.token()
		}
		dec.push(EndList(')'))
		dec.lex.next() // consume ')'
	}
	return nil
}

type Token interface{}
type Symbol string
type String string
type Int int
type StartList byte
type EndList byte
