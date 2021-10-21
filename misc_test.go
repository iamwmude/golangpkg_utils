package utils

import (
	"reflect"
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

func TestGetMapValue(t *testing.T) {
	type args struct {
		m   interface{}
		key interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				m:   map[string]string{"k1": "none", "k2": "test"},
				key: "k2",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "map_type_error",
			args: args{
				m:   "test",
				key: "k2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "key_type_error",
			args: args{
				m:   map[string]string{"k1": "none", "k2": "test"},
				key: 666,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not_found",
			args: args{
				m:   map[string]string{"k1": "none", "k2": "test"},
				key: "k3",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMapValue(tt.args.m, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMapValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
