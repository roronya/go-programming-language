package sexpr

import (
	"testing"
)

func TestBool(t *testing.T) {
	var got bool
	err := Unmarshal([]byte("true"), &got)
	if err != nil {
		t.Error(err)
	}
	if got != true {
		t.Error("want true, got false")
	}

	err = Unmarshal([]byte("false"), &got)
	if err != nil {
		t.Error(err)
	}
	if got != false {
		t.Error("want false, got true")
	}

}

func TestFloat64(t *testing.T) {
	tests := []struct {
		want float64
		in   []byte
	}{
		{1.000, []byte("1.000")},
		{6.02214129e23, []byte("6.02214129e+23")},
		{6.62606957e-34, []byte("6.62606957e-34")},
	}
	for _, test := range tests {
		var got float64
		err := Unmarshal(test.in, &got)
		if err != nil {
			t.Error(err)
		}
		if test.want != got {
			t.Errorf("want %g, got %g", test.want, got)
		}
	}
}

func TestComplex(t *testing.T) {
	tests := []struct {
		want complex128
		in   []byte
	}{
		{complex(1, 2), []byte("#(1 2)")},
		{complex(1.1, 2.1), []byte("#(1.1 2.1)")},
		{complex(0, 0), []byte("#(0 0)")},
		{complex(1, 0), []byte("#(1 0)")},
		{complex(0, 1), []byte("#(0 1)")},
	}
	for _, test := range tests {
		var got complex128
		err := Unmarshal(test.in, &got)
		if err != nil {
			t.Error(err)
		}
		if got != test.want {
			t.Errorf("want %v, got %v", test.want, got)
		}
	}
}

func TestInterface(t *testing.T) {
	var x interface{} = 1
	var y interface{} = []int{1, 2, 3}
	var z interface{} = struct {
		a int
		b byte
	}{1, 'a'}
	tests := []struct {
		want interface{}
		in   []byte
	}{
		{&x, []byte("(int 1)")},
		{&y, []byte("([]int [1 2 3])")},
		{&z, []byte("(struct { a int; b uint8 } {1 97})")},
	}
	for _, test := range tests {
		var got interface{}
		err := Unmarshal(test.in, &got)
		if err != nil {
			t.Error(err)
		}
		if got != test.want {
			t.Errorf("want %v, got %v", test.want, got)
		}
	}
}
