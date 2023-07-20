package main

import "fmt"

func main() {
	s := append([]byte(`"t\e\st"`), 0)

	read := newRead(s)
	var ok bool = true
	for ok {
		ok = switchState(read)
	}
}

func newRead(s []byte) func() byte {
	i := 0

	return func() byte {
		defer func() {
			i++
		}()

		return s[i]
	}
}

func switchState(read func() byte) bool {
	state := 0

	for {
		c := read()

		switch state {
		case 0:
			if c != '"' {
				return false
			}
			state = 1
		case 1:
			if c == '"' {
				return true
			}
			if c == '\\' {
				state = 2
			} else {
				state = 1
			}
		case 2:
			fmt.Printf("%c", c)
			state = 1
		}
	}
}
