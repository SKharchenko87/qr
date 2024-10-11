package qr

import (
	"reflect"
	"testing"
)

func Test_getCountByteOfBlock(t *testing.T) {
	type args struct {
		level   levelCorrection
		version int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{Medium, 9 - 1}, []int{36, 36, 36, 37, 37}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCountByteOfBlock(tt.args.level, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCountByteOfBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
