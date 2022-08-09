package mem

import "testing"

const times = 100000

func BenchamarkByRef(b *testing.B) {
	s := foo{a: "somevalue"}
	concat := ""
	for i := 0; i < times; i++ {
		concat += printer(&s)
	}
}

func BenchmarkByVal(b *testing.B) {
	s := foo{a: "somevalue"}
	concat := ""
	for i := 0; i < times; i++ {
		concat += printer(s)
	}
}
