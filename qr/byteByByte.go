package qr

import (
	"fmt"
	"math"
	"slices"
)

// getVersion - получаем версию по кол-ву бит
func getVersion(level levelCorrection, bitSizeData int) int {
	version, _ := slices.BinarySearch(levelToCountBits[level], bitSizeData)
	return version + 1
}

//101110001100000
//________ ________
//    0100 01100100
//<Способ кодирования><Количество данных><исходная последовательность>

// Заполнение
// На выходе последовательность бит кратная 8 и
// если количество бит в текущей последовательности байт меньше того,
// которое нужно для выбранной версии, то оно дополнено чередующимися
// байтами 11101100 и 00010001.
func fill(data []byte, level levelCorrection) ([]byte, byte) {
	l := len(data)
	bitSizeData := l * 8
	version, _ := slices.BinarySearch(levelToCountBits[level], bitSizeData)
	bitSizeFieldData := lengthFieldData(version, byteByByte)
	if bitSizeData+bitSizeFieldData+4 > levelToCountBits[level][version] {
		version++
		bitSizeFieldData = lengthFieldData(version, byteByByte)
	}
	res := make([]byte, levelToCountBits[level][version]/8)
	res[0] = byte(byteByByte << 4)
	index := 0
	if bitSizeFieldData == 16 {
		res[index] |= byte(l & 0b0000_0000_0000_0000_1111_0000_0000_0000 >> 12)
		index++
		res[index] |= byte(l & 0b0000_0000_0000_0000_0000_1111_0000_0000 >> 4)
	}
	res[index] |= byte(l & 0b0000_0000_0000_0000_0000_0000_1111_0000 >> 4)
	index++
	res[index] |= byte(l & 0b0000_0000_0000_0000_0000_0000_0000_1111 << 4)
	i := 0
	for ; i < l; i++ {
		res[index] |= byte(data[i] & 0b0000_0000_0000_0000_0000_0000_1111_0000 >> 4)
		index++
		res[index] |= byte(data[i] & 0b0000_0000_0000_0000_0000_0000_0000_1111 << 4)
	}
	index++
	tmp := [2]byte{0b11101100, 0b00010001}
	tmpIndex := 0
	for ; index < levelToCountBits[level][version]/8; i++ {
		res[index] = tmp[tmpIndex]
		tmpIndex = (tmpIndex + 1) % 2
		index++
	}
	return res, byte(version + 1)
}

// Разделение информации на блоки - Определение количество байт в каждом блоке
func getCountByteOfBlock(level levelCorrection, version byte) []int {
	bytesData := levelToCountBits[level][version] / 8
	byteOfBlock := bytesData / CountOfBlocks[level][version]
	reminderByteOfBlock := bytesData % CountOfBlocks[level][version]
	res := make([]int, CountOfBlocks[level][version])
	index := CountOfBlocks[level][version] - 1
	for ; reminderByteOfBlock > 0 && index >= 0; index, reminderByteOfBlock = index-1, reminderByteOfBlock-1 {
		res[index] = byteOfBlock + 1
	}
	for ; index >= 0; index-- {
		res[index] = byteOfBlock
	}
	return res
}

// Заполнение блоков
// Блок заполняется байтами из данных полностью.
// Когда текущий блок полностью заполняется, очередь переходит к следующему.
// Байтов данных должно хватить ровно на все блоки, ни больше и ни меньше.
func fillBlocks(data []byte, countByteOfBlock []int) [][]byte {
	l := len(countByteOfBlock)
	res := make([][]byte, l)
	index := 0
	for i, cnt := range countByteOfBlock {
		res[i] = make([]byte, cnt)
		for j := 0; j < cnt; j++ {
			res[i][j] = data[index]
			index++
		}
	}
	return res
}

// createByteCorrection Создание байтов коррекции
func createByteCorrection(level levelCorrection, version byte, data *[]byte) []byte {
	l := len(*data)
	lengthCorrectionBytes := NumberOfCorrectionBytesPerBlock[level][version] // lengthCorrectionBytes - Количество байтов коррекции на один блок.
	polynomials := GeneratePolynomial[lengthCorrectionBytes]                 // polynomials - Генерирующие многочлены.
	correctionBytes := make([]byte, l+int(lengthCorrectionBytes))            // correctionBytes - Блок байтов коррекции(подготовленный массив)
	// Копируем исходный блок данных
	for i := 0; i < l; i++ {
		correctionBytes[i] = (*data)[i]
	}

	// Выполняем алгоритм по количеству байт в исходном блоке данных
	for i := 0; i < l; i++ {
		element := correctionBytes[i] // element - обрабатываемый элемент
		if element > 0 {
			reverseElement := int(ReverseGaloisField[element]) // reverseElement - обратный элемент поля Галуа
			for j := 0; j < int(lengthCorrectionBytes); j++ {
				mask := (int(polynomials[j]) + reverseElement) % 255
				correctionBytes[j+i+1] ^= GaloisField[mask]
			}
		}
	}
	return correctionBytes[l:]
}

// mergeBlocks Объединение блоков
func mergeBlocks(data, correction [][]byte) []byte {
	lengthRes := 0
	for i := range data {
		lengthRes += len(data[i])
	}
	for i := range correction {
		lengthRes += len(correction[i])
	}
	index := 0
	res := make([]byte, lengthRes)
	//ToDo
	//for _, ds := range data {
	//	for _, d := range ds {
	//		res[index] = d
	//		index++
	//	}
	//}
	//for _, ds := range correction {
	//	for _, d := range ds {
	//		res[index] = d
	//		index++
	//	}
	//}
	//return res

	var f = func(arr *[][]byte) {
		l0 := len((*arr)[0])
		for j := 0; j < l0; j++ {
			for i := range *arr {
				res[index] = (*arr)[i][j]
				index++
			}
		}
		for i := range *arr {
			if len((*arr)[i]) > l0 {
				res[index] = (*arr)[i][l0]
				index++
			}
		}
	}
	f(&data)
	f(&correction)

	return res
}

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
func drawSearchNodes(canvas *[][]bool, busyRangeModuls *[]Rectangle) {
	drawSearchNode(canvas, 3, 3)
	drawSearchNode(canvas, byte(len(*canvas)-4), 3)
	drawSearchNode(canvas, 3, byte(len(*canvas)-4))
	lengthCanvas := byte(len(*canvas))
	*busyRangeModuls = append(*busyRangeModuls, Rectangle{0, 0, 8, 8})
	*busyRangeModuls = append(*busyRangeModuls, Rectangle{0, lengthCanvas - 8, 8, lengthCanvas - 1})
	*busyRangeModuls = append(*busyRangeModuls, Rectangle{lengthCanvas - 8, 0, lengthCanvas - 1, 8})
}

// drawAlignmentNode наносим на холст выравнивающий узор
func drawAlignmentNode(canvas *[][]bool, i, j byte) {
	(*canvas)[i][j] = Black
	drawSquare(canvas, 4, i-2, j-2)
}

// drawSearchNodes наносим на холст выравнивающие узоры
func drawAlignmentNodes(canvas *[][]bool, version byte, busyRangeModuls *[]Rectangle) {
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
			*busyRangeModuls = append(*busyRangeModuls, Rectangle{locationAlignmentPatterns[i] - 2, locationAlignmentPatterns[j] - 2, locationAlignmentPatterns[i] + 2, locationAlignmentPatterns[j] + 2})
		}
	}
}

// drawSynchronizationLine наносим на холст линии синхронизации
func drawSynchronizationLines(canvas *[][]bool, busyRangeModuls *[]Rectangle) {
	y := 6
	for x := 8; x < len(*canvas)-8; x += 2 {
		(*canvas)[x][y] = I
		(*canvas)[y][x] = I
	}
	lengthCanvas := byte(len(*canvas))
	*busyRangeModuls = append(*busyRangeModuls, Rectangle{6, 0, 6, lengthCanvas - 1})
	*busyRangeModuls = append(*busyRangeModuls, Rectangle{0, 6, lengthCanvas - 1, 6})
}

// drawCodeVersion наносим код версии
func drawCodeVersion(canvas *[][]bool, version byte, busyRangeModuls *[]Rectangle) {
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
		*busyRangeModuls = append(*busyRangeModuls, Rectangle{0, lengthCanvas - 11, 6, lengthCanvas - 9})
		*busyRangeModuls = append(*busyRangeModuls, Rectangle{lengthCanvas - 11, 0, lengthCanvas - 9, 6})
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

	busyRangeModuls := []Rectangle{}

	drawSearchNodes(&canvas, &busyRangeModuls)
	drawSynchronizationLines(&canvas, &busyRangeModuls)
	drawAlignmentNodes(&canvas, version, &busyRangeModuls)
	drawCodeVersion(&canvas, version, &busyRangeModuls)

	return canvas, busyRangeModuls
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

func checkFree(busyRangeModuls *[]Rectangle, i, j byte) bool {
	for _, rectangle := range *busyRangeModuls {
		if rectangle.iLeftUp <= i && i <= rectangle.iRightDown &&
			rectangle.jLeftUp <= j && j <= rectangle.jRightDown {
			return false
		}
	}
	return true
}

func nextPosition(busyRangeModuls *[]Rectangle, i, j byte, lengthCanvas byte) (byte, byte) {
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
		flg = !checkFree(busyRangeModuls, byte(candidateI), byte(candidateJ))
	}
	return byte(candidateI), byte(candidateJ)
}

// generatePreCode генерируем qr код без масок
func generatePreCode(data []byte, canvas *[][]bool, busyRangeModuls *[]Rectangle) {
	lengthCanvas := byte(len(*canvas))
	i, j := lengthCanvas-1, lengthCanvas-1
	for _, d := range data {
		for _, b := range bitMap {
			if d&b > 0 {
				(*canvas)[i][j] = I
			}
			i, j = nextPosition(busyRangeModuls, i, j, lengthCanvas)
		}
	}
}

// getScoreRule1 вычисляем балы для выбора маски по Правилу 1:
// горизонтальный подсчет
func getScoreRule1(canvas *[][]bool) int {
	lengthCanvas := len(*canvas)
	rule1Score := 0
	for i := 0; i < lengthCanvas; i++ {
		color := (*canvas)[i][0]
		cnt := 1
		for j := 1; j < lengthCanvas; j++ {
			if color == (*canvas)[i][j] {
				cnt++
			} else {
				color = (*canvas)[i][j]
				if cnt >= 5 {
					rule1Score += (cnt - 2)
				}
				cnt = 1
			}
		}
		if cnt >= 5 {
			rule1Score += cnt - 2
		}
	}
	return rule1Score
}

// getScoreRule2 вычисляем балы для выбора маски по Правилу 2:
// вертикальный подсчет
func getScoreRule2(canvas *[][]bool) int {
	lengthCanvas := len(*canvas)
	rule2Score := 0
	for j := 0; j < lengthCanvas; j++ {
		color := (*canvas)[0][j]
		cnt := 1
		for i := 1; i < lengthCanvas; i++ {
			if color == (*canvas)[i][j] {
				cnt++
			} else {
				color = (*canvas)[i][j]
				if cnt >= 5 {
					rule2Score += (cnt - 2)
				}
				cnt = 1
			}
		}
		if cnt >= 5 {
			rule2Score += cnt - 2
		}
	}
	return rule2Score
}

// getScoreRule3 вычисляем балы для выбора маски по Правилу 3:
// За каждый квадрат модулей одного цвета размером 2 на 2 начисляется по 3 очка.
func getScoreRule3(canvas *[][]bool) int {
	lengthCanvas := len(*canvas)
	rule3Score := 0
	for i := 0; i < lengthCanvas-1; i++ {
		for j := 0; j < lengthCanvas-1; j++ {
			if (*canvas)[i][j] == (*canvas)[i+1][j] &&
				(*canvas)[i][j] == (*canvas)[i][j+1] &&
				(*canvas)[i][j] == (*canvas)[i+1][j+1] {
				rule3Score += 3
			}
		}
	}
	return rule3Score
}

// getScoreRule4 вычисляем балы для выбора маски по Правилу 4:
// За каждую последовательность по горизонтали модулей ЧБЧЧЧБЧ,
// с 4-мя белыми модулями с одной из сторон (или с 2-х сразу),
// добавляется 40 очков
func getScoreRule4(canvas *[][]bool) int {
	lengthCanvas := byte(len(*canvas))
	bitRule := 0b11111111111
	ruleL := 0b00001011101
	ruleR := 0b10111010000
	m := map[[2]byte]struct{}{}
	var i, j byte
	for i = 0; i < lengthCanvas; i++ {
		x := 0
		for j = 0; j < lengthCanvas; j++ {
			x <<= 1
			if (*canvas)[i][j] {
				x += 1
			}
			if j > 9 && x&bitRule == ruleL {
				m[[2]byte{i, j - 3}] = struct{}{}
			}
			if x&bitRule == ruleR {
				m[[2]byte{i, j - 7}] = struct{}{}
			}
		}
	}
	return len(m) * 40
}

// getScoreRule5 вычисляем балы для выбора маски по Правилу 4:
// За каждую последовательность по вертикали модулей ЧБЧЧЧБЧ,
// с 4-мя белыми модулями с одной из сторон (или с 2-х сразу),
// добавляется 40 очков
func getScoreRule5(canvas *[][]bool) int {
	lengthCanvas := byte(len(*canvas))
	bitRule := 0b11111111111
	ruleL := 0b00001011101
	ruleR := 0b10111010000
	m := map[[2]byte]struct{}{}
	var i, j byte
	for j = 0; j < lengthCanvas; j++ {
		x := 0
		for i = 0; i < lengthCanvas; i++ {
			x <<= 1
			if (*canvas)[i][j] {
				x += 1
			}
			if i > 9 && x&bitRule == ruleL {
				m[[2]byte{j, i - 3}] = struct{}{}
			}
			if x&bitRule == ruleR {
				m[[2]byte{j, i - 7}] = struct{}{}
			}
		}
	}
	return len(m) * 40
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

// getScoreRule6 вычисляем балы по соотношения количества чёрных и белых модулей
func getScoreRule6(canvas *[][]bool) int {
	lengthCanvas := byte(len(*canvas))
	cntI := 0
	var i, j byte
	for i = 0; i < lengthCanvas; i++ {
		for j = 0; j < lengthCanvas; j++ {
			if (*canvas)[i][j] {
				cntI++
			}
		}
	}
	return abs(int(math.Floor(float64(cntI*100)/float64(lengthCanvas)/float64(lengthCanvas)-50.0)) * 2)
}

// getScore вычисляем балы для выбора маски по горизонтали
func getScore(canvas *[][]bool) int {
	return getScoreRule1(canvas) +
		getScoreRule2(canvas) +
		getScoreRule3(canvas) +
		getScoreRule4(canvas) +
		getScoreRule5(canvas) +
		getScoreRule6(canvas)
}

// drawCodeMaskLevelCorrection наносим на холст код маски и уровень коррекции
func drawCodeMaskLevelCorrection(canvas *[][]bool, level levelCorrection, mask byte) {
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
		//{-3, 8}, {-4, 8}, {-5, 8}, {-6, 8}, {-7, 8}, {8, -8}, {8, -7}, {8, -6}, {8, -5}, {8, -4}, {8, -3}, {8, -2}, {8, -1}
		{8, -1}, {8, -2}, {8, -3}, {8, -4}, {8, -5}, {8, -6}, {8, -7}, {8, -8}, {-7, 8}, {-6, 8}, {-5, 8}, {-4, 8}, {-3, 8},
	}
	i1 := [][]int{
		//{8, 2}, {8, 3}, {8, 4}, {8, 5}, {8, 7}, {8, 8}, {7, 8}, {5, 8}, {4, 8}, {3, 8}, {2, 8}, {1, 8}, {0, 8},
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
func drawMask(canvas *[][]bool, busyRangeModuls *[]Rectangle, oldMask, newMask byte) {
	lengthCanvas := byte(len(*canvas))
	x, y := lengthCanvas-1, lengthCanvas-1
	for x != 255 && y != 255 {
		// ToDo
		//fmt.Print(y, x, (*canvas)[y][x], " ")
		if Masks[oldMask](x, y) == 0 && Masks[newMask](x, y) != 0 ||
			Masks[oldMask](x, y) != 0 && Masks[newMask](x, y) == 0 {
			(*canvas)[y][x] = !(*canvas)[y][x]
		}
		// ToDo
		//fmt.Println((*canvas)[y][x])
		y, x = nextPosition(busyRangeModuls, y, x, lengthCanvas)
	}
	// ToDo
	//printQR(canvas)
}

// ToDo
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

func generateQR(text string, level levelCorrection) [][]bool {
	data, version := fill([]byte(text), level)
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
	oldMask := byte(8)
	for i := 0; i < 8; i++ {
		drawMask(&candidateCanvas, &busyRangeModuls, oldMask, byte(i))
		printQR(&candidateCanvas)
		//drawCodeMaskLevelCorrection(&candidateCanvas, level, byte(i))
		oldMask = byte(i)
	}

	//drawMask(&candidateCanvas, &busyRangeModuls, 8, 2)
	//drawCodeMaskLevelCorrection(&candidateCanvas, level, 2)

	// ToDo применить маску
	return canvas
}
