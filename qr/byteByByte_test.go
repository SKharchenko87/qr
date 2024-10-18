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

func Test_createByteCorrection(t *testing.T) {
	type args struct {
		level   levelCorrection
		version int
		data    *[]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{High, 2 - 1, &[]byte{64, 196, 132, 84, 196, 196, 242, 194, 4, 132, 20, 37, 34, 16, 236, 17}}, []byte{16, 85, 12, 231, 54, 54, 140, 70, 118, 84, 10, 174, 235, 197, 99, 218, 12, 254, 246, 4, 190, 56, 39, 217, 115, 189, 193, 24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createByteCorrection(tt.args.level, tt.args.version, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createByteCorrection() = %v, want %v", got, tt.want)
			}
		})
	}
}
