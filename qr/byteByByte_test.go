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

func Test_drawSearchNodes(t *testing.T) {
	type args struct {
		canvas *[][]bool
		i      byte
		j      byte
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"1", args{&[][]bool{
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
		}, 3, 3}, [][]bool{
			{I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawSearchNode(tt.args.canvas, tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.args.canvas, &tt.want) {
				t.Errorf("drawSearchNode() = %v, want %v", tt.args.canvas, tt.want)
			}
		})
	}
}

func Test_drawAlignmentNode(t *testing.T) {
	type args struct {
		canvas *[][]bool
		i      byte
		j      byte
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"1", args{&[][]bool{
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O},
		}, 3, 3}, [][]bool{
			{O, O, O, O, O, O, O},
			{O, I, I, I, I, I, O},
			{O, I, O, O, O, I, O},
			{O, I, O, I, O, I, O},
			{O, I, O, O, O, I, O},
			{O, I, I, I, I, I, O},
			{O, O, O, O, O, O, O},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawAlignmentNode(tt.args.canvas, tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.args.canvas, &tt.want) {
				t.Errorf("drawAlignmentNode() = %v, want %v", tt.args.canvas, tt.want)
			}
		})
	}
}

func Test_drawSynchronizationLines(t *testing.T) {
	type args struct {
		canvas          *[][]bool
		busyRangeModuls *[]Rectangle
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"1", args{&[][]bool{
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
		}, &[]Rectangle{}}, [][]bool{
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, I, O, I, O, I, O, I, O, I, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawSynchronizationLines(tt.args.canvas, tt.args.busyRangeModuls)
			if !reflect.DeepEqual(tt.args.canvas, &tt.want) {
				t.Errorf("drawSynchronizationLine() = %v, want %v", tt.args.canvas, tt.want)
			}
		})
	}
}

func Test_generateInfoCanvas1(t *testing.T) {
	type args struct {
		version byte
	}
	tests := []struct {
		name  string
		args  args
		want  [][]bool
		want1 []Rectangle
	}{
		//{"1", args{1}, [][]bool{{true, false, true, false}}},
		{"2", args{2}, [][]bool{
			{I, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, I, O, I, O, I, O, O, O, O},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
		}, []Rectangle{
			{iLeftUp: 0, jLeftUp: 0, iRightDown: 8, jRightDown: 8},
			{iLeftUp: 0, jLeftUp: 17, iRightDown: 8, jRightDown: 24},
			{iLeftUp: 17, jLeftUp: 0, iRightDown: 24, jRightDown: 8},
			{iLeftUp: 6, jLeftUp: 0, iRightDown: 6, jRightDown: 24},
			{iLeftUp: 0, jLeftUp: 6, iRightDown: 24, jRightDown: 6},
			{iLeftUp: 16, jLeftUp: 16, iRightDown: 20, jRightDown: 20},
		}},
		{"7", args{7}, [][]bool{
			{I, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O},
			{O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{O, O, O, O, I, O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, I, O, I, O, O, O, O},
			{O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, O, O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{O, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, O, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, I, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, I, O, I, O, O, O, O},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
			{I, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
		}, []Rectangle{
			{iLeftUp: 0, jLeftUp: 0, iRightDown: 8, jRightDown: 8},
			{iLeftUp: 0, jLeftUp: 37, iRightDown: 8, jRightDown: 44},
			{iLeftUp: 37, jLeftUp: 0, iRightDown: 44, jRightDown: 8},
			{iLeftUp: 6, jLeftUp: 0, iRightDown: 6, jRightDown: 44},
			{iLeftUp: 0, jLeftUp: 6, iRightDown: 44, jRightDown: 6},
			{iLeftUp: 4, jLeftUp: 20, iRightDown: 8, jRightDown: 24},
			{iLeftUp: 20, jLeftUp: 4, iRightDown: 24, jRightDown: 8},
			{iLeftUp: 20, jLeftUp: 20, iRightDown: 24, jRightDown: 24},
			{iLeftUp: 20, jLeftUp: 36, iRightDown: 24, jRightDown: 40},
			{iLeftUp: 36, jLeftUp: 20, iRightDown: 40, jRightDown: 24},
			{iLeftUp: 36, jLeftUp: 36, iRightDown: 40, jRightDown: 40},
			{iLeftUp: 0, jLeftUp: 34, iRightDown: 6, jRightDown: 36},
			{iLeftUp: 34, jLeftUp: 0, iRightDown: 36, jRightDown: 6},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateInfoCanvas(tt.args.version)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateInfoCanvas() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("generateInfoCanvas() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
