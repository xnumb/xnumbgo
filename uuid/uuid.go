package uuid

import "github.com/gogf/gf/v2/util/guid"

func S(s ...string) string {
	r := ""
	if len(s) == 0 {
		r = guid.S()
	}
	if len(s) == 1 {
		r = guid.S([]byte(s[0]))
	}
	if len(s) == 2 {
		r = guid.S([]byte(s[0]), []byte(s[1]))
	}
	return r
}
