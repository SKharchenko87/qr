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
		canvas *[][]bool
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
		}}, [][]bool{
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
			drawSynchronizationLines(tt.args.canvas)
			if !reflect.DeepEqual(tt.args.canvas, &tt.want) {
				t.Errorf("drawSynchronizationLine() = %v, want %v", tt.args.canvas, tt.want)
			}
		})
	}
}

func Test_generateInfoCanvas(t *testing.T) {
	type args struct {
		version byte
	}
	tests := []struct {
		name string
		args args
		want [][]bool
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
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateInfoCanvas(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateInfoCanvas() = %v, want %v", got, tt.want)
			}
		})
	}
}
