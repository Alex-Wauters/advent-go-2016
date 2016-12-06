package main

import (
	"fmt"
	"strings"
)

func part2(rooms []Room) {
	for _, room := range rooms {
		if room.Decrypt() == "northpole object storage" {
			fmt.Printf("The north pole has sector id: %v", room.sector)
			return
		}
	}
}

func (r Room) Decrypt() string {
	return strings.Map(func(ru rune) rune {
		if ru == ' ' {
			return ru
		}
		shifted := int(ru) + r.sector%26
		if shifted > 'z' {
			return rune(shifted - 26)
		}
		return rune(shifted)
	}, r.encrypted)
}
