package main

func f() (r string) {
	defer func() {
		if p := recover(); p != nil {
			r = "recovered"
		}
	}()
	panic("panic")
}
