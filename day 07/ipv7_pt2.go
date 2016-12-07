package main

import "strings"

func aba(seg string) (abas []string) {
	for i := 0; i < len(seg)-2; i++ {
		if isABA(seg[i : i+3]) {
			abas = append(abas, seg[i:i+3])
		}
	}
	return
}

func isABA(s string) bool {
	return s[0] == s[2] && s[0] != s[1]
}

func containsBab(seq string, aba string) bool {
	return strings.Index(seq, reverseAba(aba)) > -1
}

func reverseAba(aba string) string {
	return string(aba[1]) + string(aba[0]) + string(aba[1])
}

func (i ipv7) supportsSSL() bool {
	var abas []string
	for _, supernet := range i.supernet {
		abas = append(abas, aba(supernet)...)
	}
	for _, aba := range abas {
		for _, hypernet := range i.hypernet {
			if containsBab(hypernet, aba) {
				return true
			}
		}
	}
	return false
}
