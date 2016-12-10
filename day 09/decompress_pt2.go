package main

import (
	"fmt"
	"time"
)

func part2(input string) {
	defer track(time.Now(), "part 2")
	fmt.Println(decompress2(input))
}

type multiplier struct {
	len, repeat int
}

type repeater struct {
	multipliers []*multiplier
}

func decompress2(s string) (count uint64) {
	u := unpack{}
	r := repeater{}
	for _, c := range s {
		switch {
		case u.inCommand && isNumber(c):
			u.setDigit(c)
		case c == ')':
			r.Add(&multiplier{u.length, u.repeat})
			u = unpack{}
		case c == '(':
			u.inCommand = true
		case c == 'x' && u.inCommand:
			u.lengthSet = true
		case c == ' ':
		default:
			count += uint64(r.Count())
		}
		r.Decrease()
	}
	return
}

func (r *repeater) Add(m *multiplier) {
	r.multipliers = append(r.multipliers, m)
}

func (r *repeater) Count() int {
	i := 1
	for _, m := range r.multipliers {
		i = i * m.repeat
	}
	return i
}

func (r *repeater) Decrease() {
	for i := len(r.multipliers) - 1; i >= 0; i-- {
		m := r.multipliers[i]
		m.len = m.len - 1
		if m.len < 0 {
			r.multipliers = append(r.multipliers[:i], r.multipliers[i+1:]...)
		}
	}
}
