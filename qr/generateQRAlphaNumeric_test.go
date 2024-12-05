package qr

import (
	"fmt"
	"reflect"
	"testing"
)

func generateQR2(text string, level levelCorrection) [][]bool {
	data, version := fillAlphanumeric([]byte(text), level)
	fmt.Println(version)
	fmt.Println(data)

	countByteOfBlock := getCountByteOfBlock(level, version-1)
	dataBlock := fillBlocks(data, countByteOfBlock)

	correctionBlock := make([][]byte, len(dataBlock))
	for i, bytes := range dataBlock {
		correctionBlock[i] = createByteCorrection(level, version-1, &bytes)
	}

	blocks := mergeBlocks(dataBlock, correctionBlock)

	canvas, busyRangeModuls := generateInfoCanvas(version)
	generatePreCode(blocks, &canvas, &busyRangeModuls)

	lengthCanvas := len(canvas)
	candidateCanvas := make([][]bool, lengthCanvas)
	copy(candidateCanvas, canvas)
	//oldMask := byte(8)
	//printQR(&candidateCanvas)
	drawMask(&candidateCanvas, &busyRangeModuls, 8, 1)
	drawCodeMaskLevelCorrection(&candidateCanvas, level, 1)
	printQR(&candidateCanvas)

	//for i := 0; i < 8; i++ {
	//	drawMask(&candidateCanvas, &busyRangeModuls, oldMask, byte(i))
	//	drawCodeMaskLevelCorrection(&candidateCanvas, level, byte(i))
	//	printQR(&candidateCanvas)
	//	//drawCodeMaskLevelCorrection(&candidateCanvas, level, byte(i))
	//	oldMask = byte(i)
	//}

	//drawMask(&candidateCanvas, &busyRangeModuls, 8, 2)
	//drawCodeMaskLevelCorrection(&candidateCanvas, level, 2)

	// ToDo применить маску
	return canvas
}

func Test_generateQR2(t *testing.T) {
	type args struct {
		text  string
		level levelCorrection
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"Test A", args{"A", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, O, O, O, O, O, I, O, O, I, O, I},
			{O, I, O, O, I, O, O, O, O, O, I, I, O, I, I, I, O, I, I, O, O},
			{O, I, O, O, O, O, I, O, O, O, I, I, I, I, O, I, I, I, O, I, I},
			{O, O, O, I, I, I, O, I, O, O, I, I, O, I, I, I, O, I, I, O, O},
			{I, O, O, I, O, O, I, O, I, I, O, I, I, I, I, I, I, I, I, O, I},
			{O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, O, O, O, I, I, O},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, I, O, O, O, I, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, I, I, I, O, I, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, I, I, I, O, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, I, I, I, I, I, I},
		}},
		{"Test A0", args{"A0", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{O, O, I, I, I, I, O, I, I, I, I, O, O, O, O, O, O, O, O, O, I},
			{O, O, I, O, O, I, I, O, I, O, O, O, O, O, I, O, O, I, I, I, I},
			{O, I, I, I, O, I, O, I, O, O, O, O, I, O, O, O, I, O, O, O, O},
			{I, I, I, I, O, O, I, I, O, I, I, O, O, O, I, O, O, I, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, I, I, O, I, I, I, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, I, I, I, O, I, I, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, I, I, I, I, I, I, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, O, O, O, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, O, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, O, O, O, I, I, I},
		}},
		{"Test A01", args{"A01", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{I, I, I, O, I, I, O, I, I, O, I, I, O, I, I, I, O, O, O, O, I},
			{O, O, I, I, O, O, I, I, I, O, I, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, O, I, O, I, I, O, O, I, O, I, I, I, O, I, O, O, O},
			{I, I, O, O, I, I, I, O, I, O, O, I, I, I, I, I, I, I, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, O, O, I, I, I, O},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, O, O, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, I, O, O, I, I, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, I, I, I, O, I, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, I, I, I, O, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, I, I, I, I, I, I, I},
		}},
		{"Test A012", args{"A012", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{O, I, O, O, I, O, O, O, I, I, I, I, O, I, I, I, O, O, O, O, I},
			{I, O, O, O, O, I, I, O, O, I, O, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, I, O, I, O, I, O, I, I, I, O, I, I, I, O, I, O, O, O},
			{I, I, I, I, O, I, I, I, O, O, I, I, I, I, I, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, O, O, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, O, I, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, I, O, O, I, I, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, I, I, I, O, I, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, I, I, I, O, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, I, I, I, I, I, I},
		}},
		{"Test A0123", args{"A0123", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, I, I, O, O, I, O, O, I, O, I},
			{I, O, I, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, I},
			{O, I, I, O, I, O, I, I, O, I, I, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, I, O, O, O, I, I, O, I, O, O, O, I, I, O, O, O},
			{I, I, O, I, I, O, I, O, O, I, O, O, O, O, I, O, O, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, O, I, O, I, I, I, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, O, I, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, I, I, I, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, O, O, O, O, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, O, O, O, I, O, O, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, O, O, I, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, O, O, I, I, I, I},
		}},
		{"Test A01234", args{"A01234", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, I, O, O, O, O, I, O, O, I, O, I},
			{O, I, O, O, I, O, O, I, I, O, I, I, O, I, I, I, O, O, O, O, I},
			{I, I, I, O, I, I, I, I, I, O, I, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, I, I, O, I, I, O, I, I, O, I, I, I, O, I, O, O, O},
			{O, I, I, I, O, I, I, O, I, O, I, I, I, I, I, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, O, I, O, O, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, O, I, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, I, O, O, I, O, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, I, I, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, O, O, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, I, I, I, I, I, I, I},
		}},
		{"Test A012345", args{"A012345", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{O, O, O, O, O, I, O, I, O, O, I, I, O, I, I, I, O, O, O, O, I},
			{I, I, O, I, I, I, I, I, I, O, O, I, I, I, O, I, I, I, I, I, I},
			{I, O, O, O, O, O, O, I, I, I, I, I, O, I, I, I, O, I, O, O, O},
			{I, I, I, O, I, O, I, O, O, O, O, I, I, I, I, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, O, O, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, O, I, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, O, O, I, O, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, I, I, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, I, O, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, O, I, I, I, I, I},
		}},
		{"Test A0123456", args{"A0123456", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{I, I, O, I, O, O, O, O, I, I, O, O, O, O, O, O, O, O, O, O, I},
			{I, O, O, O, I, O, I, O, I, O, O, O, O, O, I, O, O, I, I, I, I},
			{I, I, I, O, I, O, O, I, O, I, O, O, I, O, O, O, I, I, O, O, O},
			{O, O, I, I, I, I, I, O, O, I, I, O, O, O, I, O, O, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, O, I, O, I, I, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, I, I, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, I, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, O, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, O, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, O, I, I, I, I, I},
		}},
		{"Test A01234567", args{"A01234567", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{I, O, O, I, I, O, O, O, O, I, I, I, O, I, I, O, O, O, O, O, I},
			{I, I, I, O, O, I, I, I, O, O, O, I, I, I, O, I, I, I, I, I, I},
			{O, O, O, I, I, O, O, O, I, O, O, I, O, I, I, O, O, I, O, O, O},
			{I, I, O, O, O, I, I, I, I, O, I, I, I, I, I, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, I, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, O, O, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, O, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, O, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, I, I, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, O, I, I, I, I, I},
		}},
		{"Test A012345678", args{"A012345678", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, O, O, O, O, I, O, O, I, O, I},
			{I, O, O, I, I, O, O, I, I, O, I, I, O, I, I, O, O, O, O, O, I},
			{I, O, O, I, O, I, I, I, I, O, O, I, I, I, O, I, I, I, I, I, I},
			{O, I, O, O, I, O, O, I, I, I, I, I, O, I, I, I, I, I, O, O, O},
			{O, I, O, I, I, I, I, I, O, O, O, I, I, I, I, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, O, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, O, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, I, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, O, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, O, I, I, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, O, I, I, I, I, I},
		}},
		{"Test A0123456789", args{"A0123456789", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{O, I, O, O, I, O, O, I, I, I, I, O, O, I, O, O, I, O, O, O, I},
			{I, O, O, O, I, O, I, O, I, I, I, O, O, I, I, I, I, I, I, I, I},
			{I, O, I, O, O, O, O, I, O, O, I, O, I, O, O, I, I, I, O, O, O},
			{I, I, O, O, I, O, I, I, O, O, I, O, O, I, I, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, I, I, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, I, I, O, O, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, O, I, I, O, I, I, I, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, I, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, O, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, O, O, O, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, O, I, I, I, I, I},
		}},
		{"Test A0123456789A", args{"A0123456789A", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, O, O, I, O, O, I, O, I},
			{O, I, O, O, I, O, O, I, O, O, I, I, O, O, O, I, I, O, O, O, I},
			{O, I, O, I, O, O, I, I, I, O, O, I, I, O, O, I, I, I, I, I, I},
			{I, O, O, I, O, O, O, I, I, O, I, I, O, I, I, I, I, I, O, O, O},
			{O, O, I, I, O, O, I, O, O, O, I, I, I, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, I, I, O, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, I, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, I, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, I, I, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, I, I, I, O, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, I, I, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, I, O, I, I, I, I, I},
		}},
		{"Test A0123456789A0", args{"A0123456789A0", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, O, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, I, I, O, O, I, O, O, I, O, I},
			{I, I, I, O, I, I, O, I, I, O, I, O, O, O, O, I, I, O, O, O, I},
			{O, O, O, I, I, O, I, I, O, I, O, O, O, O, O, I, I, I, I, I, I},
			{I, O, I, O, O, O, O, I, I, I, I, O, I, I, I, I, I, I, O, O, O},
			{I, O, O, I, I, I, I, O, I, I, I, O, O, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, I, I, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, I, O, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, O, O, I, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, I, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, O, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, O, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, I, I, O, I, I, I, I, I},
		}},
		{"Test A0123456789A01", args{"A0123456789A01", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, I, I, O, O, I, O, O, I, O, I},
			{I, O, O, O, O, I, O, I, O, I, I, O, O, O, O, I, I, O, O, O, I},
			{I, I, O, I, I, I, I, I, O, O, O, O, O, O, O, I, I, I, I, I, I},
			{I, I, O, I, I, O, O, O, O, I, I, O, I, I, I, I, I, I, O, O, O},
			{O, O, O, O, O, I, I, I, I, O, I, O, O, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, I, I, I, I, O, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, O, O, I, O, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, I, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, O, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, I, I, O, I, I, I, I, I},
		}},
		{"Test A0123456789A012", args{"A0123456789A012", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, O, O, O, O, I, O, O, I, O, I},
			{I, I, O, O, O, I, O, I, O, O, O, I, O, O, O, I, I, O, O, O, I},
			{I, I, I, I, O, I, I, O, O, I, I, I, I, O, O, I, I, I, I, I, I},
			{I, O, O, I, O, O, O, O, I, O, I, I, O, I, I, I, I, I, O, O, O},
			{I, O, I, O, O, O, I, O, O, O, I, I, I, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, I, I, O, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, I, O, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, I, I, I, O, I, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, I, I, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, O, O, I, I, I, I, I},
		}},
		{"Test A0123456789A0123", args{"A0123456789A0123", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, I, I, O, O, I, O, O, I, O, I},
			{O, O, O, I, O, O, O, O, O, I, I, O, O, O, O, I, I, O, O, O, I},
			{I, I, I, O, I, I, I, I, O, O, I, O, O, O, O, I, I, I, I, I, I},
			{O, O, I, I, I, O, O, O, O, I, O, O, I, I, I, I, I, I, O, O, O},
			{I, I, O, I, O, I, I, O, I, O, O, O, O, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, I, O, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, O, O, I, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, I, I, I, O, I, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, O, I, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, I, O, O, I, I, I, I, I},
		}},
		{"Test A0123456789A01234", args{"A0123456789A01234", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, I, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, I, I, I, O, O, I, O, O, I, O, I},
			{I, I, I, O, I, O, O, O, I, I, O, O, O, O, O, I, I, O, O, O, I},
			{O, I, I, O, I, I, I, O, O, O, I, O, O, O, O, I, I, I, I, I, I},
			{O, I, I, I, O, I, O, I, O, I, I, O, I, I, I, I, I, I, O, O, O},
			{I, O, I, O, I, O, I, O, O, O, O, O, O, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, I, O, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, I, I, I, I, O, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I, O, O, O, O, O},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, I, I, I, O, I, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, O, O, O, I, O, O, I, I, I, I, I},
		}},
		{"Test A0123456789A012345", args{"A0123456789A012345", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, O, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, I, O, O, O, O, I, O, O, I, O, I},
			{I, I, O, O, O, I, O, I, I, I, O, O, O, O, O, I, I, O, O, O, I},
			{O, I, I, O, I, I, I, I, I, O, O, I, I, O, O, I, I, I, I, I, I},
			{I, O, O, O, O, I, O, I, I, O, I, O, O, I, I, I, I, I, O, O, O},
			{O, O, I, I, I, I, I, O, O, O, I, I, I, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, I, O, O, I, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, I, I, I, O, I, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, I, I, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, I, O, O, I, I, I, I, I},
		}},
		{"Test A0123456789A0123456", args{"A0123456789A0123456", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, I, I, I, O, O, I, O, O, I, O, I},
			{O, O, I, I, O, I, O, I, I, O, O, O, O, O, O, I, I, O, O, O, I},
			{I, O, O, I, O, O, I, I, I, I, O, I, O, O, O, I, I, I, I, I, I},
			{O, O, I, O, O, I, O, O, O, O, O, I, O, I, I, I, I, I, O, O, O},
			{O, I, O, O, I, I, I, I, I, O, O, I, I, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, O, I, I, O, I, I, O, I, O},
			{I, O, O, O, O, O, I, O, O, I, I, I, I, O, O, I, O, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, O, I, O, I, I, I, O, I, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, I, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, I, I, O, O, I, O, O, I, I, I, I, I},
		}},
		{"Test A0123456789A01234567", args{"A0123456789A01234567", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, O, O, O, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, I, I, O, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, I, O, I, O, O, O, I, O, O, I, O, I},
			{I, O, O, O, I, I, O, O, O, I, O, O, I, O, O, I, I, O, O, O, I},
			{I, O, I, O, I, I, I, O, I, O, O, I, I, O, O, I, I, I, I, I, I},
			{I, O, I, I, O, I, O, I, I, O, O, O, I, I, I, I, I, I, O, O, O},
			{O, O, O, O, I, O, I, I, O, I, O, I, I, I, O, I, I, O, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, I, I, I, O},
			{I, I, I, I, I, I, I, O, I, I, O, I, O, I, I, O, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, I, I, O, O, I, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, I, I, I, O, I, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, I, O, O, I, O, I, I, O, O, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, I, I, I, I, I, O, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, I, O, O, I, I, I, I, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, O, O, I, O, O, I, I, I, I, I},
		}},
		{"Test A0123456789A012345678", args{"A0123456789A012345678", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, I, O, I, O, O, I, I, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, I, O, O, I, I, O, I, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, O, I, I, I, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, O, I, I, O, O, I, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, O, I, I, I, I, O, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, O, I, O, O, O, I, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, O, O, O, O, I, I, I, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, I, I, I, I, I, O, I, O, O, I, O, O, I, O, I},
			{O, O, O, O, O, O, O, O, O, O, I, O, I, I, I, O, O, O, O, O, O, I, I, I, O},
			{I, O, I, O, I, O, I, I, I, I, O, I, O, O, O, I, I, O, I, I, I, I, O, O, I},
			{I, O, I, O, O, I, O, O, I, I, I, I, I, I, O, O, O, O, O, O, O, O, O, O, O},
			{O, O, I, I, O, O, I, O, O, O, I, O, I, I, O, I, I, I, I, I, O, I, O, I, I},
			{O, I, I, I, O, O, O, O, I, I, I, O, I, O, O, O, I, O, I, I, I, O, O, O, I},
			{I, I, O, I, I, O, I, I, O, O, O, O, I, I, I, I, I, O, O, O, O, I, O, I, I},
			{O, O, O, I, I, I, O, O, O, O, O, I, I, O, I, I, O, O, I, O, O, I, I, O, O},
			{I, I, O, I, O, I, I, O, I, I, I, O, O, I, I, O, I, I, I, I, I, I, I, O, O},
			{O, O, O, O, O, O, O, O, I, I, O, O, O, I, I, O, I, O, O, O, I, I, O, I, O},
			{I, I, I, I, I, I, I, O, I, I, I, I, O, O, O, O, I, O, I, O, I, O, O, I, O},
			{I, O, O, O, O, O, I, O, O, O, O, O, O, I, O, O, I, O, O, O, I, O, O, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, I, I, O, I, I, I, I, I, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, O, O, O, I, O, O, O, I, O, I, I, I, O, I, O, O},
			{I, O, I, I, I, O, I, O, I, O, O, O, I, I, I, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, O, I, I, O, I, I, O, O, I, I, I, I, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, O, I, I, I, O, O, O, I, I, I, I, I, I},
		}},
		{"Test A0123456789A0123456789", args{"A0123456789A0123456789", Medium}, [][]bool{
			{I, I, I, I, I, I, I, O, I, I, O, I, I, I, I, I, O, O, I, I, I, I, I, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, O, O, O, O, I, O, O, I, O, O, O, O, O, I},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, I, O, O, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, I, I, O, O, O, I, O, O, I, O, I, I, I, O, I},
			{I, O, I, I, I, O, I, O, O, I, O, I, O, O, O, I, I, O, I, O, I, I, I, O, I},
			{I, O, O, O, O, O, I, O, I, O, I, I, O, I, I, O, O, O, I, O, O, O, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, O, I, O, I, O, I, I, I, I, I, I, I},
			{O, O, O, O, O, O, O, O, O, O, I, O, O, I, O, O, O, O, O, O, O, O, O, O, O},
			{I, O, I, O, O, O, I, I, O, O, O, O, O, I, O, I, I, O, O, I, O, O, I, O, I},
			{I, I, O, O, O, I, O, I, I, O, I, O, O, O, O, O, I, O, O, O, O, I, I, I, O},
			{I, I, O, I, O, I, I, I, O, I, I, I, I, I, I, I, I, O, I, I, I, I, O, O, I},
			{O, O, O, I, I, I, O, I, O, O, I, O, O, O, I, O, I, O, O, O, O, O, O, O, O},
			{O, O, O, I, O, I, I, O, O, O, O, I, O, I, I, O, O, I, I, I, O, I, O, I, I},
			{O, O, I, O, O, I, O, O, O, I, I, O, O, I, I, O, I, O, I, I, I, O, O, O, I},
			{I, I, O, I, I, O, I, I, O, O, O, I, O, O, O, I, I, O, O, O, O, I, O, I, I},
			{O, O, I, O, O, O, O, I, I, O, I, O, O, I, O, I, O, O, I, O, O, I, I, O, O},
			{I, I, I, I, O, I, I, I, O, I, O, I, I, I, O, O, I, I, I, I, I, I, I, O, O},
			{O, O, O, O, O, O, O, O, I, O, O, O, I, O, O, O, I, O, O, O, I, I, O, I, O},
			{I, I, I, I, I, I, I, O, I, O, I, O, I, I, I, O, I, O, I, O, I, O, O, I, I},
			{I, O, O, O, O, O, I, O, O, O, I, I, I, O, I, O, I, O, O, O, I, O, O, I, I},
			{I, O, I, I, I, O, I, O, O, I, O, O, O, I, I, I, I, I, I, I, I, O, O, I, O},
			{I, O, I, I, I, O, I, O, O, I, I, O, O, I, I, O, I, O, I, I, I, O, I, O, O},
			{I, O, I, I, I, O, I, O, I, I, I, I, O, O, O, I, I, I, I, I, I, O, I, I, I},
			{I, O, O, O, O, O, I, O, O, I, O, O, O, I, O, I, O, O, I, I, I, I, O, O, I},
			{I, I, I, I, I, I, I, O, I, O, O, I, I, I, O, I, O, O, O, I, I, I, I, I, I},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateQR2(tt.args.text, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateQR2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillAlphanumeric(t *testing.T) {
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
		{"Test . 'A' ", args{[]byte("A"), Medium}, []byte{32, 9, 64, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A0' ", args{[]byte("A0"), Medium}, []byte{32, 17, 194, 0, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'A01' ", args{[]byte("A01"), Medium}, []byte{32, 25, 194, 4, 0, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A012' ", args{[]byte("A012"), Medium}, []byte{32, 33, 194, 5, 224, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A0123' ", args{[]byte("A0123"), Medium}, []byte{32, 41, 194, 5, 225, 128, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'A01234' ", args{[]byte("A01234"), Medium}, []byte{32, 49, 194, 5, 226, 44, 0, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A012345' ", args{[]byte("A012345"), Medium}, []byte{32, 57, 194, 5, 226, 44, 80, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A0123456' ", args{[]byte("A0123456"), Medium}, []byte{32, 65, 194, 5, 226, 44, 115, 128, 236, 17, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'A01234567' ", args{[]byte("A01234567"), Medium}, []byte{32, 73, 194, 5, 226, 44, 115, 142, 0, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A012345678' ", args{[]byte("A012345678"), Medium}, []byte{32, 81, 194, 5, 226, 44, 115, 148, 48, 236, 17, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A0123456789' ", args{[]byte("A0123456789"), Medium}, []byte{32, 89, 194, 5, 226, 44, 115, 148, 50, 64, 236, 17, 236, 17, 236, 17}, 1},
		{"Test . 'A0123456789A' ", args{[]byte("A0123456789A"), Medium}, []byte{32, 97, 194, 5, 226, 44, 115, 148, 51, 62, 0, 236, 17, 236, 17, 236}, 1},
		{"Test . 'A0123456789A0' ", args{[]byte("A0123456789A0"), Medium}, []byte{32, 105, 194, 5, 226, 44, 115, 148, 51, 62, 0, 0, 236, 17, 236, 17}, 1},
		{"Test . 'A0123456789A01' ", args{[]byte("A0123456789A01"), Medium}, []byte{32, 113, 194, 5, 226, 44, 115, 148, 51, 62, 0, 64, 236, 17, 236, 17}, 1},
		{"Test . 'A0123456789A012' ", args{[]byte("A0123456789A012"), Medium}, []byte{32, 121, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 0, 236, 17, 236}, 1},
		{"Test . 'A0123456789A0123' ", args{[]byte("A0123456789A0123"), Medium}, []byte{32, 129, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 0, 236, 17}, 1},
		{"Test . 'A0123456789A01234' ", args{[]byte("A0123456789A01234"), Medium}, []byte{32, 137, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 128, 236, 17}, 1},
		{"Test . 'A0123456789A012345' ", args{[]byte("A0123456789A012345"), Medium}, []byte{32, 145, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 185, 0, 236}, 1},
		{"Test . 'A0123456789A0123456' ", args{[]byte("A0123456789A0123456"), Medium}, []byte{32, 153, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 185, 24, 0}, 1},
		{"Test . 'A0123456789A01234567' ", args{[]byte("A0123456789A01234567"), Medium}, []byte{32, 161, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 185, 34, 160}, 1},
		{"Test . 'A0123456789A012345678' ", args{[]byte("A0123456789A012345678"), Medium}, []byte{32, 169, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 185, 34, 164, 0, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17, 236}, 2},
		{"Test . 'A0123456789A0123456789' ", args{[]byte("A0123456789A0123456789"), Medium}, []byte{32, 177, 194, 5, 226, 44, 115, 148, 51, 62, 0, 66, 232, 185, 34, 165, 196, 0, 236, 17, 236, 17, 236, 17, 236, 17, 236, 17}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fillAlphanumeric(tt.args.data, tt.args.level)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillAlphanumeric() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fillAlphanumeric() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
