package qr

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getCountByteOfBlock(t *testing.T) {
	type args struct {
		level   levelCorrection
		version byte
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
		version byte
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

func Test_mergeBlocks(t *testing.T) {
	type args struct {
		data       [][]byte
		correction [][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Test 1", args{[][]byte{{1, 4}, {2, 5}, {3, 6, 7}}, [][]byte{{10, 13}, {11, 14}, {12, 15, 16}}}, []byte{1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 13, 14, 15, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeBlocks(tt.args.data, tt.args.correction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fill(t *testing.T) {
	type args struct {
		data  []byte
		level levelCorrection
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 byte
	}{
		{"Test 1", args{[]byte("HELLO"), Medium}, []byte{64, 84, 132, 84, 196, 196, 240, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fillBinary(tt.args.data, tt.args.level)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillBinary() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fillBinary() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getKind(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want func(data []byte, level levelCorrection) ([]byte, byte)
	}{
		{"Test 1", args{"0123456789"}, fillNumeric},
		{"Test 2", args{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./"}, fillAlphanumeric},
		{"Test 3", args{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./Привет"}, fillBinary},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKind(tt.args.text); fmt.Sprintf("%p", got) != fmt.Sprintf("%p", tt.want) {
				t.Errorf("getKind() = %p, want %p", got, tt.want)
			}
		})
	}
}

func Test_generateQR(t *testing.T) {
	type args struct {
		text  string
		level levelCorrection
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"Test 1", args{"0123456789", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, O, O, O, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, I, O, I, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, I, O, I, O, O, I, O, I, O, O, O, O, I, O, O, I, O},
			{O, I, I, I, O, O, O, I, I, O, I, O, O, O, I, O, O, O, O, I, O},
			{O, O, I, I, I, O, I, O, O, I, O, O, I, O, O, O, I, I, I, I, O},
			{O, I, I, I, I, O, O, I, I, I, I, O, O, O, I, O, O, O, O, I, O},
			{O, O, O, I, O, I, I, O, O, I, O, O, I, O, I, O, I, O, O, O, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, O, I, O, I, O, O, O, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, I, I, I, O, O, O, O, O},
			{I, O, O, O, O, O, I, O, O, O, O, I, I, I, O, I, I, I, I, I, I},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, I, I, I, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, I, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, I, O, O, O, I, O, O, O, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, O, O, I, I, O, I, I, O, O},
			{I, I, I, I, I, I, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
		}},
		{"Test 2", args{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, I, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, O, I, O, I, O, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, O, I, O, O, O, I, O, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, I, I, I, O, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, O, I, I, I, I, I, O, O, I, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, O, I, O, I, O, O, I, I, I, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, I, I, O, O, O, I, I, O, O, I, I, O, I, I, O, O, O, I, I, I, I, I, O, O},
			{O, I, O, I, O, I, O, O, I, O, O, O, I, O, O, I, I, I, O, I, O, O, O, O, I, I, I, I, I},
			{O, I, I, I, O, O, I, O, O, O, O, I, O, O, O, O, O, I, I, O, O, O, O, O, O, O, I, I, O},
			{O, I, I, I, I, O, O, O, I, I, O, I, I, O, I, O, I, I, O, O, I, I, I, I, O, I, O, I, I},
			{O, I, O, I, O, I, I, I, O, I, I, I, I, O, O, O, I, O, I, O, I, O, O, O, O, O, O, I, I},
			{O, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, O, I, O, I, O, O, I, O, O, O, I, O, I},
			{I, I, I, I, I, O, I, I, O, I, I, I, O, I, I, I, I, I, O, I, I, I, I, O, O, O, O, I, I},
			{I, I, I, I, O, O, O, O, I, I, O, O, O, I, O, I, O, O, I, O, I, I, I, O, I, O, I, I, O},
			{O, O, I, O, O, I, I, I, O, O, I, O, I, O, I, I, O, I, O, O, I, O, I, O, I, O, I, I, O},
			{I, I, O, I, O, O, O, O, O, I, I, O, O, O, O, I, I, O, O, O, O, I, I, I, I, O, I, O, O},
			{I, O, O, I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, O, I, I, O, O, I, I, I, O, I, O},
			{I, O, I, I, O, I, O, I, O, O, I, O, I, O, I, O, I, I, I, O, O, I, I, I, I, I, O, I, O},
			{I, O, O, O, I, I, I, I, I, O, I, O, I, O, O, O, O, I, I, O, I, I, I, I, I, I, I, I, O},
			{O, O, O, O, O, O, O, O, I, O, I, I, O, I, I, I, O, I, I, I, I, O, O, O, I, O, O, I, O},
			{I, I, I, I, I, I, I, O, O, O, O, O, O, I, I, I, O, I, O, O, I, O, I, O, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, I, O, I, O, O, O, I, I, O, O, O, I, O, O, I, I},
			{I, O, I, I, I, O, I, O, I, O, O, I, O, O, I, I, O, I, I, O, I, I, I, I, I, I, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, O, O, O, O, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, O, O, O, I, I, I, O, I, O, O, I, O, I, O, I, O, I, I, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, I, I, I, O, O, I, O, I, O, O, O, I, I, O},
		}},
		{"Test 3", args{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./Привет", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, O, O, I, O, I, I, O, O, O, O, I, O, O, O, I, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, I, I, I, O, I, O, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, I, O, I, I, O, I, I, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, I, I, I, I, I, O, O, O, I, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, I, I, O, O, I, O, I, O, O, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, I, O, I, O, I, I, O, I, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, I, I, I, O, I, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, O, I, O, I, I, O, I, O, I, I, O, O, I, O, O, I, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O},
			{I, O, I, O, O, O, O, O, O, I, O, I, O, I, I, I, I, I, O, I, I, O, O, I, I, O, I, I, O, I, O, O, O},
			{O, I, I, O, O, O, I, I, I, O, O, I, I, I, I, O, O, O, O, I, I, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, I, I, I, O, I, O, O, I, I, O, I, I, I, O, O, O, I, I, O, I, O, I, I, O, O, I, O, O, O, O, O, O},
			{I, I, O, O, I, O, I, I, O, O, I, I, I, O, I, O, I, O, O, I, O, O, O, O, I, O, I, O, O, I, O, I, I},
			{I, O, O, O, I, O, O, I, I, O, O, O, I, O, I, I, O, O, I, I, O, I, O, O, O, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, I, I, O, O, I, I, I, I, O, O, I, I, I, O, I, O, O, O, I, I, I, O, O, I, I, I, O, O},
			{O, O, I, O, I, I, O, I, O, I, I, O, I, O, O, O, I, I, I, I, O, I, I, O, I, O, I, O, O, O, O, I, O},
			{O, I, O, I, I, O, I, O, O, O, O, O, I, O, O, I, I, O, O, I, I, O, I, I, I, I, I, I, I, I, I, I, I},
			{I, I, O, I, O, I, O, O, O, I, O, O, I, O, I, I, I, O, O, I, O, I, O, O, I, O, O, I, O, O, O, O, I},
			{O, O, O, I, I, O, I, I, I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, I, I, O, O, I, I},
			{I, I, O, I, I, I, O, O, I, O, I, O, O, I, I, I, I, O, I, O, O, O, O, I, O, O, I, O, I, I, O, O, I},
			{I, O, I, O, I, O, I, I, O, I, I, I, I, O, I, O, O, I, I, I, I, O, I, I, I, O, O, O, O, O, O, I, O},
			{O, I, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, O, O, I, I, I, O, O, O, I, I, I, O, I, I, O, I},
			{I, I, I, O, I, I, I, O, O, O, I, I, O, O, I, O, I, I, O, O, I, I, I, O, O, I, O, O, O, I, I, I, I},
			{O, I, I, O, O, I, O, I, I, O, I, O, I, I, O, I, O, O, O, I, I, I, O, O, I, I, I, I, O, I, O, I, I},
			{I, O, O, I, O, O, I, I, I, I, I, I, O, O, I, O, O, O, O, I, I, O, O, I, I, I, I, I, I, I, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, I, O, I, I, I, I, O, I, O, O, I, I, I, I, I, O, O, O, I, I, O, I, O},
			{I, I, I, I, I, I, I, O, O, I, O, I, O, I, O, O, I, O, I, O, O, O, I, I, I, O, I, O, I, O, O, I, O},
			{I, O, O, O, O, O, I, O, I, I, O, O, O, O, O, I, O, I, O, I, O, O, O, O, I, O, O, O, I, I, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, I, I, O, O, O, I, O, O, O, I, O, I, I, I, I, I, I, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, I, I, I, O, O, I, I, I, I, O, I, I, O, I, O, I, O, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, I, I, I, I, I, I, O, I, I, I, O, O, O, I, O, O, I, O, O, O, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, I, I, O, O, O, I, I, O, O, I, O, I, I, O, I, I, I, I, O, O, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, I, I, I, I, I, I, I, O, I, O},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateQR(tt.args.text, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateQR() = %v, want %v", got, tt.want)
			}
		})
	}
}
