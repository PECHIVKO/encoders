package encoder

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"test1", args{"Owl, EGG I, ."}, "Owlyay EGGYAY IYAY"},
		{"test2", args{"./.hono/123ur pig!!! laTIN "}, "honouryay igpay atinlay"},
		{"test3", args{"  That's a test    String"}, "at'sthay ayay esttay Ingstray"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Encode(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("Encode() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_vowelsPigLatin(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"test1", args{"Owl"}, "Owlyay"},
		{"test2", args{"EGG"}, "EGGYAY"},
		{"test3", args{"I"}, "IYAY"},
		{"test4", args{"honour"}, "honouryay"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := vowelsPigLatin(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("vowelsPigLatin() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_consonantsPigLatin(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"test1", args{"PIG"}, "IGPAY"},
		{"test2", args{"Latin"}, "Atinlay"},
		{"test3", args{"That's"}, "at'sthay"},
		{"test4", args{"StranGECasE"}, "angecasestray"},
		{"test5", args{"lowercase"}, "owercaselay"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := consonantsPigLatin(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("consonantsPigLatin() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_isConsostant(t *testing.T) {
	type args struct {
		character rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"vowel 1", args{'a'}, false},
		{"vowel 2", args{'o'}, false},
		{"consostant", args{'b'}, true},
		{"consostant", args{'y'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isConsostant(tt.args.character); got != tt.want {
				t.Errorf("isConsostant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startsWithVowelSound(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"starts with vowel 1", args{"honour"}, true},
		{"starts with vowel 2", args{"owl"}, true},
		{"starts with consostant", args{"ball"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startsWithVowelSound(tt.args.s); got != tt.want {
				t.Errorf("startsWithVowelSound() = %v, want %v", got, tt.want)
			}
		})
	}
}
