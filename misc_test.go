package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstReturn(t *testing.T) {
	assert.Equal(t, GetFirstReturn(strconv.Atoi("150")), 150)
	assert.Equal(t, GetFirstReturn(nil), nil)
}

func TestGetString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "int",
			args: args{5},
			want: "5",
		},
		{
			name: "float",
			args: args{5.5},
			want: "5.5",
		},
		{
			name: "struct",
			args: args{struct{ Name string }{Name: "test"}},
			want: `{"Name":"test"}`,
		},
		{
			name: "map",
			args: args{map[string]int{"test": 5}},
			want: `{"test":5}`,
		},
		{
			name: "slice",
			args: args{v: []int{5, 5, 5}},
			want: "[5,5,5]",
		},
		{
			name: "array",
			args: args{v: [3]int{5, 5, 5}},
			want: "[5,5,5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.v); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
