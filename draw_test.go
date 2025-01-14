package qr

import (
	"reflect"
	"testing"
)

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

func Test_nextPosition(t *testing.T) {
	rectangles1 := []Rectangle{
		{iLeftUp: 0, jLeftUp: 0, iRightDown: 8, jRightDown: 8},
		{iLeftUp: 0, jLeftUp: 13, iRightDown: 8, jRightDown: 20},
		{iLeftUp: 13, jLeftUp: 0, iRightDown: 20, jRightDown: 8},
		{iLeftUp: 6, jLeftUp: 0, iRightDown: 6, jRightDown: 24},
		{iLeftUp: 0, jLeftUp: 6, iRightDown: 24, jRightDown: 6},
	}
	rectangles2 := []Rectangle{
		{iLeftUp: 0, jLeftUp: 0, iRightDown: 8, jRightDown: 8},
		{iLeftUp: 0, jLeftUp: 17, iRightDown: 8, jRightDown: 24},
		{iLeftUp: 17, jLeftUp: 0, iRightDown: 24, jRightDown: 8},
		{iLeftUp: 6, jLeftUp: 0, iRightDown: 6, jRightDown: 24},
		{iLeftUp: 0, jLeftUp: 6, iRightDown: 24, jRightDown: 6},
		{iLeftUp: 16, jLeftUp: 16, iRightDown: 20, jRightDown: 20},
	}
	rectangles7 := []Rectangle{
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
	}
	type args struct {
		busyRangeModuls *[]Rectangle
		i               byte
		j               byte
		lengthCanvas    byte
	}
	tests := []struct {
		name  string
		args  args
		want  byte
		want1 byte
	}{
		{"Test 1. QRVersion 2", args{&rectangles2, 24, 24, 25}, 24, 23},
		{"Test 2. QRVersion 7", args{&rectangles7, 44, 44, 45}, 44, 43},
		{"Test 3. QRVersion 2", args{&rectangles2, 9, 23, 25}, 9, 22},
		{"Test 4. QRVersion 2", args{&rectangles2, 21, 20, 25}, 21, 19},
		{"Test 5. QRVersion 2", args{&rectangles2, 15, 18, 25}, 15, 17},
		{"Test 6. QRVersion 2", args{&rectangles2, 15, 17, 25}, 21, 18},
		{"Test 7. QRVersion 2", args{&rectangles2, 21, 15, 25}, 20, 15},
		{"Test 8. QRVersion 2", args{&rectangles2, 7, 15, 25}, 5, 16},
		{"Test 9. QRVersion 2", args{&rectangles2, 24, 9, 25}, 16, 8},
		{"Test 10. QRVersion 2", args{&rectangles2, 9, 7, 25}, 9, 5},
		{"Test 11. QRVersion 2", args{&rectangles2, 16, 5, 25}, 16, 4},
		{"Test 12. QRVersion 2", args{&rectangles2, 16, 4, 25}, 16, 3},
		{"Test 13. QRVersion 7", args{&rectangles7, 7, 35, 45}, 0, 33},
		{"Test 14. QRVersion 7", args{&rectangles7, 9, 7, 45}, 9, 5},
		{"Test 15. QRVersion 7", args{&rectangles7, 33, 4, 45}, 33, 3},

		{"Test 3.1 QRVersion 2", args{&rectangles2, 9, 22, 25}, 9, 21},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 9, 21, 25}, 10, 22},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 10, 22, 25}, 10, 21},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 10, 21, 25}, 11, 22},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 11, 22, 25}, 11, 21},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 11, 21, 25}, 12, 22},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 12, 22, 25}, 12, 21},
		{"Test 3.1 QRVersion 2", args{&rectangles2, 12, 21, 25}, 13, 22},

		{"Test 16. QRVersion 1", args{&rectangles1, 9, 2, 21}, 9, 1},
		{"Test 16. QRVersion 1", args{&rectangles1, 9, 1, 21}, 9, 0},
		{"Test 16. QRVersion 1", args{&rectangles1, 9, 0, 21}, 10, 1},
		{"Test 16. QRVersion 1", args{&rectangles1, 10, 1, 21}, 10, 0},
		{"Test 16. QRVersion 1", args{&rectangles1, 10, 0, 21}, 11, 1},
		{"Test 16. QRVersion 1", args{&rectangles1, 11, 1, 21}, 11, 0},
		{"Test 16. QRVersion 1", args{&rectangles1, 12, 1, 21}, 12, 0},
		{"Test 16. QRVersion 1", args{&rectangles1, 12, 0, 21}, 20, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := nextPosition(tt.args.busyRangeModuls, tt.args.i, tt.args.j, tt.args.lengthCanvas)
			if got != tt.want {
				t.Errorf("nextPosition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("nextPosition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_drawMask(t *testing.T) {
	type args struct {
		canvas          *[][]bool
		busyRangeModuls *[]Rectangle
		oldMask         byte
		newMask         byte
	}
	tests := []struct {
		name string
		args args
		want *[][]bool
	}{
		{
			"Тест 0. Текст в формате байт `Привет` маска 0", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 0},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, O, O, O, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, O, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, I, I, I, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
				{O, I, I, I, I, I, O, O, O, O, I, I, O, I, O, O, I, O, I, O, O},
				{I, O, I, O, I, O, I, O, I, I, I, I, O, O, O, O, O, I, O, O, O},
				{O, O, I, O, I, I, O, O, O, O, O, I, I, O, I, I, O, O, I, I, O},
				{I, I, I, I, O, I, I, O, O, I, O, I, O, I, I, O, I, I, I, O, I},
				{O, O, O, O, O, O, O, O, O, I, O, I, O, I, O, O, O, I, O, O, O},
				{I, I, I, I, I, I, I, O, O, O, O, O, I, O, I, I, O, I, I, I, O},
				{I, O, O, O, O, O, I, O, O, O, O, I, I, I, I, I, O, I, I, I, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, O, O, I, I, O},
				{I, O, I, I, I, O, I, O, O, I, O, I, O, O, O, O, I, I, O, I, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, O, O, I, O, I},
				{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, I, O, O, I, I, O},
				{I, I, I, I, I, I, I, O, O, I, O, O, I, I, I, O, I, I, I, I, I},
			}},
		{
			"Тест 1. Текст в формате байт `Привет` маска 1", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 1},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, I, O, I, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, I, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, O, I, I, I, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{I, I, I, I, I, I, I, I, I, O, I, O, O, I, O, I, O, O, O, I, O},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{I, O, I, O, O, O, I, I, O, O, O, O, O, O, I, I, I, O, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, I, O, I, I, I, I, O, O, O, I, O, O},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, I, I, I, I, O, O, I, O, I, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, I, O, I, O, O, I, I, O, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, O, O, I, I, O, I, I, I, O, I, O, I},
			}},
		{
			"Тест 2. Текст в формате байт `Привет` маска 2", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 2},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, I, I, O, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, I, I, I, I, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, I, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, I, I, O, I, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, I, O, I, I, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, O, O, O, O, O},
				{I, O, I, I, I, O, O, I, O, O, I, O, I, O, O, O, I, I, O, I, O},
				{I, O, O, I, O, O, I, O, O, O, O, I, O, O, I, I, I, I, O, O, I},
				{I, I, I, O, I, O, O, I, O, O, O, O, O, I, I, I, O, I, O, O, O},
				{I, I, O, O, I, I, I, O, I, O, I, I, O, I, O, I, O, I, I, O, O},
				{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I, I, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, I, O, O, O, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, I, O, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, I, O, O, I, I, I, I, I, O, I, I, I},
				{I, O, I, I, I, O, I, O, O, I, O, O, I, I, O, O, I, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, I, I, O, O, I, O, I, I, O, I, O, O},
				{I, O, O, O, O, O, I, O, O, I, I, I, I, I, I, I, O, I, O, O, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, I, I, O, I, O, I, I, I, O},
			}},
		{
			"Тест 3. Текст в формате байт `Привет` маска 3", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 3},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, I, I, O, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, I, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, I, O, O, O, O, O, O, O, O, O, O},
				{I, O, I, I, I, O, O, I, O, O, I, O, I, O, O, O, I, I, O, I, O},
				{O, O, I, O, O, I, I, O, I, I, O, O, I, O, O, O, I, O, I, O, O},
				{O, O, I, I, O, O, O, O, O, I, I, O, I, O, I, O, I, I, I, I, O},
				{I, I, O, O, I, I, I, O, I, O, I, I, O, I, O, I, O, I, I, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, I, I, O, I, O, I, I},
				{I, I, I, I, I, I, I, O, O, O, O, O, O, I, O, I, O, I, O, O, I},
				{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, I, O, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, I, O, I, O, O, I, I, O, I, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, O, O, O, I, O, O, O, I, O},
				{I, O, I, I, I, O, I, O, O, I, I, O, O, I, O, I, I, O, I, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, I, O, O, O, O, I, O, I},
				{I, I, I, I, I, I, I, O, O, I, O, O, O, O, O, O, I, I, O, O, O},
			}},
		{
			"Тест 4. Текст в формате байт `Привет` маска 4", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 4},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, I, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
				{I, I, O, O, I, O, O, O, I, I, I, O, I, I, I, I, I, I, O, O, I},
				{O, O, O, I, I, I, I, O, O, O, I, O, I, O, I, I, O, O, I, O, I},
				{O, I, I, O, O, I, O, I, O, O, I, I, I, I, I, I, I, O, I, O, O},
				{I, O, I, I, I, I, I, I, O, I, I, I, O, O, I, O, O, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, O, O, I, O, I},
				{I, I, I, I, I, I, I, O, O, I, O, I, O, O, O, O, O, O, O, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, I, I, I, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, I, O, O, O, I, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, I, I, O, I, I, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, I, O, I, O, I, O, O, O},
				{I, O, O, O, O, O, I, O, O, I, O, O, O, I, I, I, I, O, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, I, O, I, O, O, I, I, O, I},
			}},
		{
			"Тест 5. Текст в формате байт `Привет` маска 5", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 5},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, I, O, I, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, I, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, O, O, O, O, O},
				{I, O, O, O, O, O, O, I, I, I, O, O, I, O, I, I, O, I, O, I, I},
				{I, O, O, I, O, O, I, O, O, O, O, I, O, O, I, I, I, I, O, O, I},
				{I, I, I, I, I, O, O, I, O, I, O, O, O, I, I, O, O, I, O, O, O},
				{I, O, I, O, O, O, I, I, O, O, O, O, O, O, I, I, I, O, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, I, I, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, I, O, O, O, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, I, I, O, O, O, O, O, I, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, O, I, I, I, I, I, O, I, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, I, I, O, I, I, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, I, O, I, O, O, I, I, O, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, I, I, I, O, O, I, O, O, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, I, I, O, I, O, I, I, I, O},
			}},
		{
			"Тест 6. Текст в формате байт `Привет` маска 6", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 6},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, I, O, I, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, O, I, I, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, O, I, O, I, O, O, O, O, O, O, O, O},
				{I, O, O, O, O, O, O, I, I, I, O, O, I, O, I, I, O, I, O, I, I},
				{I, O, I, I, O, I, I, O, I, O, O, O, O, O, O, I, I, O, O, O, O},
				{I, I, I, I, O, I, O, I, O, I, I, I, O, I, I, O, I, O, O, O, O},
				{I, O, I, O, O, O, I, I, O, O, O, O, O, O, I, I, I, O, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, I, O, O, I, O, I},
				{I, I, I, I, I, I, I, O, O, I, O, O, I, I, O, O, O, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, I, I, O, O, O, O, O, I, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, I, O, I, I, I, I, I, O},
				{I, O, I, I, I, O, I, O, O, O, I, I, I, I, O, I, O, I, I, O, O},
				{I, O, I, I, I, O, I, O, O, I, O, I, O, O, I, I, O, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, I, O, O, O, O, I, O, I, I},
				{I, I, I, I, I, I, I, O, O, O, O, O, I, O, O, I, I, I, I, O, O},
			}},
		{
			"Тест 7. Текст в формате байт `Привет` маска 7", args{canvas: &[][]bool{
				{I, I, I, I, I, I, I, O, O, O, I, O, I, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
				{O, O, I, O, I, O, O, I, O, I, I, O, O, O, O, I, I, I, I, I, O},
				{O, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
				{O, I, I, I, I, O, O, I, O, I, O, O, I, I, I, O, O, I, I, O, O},
				{O, I, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O},
				{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I, I, O, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
				{I, O, I, I, I, O, I, O, O, O, I, O, I, I, O, O, I, O, O, O, O},
				{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, I, O, O, I, O, O, O, I, O, I, O},
			}, busyRangeModuls: &[]Rectangle{
				{0, 0, 8, 8},
				{0, 13, 8, 20},
				{13, 0, 20, 8},
				{6, 0, 6, 20},
				{0, 6, 20, 6},
			}, oldMask: 8, newMask: 7},
			&[][]bool{
				{I, I, I, I, I, I, I, O, O, O, O, O, O, O, I, I, I, I, I, I, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, O, O, I, O, O, O, O, O, I},
				{I, O, I, I, I, O, I, O, O, I, O, I, O, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, O, I, I, I, O, I},
				{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
				{I, O, O, O, O, O, I, O, O, I, I, O, O, O, I, O, O, O, O, O, I},
				{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
				{O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O, O},
				{O, O, O, O, O, O, I, O, O, I, I, I, I, O, O, O, O, O, O, O, O},
				{O, I, I, I, I, I, O, O, O, O, I, I, O, I, O, O, I, O, I, O, O},
				{I, I, I, O, O, O, I, I, I, I, O, I, O, I, O, O, I, I, O, I, O},
				{O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, I, O, I, I, I, I},
				{I, I, I, I, O, I, I, O, O, I, O, I, O, I, I, O, I, I, I, O, I},
				{O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, I, I, O, I, O},
				{I, I, I, I, I, I, I, O, O, O, O, I, I, O, O, I, O, O, I, I, I},
				{I, O, O, O, O, O, I, O, O, O, O, I, I, I, I, I, O, I, I, I, O},
				{I, O, I, I, I, O, I, O, O, O, O, O, I, O, O, O, I, O, I, O, O},
				{I, O, I, I, I, O, I, O, O, I, O, O, O, O, I, O, I, O, O, I, I},
				{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, O, O, O, I, O, I},
				{I, O, O, O, O, O, I, O, O, I, O, O, O, I, I, I, I, O, I, O, O},
				{I, I, I, I, I, I, I, O, O, I, O, I, I, I, O, O, I, O, I, I, O},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawMask(tt.args.canvas, tt.args.busyRangeModuls, tt.args.oldMask, tt.args.newMask)
			if !reflect.DeepEqual(tt.args.canvas, tt.want) {
				t.Errorf("generateQR0() = %v, want %v", tt.args.canvas, tt.want)
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
