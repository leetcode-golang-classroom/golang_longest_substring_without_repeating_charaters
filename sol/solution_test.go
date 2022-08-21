package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "abcabcbb"
	for idx := 0; idx < b.N; idx++ {
		lengthOfLongestSubstring(s)
	}
}
func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "s = \"abcabcbb\"",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "s = \"bbbbb\"",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "s = \"pwwkew\"",
			args: args{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
