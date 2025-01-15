package qr

import (
	"fmt"
	"reflect"
	"testing"
)

func generateQR0(text string, level LevelCorrection) [][]bool {
	data, version := fillBinary([]byte(text), level)
	fmt.Println(text)
	fmt.Println(version)
	fmt.Println(data)

	countByteOfBlock := getCountByteOfBlock(level, version-1)
	dataBlock := fillBlocks(data, countByteOfBlock)

	correctionBlock := make([][]byte, len(dataBlock))
	for i, bytes := range dataBlock {
		correctionBlock[i] = createByteCorrection(level, version-1, &bytes)
	}

	blocks := mergeBlocks(dataBlock, correctionBlock)

	canvas, busyRangeModul := generateInfoCanvas(version)
	generatePreCode(blocks, &canvas, &busyRangeModul)

	lengthCanvas := len(canvas)
	candidateCanvas := make([][]bool, lengthCanvas)
	copy(candidateCanvas, canvas)

	drawMask(&candidateCanvas, &busyRangeModul, 8, 1)
	drawCodeMaskLevelCorrection(&candidateCanvas, level, 1)
	printQR(&candidateCanvas)

	return canvas
}

func Test_generateQR0(t *testing.T) {
	type args struct {
		text  string
		level LevelCorrection
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"Test a", args{"a", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, O, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{O, O, I, I, O, O, O, O, I, O, I, I, O, I, I, I, O, I, I, O, O},
			{O, I, O, O, O, O, I, I, I, I, O, I, I, I, O, I, I, I, O, I, I},
			{I, I, O, I, O, I, O, I, I, O, O, I, O, I, I, I, O, I, I, I, O},
			{I, I, O, I, O, O, I, I, I, I, O, I, I, I, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, O, O, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, I, O, O, O, I, I, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, O, I, I, I, O, I, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, I, O, I, I, O, O},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, I, I, I, I, O, I},
		}},
		{"Test a0", args{"a0", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, I, I, O, O, I, O, O, I, O, I},
			{I, I, O, I, I, I, O, O, I, I, I, O, O, O, O, O, O, O, O, I, I},
			{I, I, I, O, O, I, I, O, I, O, I, O, O, O, I, O, O, I, I, I, I},
			{O, I, O, O, O, O, O, O, I, I, O, O, I, O, O, O, I, O, O, I, O},
			{I, I, O, I, O, O, I, I, O, O, I, O, O, O, I, O, O, I, I, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, I, O, I, I, I, O, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, I, O, I, I, I, O, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, I, I, I, O, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, I, I, I, I, I, I, I, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, O, I, O, O, O, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, O, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, O, O, O, I, O, I},
		}},
		{"Test a01", args{"a01", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, O, O, O, O, I, O, O, I, O, I},
			{I, O, O, I, O, I, O, I, O, I, I, I, O, I, I, I, O, O, O, I, I},
			{O, I, O, I, I, O, I, O, I, I, I, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, I, I, O, I, I, O, I, I, O, I, I, I, O, O, O, I, O},
			{O, I, I, O, I, O, I, I, O, I, O, I, I, I, I, I, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, O, O, O, I, O, O, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, O, I, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, I, O, O, I, I, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, I, I, I, O, I, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, I, I, I, O, I, I, O, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, I, I, I, I, I, I, O, I},
		}},
		{"Test a012", args{"a012", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, O, I, I, O, O, I, O, O, I, O, I},
			{O, O, O, I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, I, I},
			{I, O, I, I, O, O, I, I, O, I, I, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, I, O, O, I, O, I, I, O, I, O, O, O, I, O, O, I, O},
			{I, O, I, O, O, O, I, O, I, O, O, O, O, O, I, O, O, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, I, I, I, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, I, I, I, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, I, I, I, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, I, I, I, I, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, O, O, O, I, O, O, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, O, O, I, I, O, I},
		}},
		{"Test a0123", args{"a0123", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, O, O, O, O, I, O, O, I, O, I},
			{I, O, O, O, I, I, O, O, O, O, O, I, O, I, I, I, O, O, O, I, I},
			{O, O, O, I, I, O, I, I, I, I, O, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, I, I, O, I, I, O, O, I, O, I, I, I, O, O, O, I, O},
			{I, I, O, O, O, O, I, I, O, I, I, I, I, I, I, I, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, O, O, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, O, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, I, I, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, I, I, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, I, I, I, I, O, O, O, I},
		}},
		{"Test a01234", args{"a01234", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, I, I, I, O, O, I, O, O, I, O, I},
			{I, O, O, O, I, I, O, I, O, I, O, O, O, O, O, O, O, O, O, I, I},
			{O, I, O, I, I, I, I, O, I, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, I, O, O, I, O, O, I, I, O, O, O, I, O, O, O, I, O, O, I, O},
			{O, I, O, I, O, I, I, I, I, I, I, O, O, O, I, O, O, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, O, I, O, I, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, I, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, I, I, O, I, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, O, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, I, I, O, O, O, I},
		}},
		{"Test a012345", args{"a012345", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{I, O, I, I, O, O, O, O, O, I, I, I, O, I, I, O, O, O, O, I, I},
			{O, I, O, O, I, I, I, I, O, O, O, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, I, O, O, O, O, O, I, I, O, I, I, I, O, O, O, I, O},
			{I, O, O, I, O, I, I, O, I, I, I, I, I, I, I, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, O, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, I, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, O, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, I, I, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, I, I, I, I, I, I, I, O, O, O, I},
		}},
		{"Test a0123456", args{"a0123456", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{I, O, O, O, O, O, O, O, I, I, I, O, O, I, O, I, I, O, O, I, I},
			{I, I, I, O, I, I, I, O, O, O, O, O, O, I, O, I, I, I, I, I, I},
			{O, O, O, I, I, O, O, I, I, I, I, O, I, O, O, I, O, O, O, I, O},
			{I, O, O, O, I, O, I, I, O, O, I, O, O, I, I, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, I, I, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, I, I, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, O, O, O, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, O, O, O, O, I, I, I, O, O, O, I},
		}},
		{"Test a01234567", args{"a01234567", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, O, O, O, O, O, I, O, O, I, O, I},
			{O, I, O, O, I, O, O, I, O, I, I, I, O, I, O, I, I, O, O, I, I},
			{O, O, I, I, I, O, I, I, O, O, I, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, O, I, O, I, O, O, O, I, O, O, O, I, O, O, O, I, O},
			{I, O, I, O, O, I, I, I, I, O, I, I, I, O, O, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, O, O, O, I, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, O, O, O, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, O, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, I, I, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, I, I, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, I, I, I, O, O, O, I},
		}},
		{"Test a012345678", args{"a012345678", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, I, I, O, O, I, O, O, I, O, I},
			{I, O, O, I, I, O, O, I, I, O, I, O, O, I, O, I, I, O, O, I, I},
			{O, O, I, I, I, O, I, I, I, O, O, O, O, I, O, I, I, I, I, I, I},
			{I, I, I, I, O, I, O, I, O, I, I, O, I, O, O, I, O, O, O, I, O},
			{O, O, I, O, I, I, I, I, I, I, I, O, O, O, O, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, I, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, I, I, O, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, O, O, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, O, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, O, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, O, O, O, I, I, I, I, O, O, O, I},
		}},
		{"Test a0123456789", args{"a0123456789", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{I, I, O, O, I, O, O, O, I, O, O, I, O, I, O, I, I, O, O, I, I},
			{I, I, I, I, I, O, I, I, I, O, I, I, I, I, O, I, I, I, I, I, I},
			{O, O, I, O, I, O, O, O, O, O, I, I, O, O, O, I, O, O, O, I, O},
			{O, I, O, I, I, I, I, I, O, O, I, I, I, O, O, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, I, O, O, I, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, O, O, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, O, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, I, I, O, O, O, I, I, O, O, O, I},
		}},
		{"Test a0123456789a", args{"a0123456789a", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{O, I, I, O, I, I, O, O, I, O, O, O, O, I, O, I, I, O, O, I, I},
			{O, O, I, O, O, O, I, I, O, O, I, O, O, I, O, I, I, I, I, I, I},
			{O, I, I, O, O, O, O, O, I, O, O, O, I, O, O, I, O, O, O, I, O},
			{O, O, O, I, I, I, I, I, I, I, I, O, O, O, O, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, I, O, O, I, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, O, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, O, O, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, I, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, O, O, I, O, O, O, I, I, O, O, O, I},
		}},
		{"Test a0123456789a0", args{"a0123456789a0", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, O, O, O, O, O, I, O, O, I, O, I},
			{I, I, I, I, O, I, O, I, I, O, I, O, O, I, O, I, I, O, O, I, I},
			{O, O, I, O, O, I, I, I, I, O, I, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, I, O, O, O, O, I, O, O, O, O, O, O, I, O, O, O, I, O},
			{I, I, O, I, I, I, I, O, O, O, I, I, I, O, O, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, O, I, I, I, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, O, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, O, O, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, O, O, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, I, O, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, O, O, O, I, I, O, O, O, I},
		}},
		{"Test a0123456789a01", args{"a0123456789a01", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, O, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{O, I, O, I, I, I, O, O, I, O, I, I, I, I, O, I, I, O, O, I, I},
			{O, I, I, O, O, I, I, O, O, O, O, I, I, I, O, I, I, I, I, I, I},
			{I, I, I, I, I, I, O, I, O, O, I, O, O, O, O, I, O, O, O, I, O},
			{I, O, I, I, O, O, I, O, I, I, O, I, I, O, O, O, I, O, O, I, I},
			{O, O, O, O, O, O, O, O, I, I, I, I, I, I, O, I, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, O, O, I, I, O, I, O, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, O, O, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, O, O, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, O, O, I, I, O, O, O, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, O, O, O, I, I, O, O, O, I},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateQR0(tt.args.text, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateQR2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillBinary(t *testing.T) {
	type args struct {
		data  []byte
		level LevelCorrection
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 byte
	}{
		{"Test . 'a' ", args{[]byte("a"), Medium}, []byte{64, 22, 16, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'a0' ", args{[]byte("a0"), Medium}, []byte{64, 38, 19, 0, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'a01' ", args{[]byte("a01"), Medium}, []byte{64, 54, 19, 3, 16, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'a012' ", args{[]byte("a012"), Medium}, []byte{64, 70, 19, 3, 19, 32, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'a0123' ", args{[]byte("a0123"), Medium}, []byte{64, 86, 19, 3, 19, 35, 48, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'a01234' ", args{[]byte("a01234"), Medium}, []byte{64, 102, 19, 3, 19, 35, 51, 64, 236, 17, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'a012345' ", args{[]byte("a012345"), Medium}, []byte{64, 118, 19, 3, 19, 35, 51, 67, 80, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'a0123456' ", args{[]byte("a0123456"), Medium}, []byte{64, 134, 19, 3, 19, 35, 51, 67, 83, 96, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'a01234567' ", args{[]byte("a01234567"), Medium}, []byte{64, 150, 19, 3, 19, 35, 51, 67, 83, 99, 112, 236, 17, 236, 17, 236}, 1},
		{"Test . 'a012345678' ", args{[]byte("a012345678"), Medium}, []byte{64, 166, 19, 3, 19, 35, 51, 67, 83, 99, 115, 128, 236, 17, 236, 17}, 1},
		{"Test . 'a0123456789' ", args{[]byte("a0123456789"), Medium}, []byte{64, 182, 19, 3, 19, 35, 51, 67, 83, 99, 115, 131, 144, 236, 17, 236}, 1},
		{"Test . 'a0123456789a' ", args{[]byte("a0123456789a"), Medium}, []byte{64, 198, 19, 3, 19, 35, 51, 67, 83, 99, 115, 131, 150, 16, 236, 17}, 1},
		{"Test . 'a0123456789a0' ", args{[]byte("a0123456789a0"), Medium}, []byte{64, 214, 19, 3, 19, 35, 51, 67, 83, 99, 115, 131, 150, 19, 0, 236}, 1},
		{"Test . 'a0123456789a01' ", args{[]byte("a0123456789a01"), Medium}, []byte{64, 230, 19, 3, 19, 35, 51, 67, 83, 99, 115, 131, 150, 19, 3, 16}, 1},
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
