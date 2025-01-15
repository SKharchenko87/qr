package qr

import "fmt"

// direction направления для рисования квадратов
var direction = [][2]int8{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// drawSquare заполняем квадрат
func drawSquare(canvas *[][]bool, l int, x, y byte) {
	for d := 0; d < 4; d++ {
		for _, dd := range direction {
			for a := 0; a < l; a++ {
				x, y = x+byte(dd[0]), y+byte(dd[1])
				(*canvas)[x][y] = Black
			}
		}
	}
}

// drawSearchNode наносим на холст поисковой узор
func drawSearchNode(canvas *[][]bool, i, j byte) {
	(*canvas)[i][j] = Black
	drawSquare(canvas, 2, i-1, j-1)
	drawSquare(canvas, 6, i-3, j-3)
}

// drawSearchNodes наносим на холст поисковые узоры
func drawSearchNodes(canvas *[][]bool, busyRangeModul *[]Rectangle) {
	drawSearchNode(canvas, 3, 3)
	drawSearchNode(canvas, byte(len(*canvas)-4), 3)
	drawSearchNode(canvas, 3, byte(len(*canvas)-4))
	lengthCanvas := byte(len(*canvas))
	*busyRangeModul = append(*busyRangeModul, Rectangle{0, 0, 8, 8})
	*busyRangeModul = append(*busyRangeModul, Rectangle{0, lengthCanvas - 8, 8, lengthCanvas - 1})
	*busyRangeModul = append(*busyRangeModul, Rectangle{lengthCanvas - 8, 0, lengthCanvas - 1, 8})
}

// drawAlignmentNode наносим на холст выравнивающий узор
func drawAlignmentNode(canvas *[][]bool, i, j byte) {
	(*canvas)[i][j] = Black
	drawSquare(canvas, 4, i-2, j-2)
}

// drawSearchNodes наносим на холст выравнивающие узоры
func drawAlignmentNodes(canvas *[][]bool, version byte, busyRangeModul *[]Rectangle) {
	locationAlignmentPatterns := LocationAlignmentPatterns[version]
	lengthAlignmentPatterns := byte(len(locationAlignmentPatterns))
	var i, j byte
	for i = 0; i < lengthAlignmentPatterns; i++ {
		for j = 0; j < lengthAlignmentPatterns; j++ {
			if version > 6 {
				if i == 0 && j == 0 || i == 0 && j == lengthAlignmentPatterns-1 || i == lengthAlignmentPatterns-1 && j == 0 {
					continue
				}
			}
			drawAlignmentNode(canvas, locationAlignmentPatterns[i], locationAlignmentPatterns[j])
			*busyRangeModul = append(*busyRangeModul, Rectangle{locationAlignmentPatterns[i] - 2, locationAlignmentPatterns[j] - 2, locationAlignmentPatterns[i] + 2, locationAlignmentPatterns[j] + 2})
		}
	}
}

// drawSynchronizationLine наносим на холст линии синхронизации
func drawSynchronizationLines(canvas *[][]bool, busyRangeModul *[]Rectangle) {
	y := 6
	for x := 8; x < len(*canvas)-8; x += 2 {
		(*canvas)[x][y] = I
		(*canvas)[y][x] = I
	}
	lengthCanvas := byte(len(*canvas))
	*busyRangeModul = append(*busyRangeModul, Rectangle{6, 0, 6, lengthCanvas - 1})
	*busyRangeModul = append(*busyRangeModul, Rectangle{0, 6, lengthCanvas - 1, 6})
}

// drawCodeVersion наносим код версии
func drawCodeVersion(canvas *[][]bool, version byte, busyRangeModul *[]Rectangle) {
	if version > 6 {
		lengthCanvas := byte(len(*canvas))
		var i byte
		for i = 0; i < 3; i++ {
			arr := CodeVersion[version]
			for j := 0; j < 6; j++ {
				(*canvas)[i+lengthCanvas-11][j] = arr[i][j]
				(*canvas)[j][i+lengthCanvas-11] = arr[i][j]
			}
		}
		*busyRangeModul = append(*busyRangeModul, Rectangle{0, lengthCanvas - 11, 6, lengthCanvas - 9})
		*busyRangeModul = append(*busyRangeModul, Rectangle{lengthCanvas - 11, 0, lengthCanvas - 9, 6})
	}
}

type Rectangle struct {
	iLeftUp, jLeftUp, iRightDown, jRightDown byte
}

// generateInfoCanvas генерируем холст и заполняем информационные данные
func generateInfoCanvas(version byte) ([][]bool, []Rectangle) {
	lengthCanvas := LengthCanvas[version-1]
	canvas := make([][]bool, lengthCanvas)
	var i byte
	for i = 0; i < lengthCanvas; i++ {
		canvas[i] = make([]bool, lengthCanvas)
	}

	busyRangeModul := []Rectangle{}

	drawSearchNodes(&canvas, &busyRangeModul)
	drawSynchronizationLines(&canvas, &busyRangeModul)
	drawAlignmentNodes(&canvas, version, &busyRangeModul)
	drawCodeVersion(&canvas, version, &busyRangeModul)

	return canvas, busyRangeModul
}

var bitMap = []byte{
	0b1000_0000,
	0b0100_0000,
	0b0010_0000,
	0b0001_0000,
	0b0000_1000,
	0b0000_0100,
	0b0000_0010,
	0b0000_0001,
}

func checkFree(busyRangeModul *[]Rectangle, i, j byte) bool {
	for _, rectangle := range *busyRangeModul {
		if rectangle.iLeftUp <= i && i <= rectangle.iRightDown &&
			rectangle.jLeftUp <= j && j <= rectangle.jRightDown {
			return false
		}
	}
	return true
}

func nextPosition(busyRangeModul *[]Rectangle, i, j byte, lengthCanvas byte) (byte, byte) {
	var candidateI, candidateJ int
	candidateI, candidateJ = int(i), int(j)
	var flg bool = !(candidateI == 0 && candidateJ == 0)
	for flg {
		x := candidateJ % 4
		if candidateJ > 6 && x == 0 || candidateJ <= 6 && x == 3 {
			candidateJ--
		} else if candidateJ > 6 && x == 3 || candidateJ <= 6 && x == 2 {
			candidateI--
			candidateJ++
		} else if candidateJ > 6 && x == 2 || candidateJ <= 6 && x == 1 {
			candidateJ--
		} else if candidateJ > 6 && x == 1 || candidateJ <= 6 && x == 0 {
			candidateI++
			candidateJ++
		}
		if candidateI == -1 {
			candidateI = 0
			candidateJ -= 2
			if candidateJ == 6 {
				candidateJ--
			}
		} else if candidateI == int(lengthCanvas) {
			candidateI = int(lengthCanvas) - 1
			candidateJ -= 2
		}
		flg = !checkFree(busyRangeModul, byte(candidateI), byte(candidateJ))
	}
	return byte(candidateI), byte(candidateJ)
}

// generatePreCode генерируем qr код без масок
func generatePreCode(data []byte, canvas *[][]bool, busyRangeModul *[]Rectangle) {
	lengthCanvas := byte(len(*canvas))
	i, j := lengthCanvas-1, lengthCanvas-1
	for _, d := range data {
		for _, b := range bitMap {
			if d&b > 0 {
				(*canvas)[i][j] = I
			}
			i, j = nextPosition(busyRangeModul, i, j, lengthCanvas)
		}
	}
}

// drawCodeMaskLevelCorrection наносим на холст код маски и уровень коррекции
func drawCodeMaskLevelCorrection(canvas *[][]bool, level LevelCorrection, mask byte) {
	lengthCanvas := len(*canvas)
	//LevelCorrection
	(*canvas)[8][0] = (level>>1)&1 == 1
	(*canvas)[8][1] = level&1 == 1

	(*canvas)[lengthCanvas-1][8] = (level>>1)&1 == 1
	(*canvas)[lengthCanvas-2][8] = level&1 == 1

	//Mask
	var b bool
	v := CodeMaskLevelCorrection[level][mask]
	i0 := [][]int{
		{8, -1}, {8, -2}, {8, -3}, {8, -4}, {8, -5}, {8, -6}, {8, -7}, {8, -8}, {-7, 8}, {-6, 8}, {-5, 8}, {-4, 8}, {-3, 8},
	}
	i1 := [][]int{
		{0, 8}, {1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8}, {7, 8}, {8, 8}, {8, 7}, {8, 5}, {8, 4}, {8, 3}, {8, 2},
	}
	(*canvas)[lengthCanvas-8][8] = I
	for i := 0; i < len(i0); i++ {
		b = v&1 == 1
		(*canvas)[(lengthCanvas+i0[i][0])%lengthCanvas][(lengthCanvas+i0[i][1])%lengthCanvas] = b
		(*canvas)[(lengthCanvas+i1[i][0])%lengthCanvas][(lengthCanvas+i1[i][1])%lengthCanvas] = b
		v >>= 1
	}
}

// drawMask наносим на холст код маски и уровень коррекции
func drawMask(canvas *[][]bool, busyRangeModul *[]Rectangle, oldMask, newMask byte) {
	lengthCanvas := byte(len(*canvas))
	x, y := lengthCanvas-1, lengthCanvas-1
	for x != 255 && y != 255 {
		if Masks[oldMask](x, y) == 0 && Masks[newMask](x, y) != 0 ||
			Masks[oldMask](x, y) != 0 && Masks[newMask](x, y) == 0 {
			(*canvas)[y][x] = !(*canvas)[y][x]
		}
		y, x = nextPosition(busyRangeModul, y, x, lengthCanvas)
	}
}

func printQR(canvas *[][]bool) {
	lengthCanvas := len(*canvas)
	for i := 0; i < lengthCanvas; i++ {
		for j := 0; j < lengthCanvas; j++ {
			if (*canvas)[i][j] {
				fmt.Print("I\t")
			} else {
				fmt.Print(".\t")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
