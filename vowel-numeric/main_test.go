package main

import (
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		inputPhrase string
	}
	tests := []struct {
		name             string
		args             args
		wantOutputPhrase string
	}{
		{"test1", args{"BanAna"}, "B1n1n1"},
		{"test2", args{"bAanEeNiI goO sUudo"}, "b11n22N33 g44 s55d4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputPhrase := encode(tt.args.inputPhrase); gotOutputPhrase != tt.wantOutputPhrase {
				t.Errorf("encode() = %v, want %v", gotOutputPhrase, tt.wantOutputPhrase)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		inputPhrase string
	}
	tests := []struct {
		name             string
		args             args
		wantOutputPhrase string
	}{
		{"test1", args{"B1n1n1"}, "Banana"},
		{"test2", args{"b11n22N33 g44 s55d4"}, "baaneeNii goo suudo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputPhrase := decode(tt.args.inputPhrase); gotOutputPhrase != tt.wantOutputPhrase {
				t.Errorf("decode() = %v, want %v", gotOutputPhrase, tt.wantOutputPhrase)
			}
		})
	}
}

func Test_getPhrase(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name       string
		args       args
		wantPhrase string
	}{
		{"test1", args{"encode(somephrase)"}, "somephrase"},
		{"test2", args{"decode(somephrase)"}, "somephrase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPhrase := getPhrase(tt.args.command); gotPhrase != tt.wantPhrase {
				t.Errorf("getPhrase() = %v, want %v", gotPhrase, tt.wantPhrase)
			}
		})
	}
}

func Test_isValidCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"encode(somephrase)"}, true},
		{"test2", args{"decode(somephrase)"}, true},
		{"test3", args{"encode(\"somephrase\")"}, true},
		{"test4", args{"code(somephrase)"}, false},
		{"test4", args{"encode(somephrase"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCommand(tt.args.command); got != tt.want {
				t.Errorf("isValidCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputHandler(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"test1", args{"encode(somephrase)"}, "s4m2phr1s2"},
		{"test2", args{"decode(s4m2phr1s2)"}, "somephrase"},
		{"test3", args{"encode(\"somephrase\")"}, "\"s4m2phr1s2\""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := inputHandler(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("inputHandler() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
