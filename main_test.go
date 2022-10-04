package main

import (
	"reflect"
	"testing"
)

func Test_coalesce(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want any
	}{{
		"1", []any{7}, 7,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coalesce(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("coalesce() = %v, want %v", got, tt.want)
			}
		})
	}
}
