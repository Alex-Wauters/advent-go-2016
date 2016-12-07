package main

import (
	"reflect"
	"testing"
)

func TestToIp(t *testing.T) {
	line := "efuoofbuyjoaqjd[achnmlslfvovmgt]xcuyvikslsewgqlx[gjxolnhgqhhglojjqhy]iarxidejlgphqwaei"
	ip := toIp(line)
	expected := ipv7{
		[]string{"efuoofbuyjoaqjd", "xcuyvikslsewgqlx", "iarxidejlgphqwaei"}, []string{"achnmlslfvovmgt", "gjxolnhgqhhglojjqhy"},
	}
	if !reflect.DeepEqual(ip, expected) {
		t.Error("For", line, "expected", expected, "got", ip)
	}
}

func TestIsValidAbbaSegment(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{"abba[mnop]qrst", true},
		{"abcd[bddb]xyyx", false},
		{"aaaa[qwer]tyui", false},
		{"ioxxoj[asdfgh]zxcvbn", true},
		{"ioxxoj[asdfgh]zxcvbn[abba]", false},
	}
	for _, test := range tests {
		supportsAbba := toIp(test.input).supportsAbba()
		if supportsAbba != test.result {
			t.Error("For", test.input, "expected", test.result, "got", supportsAbba)
		}
	}
}

func TestIsValidSSLSegment(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{"aba[bab]xyz", true},
		{"xyx[xyx]xyx", false},
		{"aaa[kek]eke", true},
		{"zazbz[bzb]cdb", true},
	}
	for _, test := range tests {
		supportsSSL := toIp(test.input).supportsSSL()
		if supportsSSL != test.result {
			t.Error("For", test.input, "expected", test.result, "got", supportsSSL)
		}
	}
}
