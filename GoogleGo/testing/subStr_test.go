package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是一个测试用例", 9},
		{"一二三二一", 3},
		{"这是在学习Golang", 11},
		{"黑化肥发黑会发挥灰化肥会挥发", 7},
	}

	for _, s := range tests {
		actual := lengthOfNonRepeatingSubStr(s.s)
		if actual != s.ans {
			t.Errorf("<func:lengthOfNonRepeatingSubStr>: args=%s, get=%d, expected=%d", s.s, actual, s.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥发黑会发挥灰化肥会挥发"
	ans := 7

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
