package main

import (
	"fmt"
	"testing"
)

func Test_validatePassword(t *testing.T) {
	tests := []struct {
		input int
		want bool
	}{
		{ 111111, true},
		{223450, false},
		{123789, false},
		{997321, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("validate_%d", tt.input), func(t *testing.T) {
			if got := validatePassword(tt.input); got != tt.want {
				t.Errorf("validatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePassword2(t *testing.T) {
	tests := []struct {
		input int
		want bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("validate2_%d", tt.input), func(t *testing.T) {
			if got := validatePassword2(tt.input); got != tt.want {
				t.Errorf("validatePassword2() = %v, want %v", got, tt.want)
			}
		})
	}
}