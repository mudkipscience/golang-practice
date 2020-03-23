package main

import "math/rand"

func main() {
	println(ship("terra"))
	println(ship("terra", "emily"))
	println(ship("terra", "emily", "yeet"))
}

func cut(s string, l, r int) string {
	if l < 0 {
		l = 0
	}
	if r > len(s) {
		r = len(s)
	}
	if l == r {
		r++
	}
	if r < l {
		t := r
		r = l
		l = t
	}
	return s[l:r]
}

func roff() int {
	var num = rand.Intn(3) - 1
	println("num:", num)
	return num
}

func ship(names ...string) string {
	if len(names) == 0 {
		return ""
	}
	if len(names) == 1 {
		return names[0]
	}
	s := cut(names[0], 0, len(names[0])/2+roff())
	for i := 1; i < len(names)-1; i++ {
		s += cut(
			names[i],
			len(names[i])/3+roff(),
			(2*len(names[i]))/3+roff(),
		)
	}
	s += cut(
		names[len(names)-1],
		len(names[len(names)-1])/2+roff(),
		len(names[len(names)-1]),
	)
	return s
}
