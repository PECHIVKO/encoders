package format

import (
	"testing"
)

func TestRemoveUnexpectedSymbols(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"test1", args{"Owl, EGG I, ."}, "Owl EGG I"},
		{"test2", args{"./.hono/123ur pig!!! laTIN "}, "honour pig laTIN"},
		{"test3", args{"  That's a test    String"}, "That's a test String"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := RemoveUnexpectedSymbols(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("RemoveUnexpectedSymbols() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestFormatOutput(t *testing.T) {
	type args struct {
		reference string
		input     string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"test1", args{"Owl", "parrot"}, "Parrot"},
		{"test2", args{"OWL", "PaRRot"}, "PARROT"},
		{"test3", args{"owl", "PARROT"}, "parrot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := FormatOutput(tt.args.reference, tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("FormatOutput() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_isUpperCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"BANANA"}, true},
		{"test2", args{"banana"}, false},
		{"test3", args{"Banana"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUpperCase(tt.args.s); got != tt.want {
				t.Errorf("isUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startsWithCapital(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"BANANA"}, false},
		{"test2", args{"banana"}, false},
		{"test3", args{"Banana"}, true},
		{"test4", args{"BaNaNa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startsWithCapital(tt.args.s); got != tt.want {
				t.Errorf("startsWithCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
