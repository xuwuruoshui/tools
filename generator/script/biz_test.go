package main

import (
	"reflect"
	"testing"
)

func TestNewBiz(t *testing.T) {
	type args struct {
		fileInputPath string
		fileSuffix string
	}
	tests := []struct {
		name string
		args args
		want *Biz
	}{
		{
			name: "biz ast",
			args: args{fileInputPath: "../model/user.go"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBiz(tt.args.fileInputPath,tt.args.fileSuffix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBiz() = %v, want %v", got, tt.want)
			}
		})
	}
}


