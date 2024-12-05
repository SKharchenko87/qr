package qr

import (
	"slices"
)

// getVersion - получаем версию по кол-ву бит
func getVersion(level levelCorrection, bitSizeData int) int {
	version, _ := slices.BinarySearch(levelToCountBits[level], bitSizeData)
	return version + 1
}

// ToDo fillKanji

// ToDo fillAlphanumeric
func fillAlphanumeric(data []byte, level levelCorrection) ([]byte, byte) {
	l := len(data)
	var bitSizeData int
	if l%2 == 1 {
		bitSizeData = (l-1)/2*11 + 6
	} else {
		bitSizeData = l / 2 * 11
	}

	version, _ := slices.BinarySearch(levelToCountBits[level], bitSizeData)
	bitSizeFieldData := lengthFieldData(version, alphanumeric)
	if bitSizeData+bitSizeFieldData+4 > levelToCountBits[level][version] {
		version++
		bitSizeFieldData = lengthFieldData(version, alphanumeric)
	}
	res := make([]byte, levelToCountBits[level][version]/8)
	index := 0
	res[index] = byte(alphanumeric << 4)

	offset := 3
	for j := bitSizeFieldData - 1; j >= 0; j-- {
		res[index] |= byte((l >> j & 1) << offset)
		offset--
		if offset < 0 {
			offset = 7
			index++
		}
	}

	var v int16
	for i := 0; i < l/2; i++ {
		v = alphanumericDictionary[data[i*2+0]]*45 + alphanumericDictionary[data[i*2+1]]
		for j := 10; j >= 0; j-- {
			res[index] |= byte((v >> j & 1) << offset)
			offset--
			if offset < 0 {
				offset = 7
				index++
			}
		}
	}
	if l%2 == 1 {
		v = alphanumericDictionary[data[l-1]]
		for j := 5; j >= 0; j-- {
			res[index] |= byte((v >> j & 1) << offset)
			offset--
			if offset < 0 {
				offset = 7
				index++
			}
		}
	}
	//ToDo Padding
	for j := 3; j >= 0 && index < levelToCountBits[level][version]/8; j-- {
		offset--
		if offset < 0 {
			offset = 7
			index++
		}
	}
	if index != levelToCountBits[level][version]/8 {
		if offset != 7 {
			//if res[index] != 7 {
			index++
		}

		tmp := [2]byte{0b11101100, 0b00010001}
		tmpIndex := 0
		for ; index < levelToCountBits[level][version]/8; index++ {
			res[index] = tmp[tmpIndex]
			tmpIndex = (tmpIndex + 1) % 2
		}
	}
	return res, byte(version + 1)
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
func fillBinary(data []byte, level levelCorrection) ([]byte, byte) {
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

// fillNumeric - заполнение цифровое
func fillNumeric(data []byte, level levelCorrection) ([]byte, byte) {
	l := len(data)
	var bitSizeData int
	if l%3 == 2 {
		bitSizeData = l/3*10 + 7
	} else if l%3 == 1 {
		bitSizeData = l/3*10 + 4
	} else {
		bitSizeData = l / 3 * 10
	}

	version, _ := slices.BinarySearch(levelToCountBits[level], bitSizeData)
	bitSizeFieldData := lengthFieldData(version, digital)
	if bitSizeData+bitSizeFieldData+4 > levelToCountBits[level][version] {
		version++
		bitSizeFieldData = lengthFieldData(version, digital)
	}
	res := make([]byte, levelToCountBits[level][version]/8)
	index := 0
	res[index] = byte(digital << 4)

	offset := 3
	for j := bitSizeFieldData - 1; j >= 0; j-- {
		res[index] |= byte((l >> j & 1) << offset)
		offset--
		if offset < 0 {
			offset = 7
			index++
		}
	}
	var v int16
	for i := 0; i < l/3; i++ {
		v = int16(data[i*3+0]-'0')*100 + int16(data[i*3+1]-'0')*10 + int16(data[i*3+2]-'0')
		for j := 9; j >= 0; j-- {
			res[index] |= byte((v >> j & 1) << offset)
			offset--
			if offset < 0 {
				offset = 7
				index++
			}
		}
	}

	if l%3 == 2 {
		v = int16(data[l-2]-'0')*10 + int16(data[l-1]-'0')
		for j := 6; j >= 0; j-- {
			res[index] |= byte((v >> j & 1) << offset)
			offset--
			if offset < 0 {
				offset = 7
				index++
			}
		}
	} else if l%3 == 1 {
		v = int16(data[l-1] - '0')
		for j := 3; j >= 0; j-- {
			res[index] |= byte((v >> j & 1) << offset)
			offset--
			if offset < 0 {
				offset = 7
				index++
			}
		}
	}
	//ToDo Padding
	for j := 3; j >= 0 && index < levelToCountBits[level][version]/8; j-- {
		offset--
		if offset < 0 {
			offset = 7
			index++
		}
	}
	if index != levelToCountBits[level][version]/8 {
		if res[index] != 7 {
			index++
		}

		tmp := [2]byte{0b11101100, 0b00010001}
		tmpIndex := 0
		for ; index < levelToCountBits[level][version]/8; index++ {
			res[index] = tmp[tmpIndex]
			tmpIndex = (tmpIndex + 1) % 2
		}
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

func generateQR(text string, level levelCorrection) [][]bool {
	data, version := fillBinary([]byte(text), level)
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
